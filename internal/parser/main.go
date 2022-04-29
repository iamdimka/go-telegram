package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/iamdimka/go-telegram/internal/htmlutil"
	"golang.org/x/net/html"
)

func must(err error, msg ...interface{}) {
	if err != nil {
		if len(msg) > 0 {
			fmt.Println(msg...)
		}
		panic(err)
	}
}

var knownStructs = map[string]bool{}

func main() {
	res, err := http.Get("https://core.telegram.org/bots/api")
	must(err, "could not perform the request")

	root, err := html.Parse(res.Body)
	must(err, "could not parse html document")

	body := htmlutil.NewNode(root).QuerySelector("body")

	models := make([]*APIStruct, 0)
	requests := make([]*APIMethod, 0)

	for _, x := range body.QuerySelectorAll("h4", "a.anchor") {
		res := parse(x.Parent())

		if s, ok := res.(*APIStruct); ok {
			models = append(models, s)
		}

		if m, ok := res.(*APIMethod); ok {
			requests = append(requests, m)
		}
	}

	must(writeJSON("models.json", models))
	must(writeJSON("requests.json", requests))
	must(writeGo("models.go", "telegram", models, func(it *APIStruct, buf *bytes.Buffer) {
		buf.WriteString("\n\n")
		writeMultilineComment(buf, "// ", it.Description)
		if it.Additional != "" {
			buf.WriteString("//\n")
			writeMultilineComment(buf, "// ", it.Additional)
		}
		buf.WriteString("type ")
		buf.WriteString(it.Name)
		buf.WriteString(" struct {\n")
		for _, f := range it.Fields {
			writeMultilineComment(buf, "\t// ", f.Description)
			buf.WriteByte('\t')
			buf.WriteString(toFieldName(f.Field))
			buf.WriteByte(' ')
			buf.WriteString(toGoType(f.Type))
			buf.WriteString(" `json:\"")
			buf.WriteString(f.Field)
			if f.Optional {
				buf.WriteString(",omitempty")
			}
			buf.WriteString("\"`\n")
		}
		buf.WriteByte('}')
	}))

	must(writeGo("requests.go", "telegram", requests, func(it *APIMethod, buf *bytes.Buffer) {
		buf.WriteString("\n\n")

		writeMultilineComment(buf, "// ", it.Description)
		if it.Additional != "" {
			buf.WriteString("//\n")
			writeMultilineComment(buf, "// ", it.Additional)
		}

		buf.WriteString("type ")
		buf.WriteByte(it.Name[0] - ('a' - 'A'))
		buf.WriteString(it.Name[1:])
		buf.WriteString("Request")
		buf.WriteString(" struct {\n")
		for _, p := range it.Params {
			writeMultilineComment(buf, "\t// ", p.Description)
			buf.WriteByte('\t')
			buf.WriteString(toFieldName(p.Name))
			buf.WriteByte(' ')
			buf.WriteString(toGoType(p.Type))
			buf.WriteString(" `json:\"")
			buf.WriteString(p.Name)
			if !p.Required {
				buf.WriteString(",omitempty")
			}
			buf.WriteString("\"`\n")
		}
		buf.WriteByte('}')
	}))

	must(writeGo("api.go", "telegram", requests, func(it *APIMethod, buf *bytes.Buffer) {
		returnType := toGoType(it.Return)

		buf.WriteString("\n\n")
		buf.WriteString("func (b *Bot) ")
		buf.WriteString(toFieldName(it.Name))
		buf.WriteByte('(')
		if len(it.Params) > 0 {
			buf.WriteString("request *")
			buf.WriteString(toFieldName(it.Name))
			buf.WriteString("Request")
		}
		buf.WriteString(") (result ")
		buf.WriteString(returnType)
		buf.WriteString(", err error) {\n")
		buf.WriteString("\terr = b.request(\"")
		buf.WriteString(it.Name)
		buf.WriteByte('"')
		if len(it.Params) > 0 {
			buf.WriteString(", request")
		} else {
			buf.WriteString(", nil")
		}
		buf.WriteString(", ")
		if returnType[0] != '*' {
			buf.WriteByte('&')
		}
		buf.WriteString("result)\n")
		buf.WriteString("\treturn\n")
		buf.WriteByte('}')
	}))
}

func toGoType(t string) string {
	if strings.Contains(t, ", ") {
		return "interface{}"
	}

	prefix := ""
	for t[0] == '[' {
		prefix += "[]"
		t = t[1 : len(t)-1]
	}

	switch t {
	case "boolean", "true":
		return prefix + "bool"

	case "string", "int", "int64":
		return prefix + t

	case "float":
		return prefix + "float32"

	default:
		if t[0] >= 'A' && t[0] <= 'Z' {
			if knownStructs[t] {
				return prefix + "*" + t
			}

			return prefix + "json.RawMessage"
		}

		panic(t)
	}
}

func writeMultilineComment(b *bytes.Buffer, prefix, data string) {
	if data == "" {
		return
	}

	for len(data) > 80 {
		idx := strings.LastIndexByte(data[:80], ' ')
		if idx < 0 {
			idx = strings.IndexByte(data, ' ')
		}

		if idx < 0 {
			break
		}

		b.WriteString(prefix)
		b.WriteString(strings.ReplaceAll(data[:idx], "\n", "\n"+prefix))
		b.WriteByte('\n')
		data = data[idx+1:]
	}

	if data != "" {
		b.WriteString(prefix)
		b.WriteString(strings.ReplaceAll(data, "\n", "\n"+prefix))
		b.WriteByte('\n')
	}
}

func toFieldName(name string) string {
	const upper = 'a' - 'A'
	name = string(name[0]-upper) + name[1:]

	for i := 1; i < len(name); i++ {
		if name[i] == '_' {
			name = name[:i] + string(name[i+1]-upper) + name[i+2:]
		}
	}

	return name
}

func writeJSON(filename string, data interface{}) error {
	os.Remove(filename)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func writeGo[T interface{}](filename, packageName string, data []T, fn func(it T, buf *bytes.Buffer)) error {
	b := new(bytes.Buffer)

	b.WriteString("package ")
	b.WriteString(packageName)
	b.WriteString("\n\n")
	b.WriteString(`import "encoding/json"`)

	for _, it := range data {
		fn(it, b)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = b.WriteTo(file)
	return err
}

func parse(node *htmlutil.Node) interface{} {
	next := node.Next("h4")
	title := node.InnerText()
	description := node.Next("p")
	quote := node.Next("blockquote")
	table := node.Next("table")
	isMethod := title[0] >= 'a' && title[0] <= 'z'

	desc := ""
	additional := ""

	if description != nil && description.Between(node, next) {
		desc = strings.TrimSpace(parseText(description))
	}

	if quote != nil && quote.Between(node, next) {
		additional = strings.TrimSpace(parseText(quote))
	}

	if table != nil && table.Between(node, next) {
		th := table.QuerySelector("th")
		if th != nil && strings.EqualFold(th.InnerText(), "field") {
			isMethod = false
		}
	} else {
		table = nil
	}

	if isMethod {
		method := &APIMethod{
			Name:        title,
			Description: desc,
			Additional:  additional,
			Return:      parseReturn(desc),
		}

		if table != nil {
			for _, row := range table.QuerySelectorAll("tbody", "tr") {
				td := row.QuerySelectorAll("td")
				name := td[0].InnerText()

				method.Params = append(method.Params, APIParameter{
					Name:        name,
					Type:        parseType(td[1], name),
					Required:    strings.EqualFold(td[2].InnerText(), "Yes"),
					Description: strings.TrimSpace(parseText(td[3])),
				})
			}
		}
		return method
	}

	if table == nil {
		return nil
	}

	knownStructs[title] = true

	s := &APIStruct{
		Name:        title,
		Description: desc,
		Additional:  additional,
	}

	for _, row := range table.QuerySelectorAll("tbody", "tr") {
		td := row.QuerySelectorAll("td")
		field := td[0].InnerText()

		s.Fields = append(s.Fields, APIField{
			Field:       field,
			Type:        parseType(td[1], field),
			Optional:    td[2].QuerySelector("em") != nil,
			Description: strings.TrimSpace(parseText(td[2])),
		})
	}

	return s
}

func parseReturn(text string) string {
	for _, tpl := range []string{
		`([aA]n [aA]rray of).+?([A-Z][a-zA-Z]+).+?is returned`,
		`[rR]eturns ([aA]rray of )?([a-zA-Z*]+) on success`,
		`[oO]n success, returns[^A-Z]+?(\*?[aA]rray\*? of )?([A-Z][a-zA-Z]+)`,
		`[oO]n success,[^A-Z]+?(\*?[aA]rray\*? of )?([A-Z][a-zA-Z*]+).+?is returned`,
		`[rR]eturns.+?([A-Z][a-zA-Z]+)`,
	} {
		if res := regexp.MustCompile(tpl).FindStringSubmatch(text); res != nil {
			l := len(res)
			if l > 2 && res[1] == "" {
				l--
				res[1] = res[l]
			}
			return formatType(res[l-1], l-2)
		}
	}

	panic("could not parse result: " + text)
}

func formatType(t string, array int) string {
	t = strings.Trim(t, "*")

	switch strings.ToLower(t) {
	case "float number":
		t = "float"

	case "integer":
		t = "int"

	case "int", "string", "boolean", "true", "float", "int64":
		t = strings.ToLower(t)

	case "array":
		panic("parsed as array")

	default:
		if t[0] >= 'A' && t[0] <= 'Z' {
			break
		}

		panic("unknown type: " + t)
	}

	if array > 0 {
		return strings.Repeat("[", array) + t + strings.Repeat("]", array)
	}

	return t
}

func parseText(node *htmlutil.Node) string {
	var b strings.Builder

	for node := node.FirstChild(); node != nil; node = node.NextSibling() {
		if node.IsElement() {
			switch node.TagName() {
			case "em":
				b.WriteByte('*')
				b.WriteString(parseText(node))
				b.WriteByte('*')

			case "strong":
				b.WriteString("**")
				b.WriteString(parseText(node))
				b.WriteString("**")

			case "a":
				b.WriteString(parseText(node))

			case "p":
				b.WriteString(parseText(node))

			case "code":
				b.WriteByte('`')
				b.WriteString(parseText(node))
				b.WriteByte('`')

			case "br":
				b.WriteByte('\n')

			case "img":
				alt := node.Attribute("alt")
				if alt != "" {
					b.WriteString(alt)
				} else {
					b.WriteByte('[')
					b.WriteString(node.Attribute("src"))
					b.WriteByte(']')
				}

			case "ul":
				b.WriteString(parseText(node))

			case "li":
				b.WriteByte('\n')
				b.WriteString("  - ")
				b.WriteString(parseText(node))

			default:
				log.Println(node.OuterHTML())
				b.WriteString(node.OuterHTML())
			}
		} else if node.IsText() {
			b.WriteString(node.OuterHTML())
		}

	}

	return strings.NewReplacer(
		"&lt;", "<",
		"&gt;", ">",
		"&#39;", "'",
		"&quot;", "\"",
		"“", "\"",
		"”", "\"",
	).Replace(b.String())
}

func parseType(node *htmlutil.Node, name string) string {
	if node == nil {
		return ""
	}

	data := strings.TrimSpace(strings.ToLower(node.InnerHTML()))
	array := 0

	for strings.HasPrefix(data, "array of ") {
		data = data[9:]
		array++
	}

	switch data {
	case "integer":
		if name == "id" || strings.Contains(name, "_id") {
			return formatType("int64", array)
		}

		return formatType("int", array)

	case "string", "boolean", "true", "float", "float number":
		return formatType(data, array)

	default:
		types := []string{}
		for _, a := range node.QuerySelectorAll("a") {
			types = append(types, a.InnerHTML())
		}

		if len(types) > 0 {
			for i, data := range types {
				types[i] = formatType(data, array)
			}

			data = strings.Join(types, ", ")
			break
		}

		if strings.Contains(data, " or ") {
			data = strings.ReplaceAll(data, " or ", ", ")
			break
		}

		log.Panic(data, " ")
	}

	return data
}

type APIField struct {
	Field       string `json:"field"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Optional    bool   `json:"optional,omitempty"`
}

type APIStruct struct {
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Additional  string     `json:"additional,omitempty"`
	Fields      []APIField `json:"fields,omitempty"`
}

type APIParameter struct {
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Description string `json:"description,omitempty"`
}

type APIMethod struct {
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Additional  string         `json:"additional,omitempty"`
	Params      []APIParameter `json:"params,omitempty"`
	Return      string         `json:"return,omitempty"`
}
