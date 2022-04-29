package telegram

import "encoding/json"

// This object represents an incoming update.
// At most **one** of the optional
// parameters can be present in any given update.
type Update struct {
	// The update's unique identifier. Update identifiers start from a certain
	// positive number and increase sequentially. This ID becomes especially handy if
	// you're using Webhooks, since it allows you to ignore repeated updates or to
	// restore the correct update sequence, should they get out of order. If there are
	// no new updates for at least a week, then identifier of the next update will be
	// chosen randomly instead of sequentially.
	UpdateId int64 `json:"update_id"`
	// *Optional*. New incoming message of any kind — text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`
	// *Optional*. New version of a message that is known to the bot and was edited
	EditedMessage *Message `json:"edited_message,omitempty"`
	// *Optional*. New incoming channel post of any kind — text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`
	// *Optional*. New version of a channel post that is known to the bot and was
	// edited
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	// *Optional*. New incoming inline query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	// *Optional*. The result of an inline query that was chosen by a user and sent to
	// their chat partner. Please see our documentation on the feedback collecting for
	// details on how to enable these updates for your bot.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	// *Optional*. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	// *Optional*. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
	// *Optional*. New incoming pre-checkout query. Contains full information about
	// checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
	// *Optional*. New poll state. Bots receive only updates about stopped polls and
	// polls, which are sent by the bot
	Poll *Poll `json:"poll,omitempty"`
	// *Optional*. A user changed their answer in a non-anonymous poll. Bots receive
	// new votes only in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
	// *Optional*. The bot's chat member status was updated in a chat. For private
	// chats, this update is received only when the bot is blocked or unblocked by the
	// user.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`
	// *Optional*. A chat member's status was updated in a chat. The bot must be an
	// administrator in the chat and must explicitly specify "chat_member" in the list
	// of *allowed_updates* to receive these updates.
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`
	// *Optional*. A request to join the chat has been sent. The bot must have the
	// *can_invite_users* administrator right in the chat to receive these updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`
}

// Contains information about the current status of a webhook.
//
// **Optional** fields may be not returned when irrelevant.
type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	Url string `json:"url"`
	// *True*, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate,omitempty"`
	// Number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`
	// *Optional*. Currently used webhook IP address
	IpAddress string `json:"ip_address,omitempty"`
	// *Optional*. Unix time for the most recent error that happened when trying to
	// deliver an update via webhook
	LastErrorDate int `json:"last_error_date,omitempty"`
	// *Optional*. Error message in human-readable format for the most recent error
	// that happened when trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message,omitempty"`
	// *Optional*. Unix time of the most recent error that happened when trying to
	// synchronize available updates with Telegram datacenters
	LastSynchronizationErrorDate int `json:"last_synchronization_error_date,omitempty"`
	// *Optional*. Maximum allowed number of simultaneous HTTPS connections to the
	// webhook for update delivery
	MaxConnections int `json:"max_connections,omitempty"`
	// *Optional*. A list of update types the bot is subscribed to. Defaults to all
	// update types except *chat_member*
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// This object represents a Telegram user or bot.
type User struct {
	// Unique identifier for this user or bot. This number may have more than 32
	// significant bits and some programming languages may have difficulty/silent
	// defects in interpreting it. But it has at most 52 significant bits, so a 64-bit
	// integer or double-precision float type are safe for storing this identifier.
	Id int64 `json:"id"`
	// *True*, if this user is a bot
	IsBot bool `json:"is_bot,omitempty"`
	// User's or bot's first name
	FirstName string `json:"first_name"`
	// *Optional*. User's or bot's last name
	LastName string `json:"last_name,omitempty"`
	// *Optional*. User's or bot's username
	Username string `json:"username,omitempty"`
	// *Optional*. IETF language tag of the user's language
	LanguageCode string `json:"language_code,omitempty"`
	// *Optional*. *True*, if the bot can be invited to groups. Returned only in getMe.
	CanJoinGroups bool `json:"can_join_groups,omitempty"`
	// *Optional*. *True*, if privacy mode is disabled for the bot. Returned only in
	// getMe.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`
	// *Optional*. *True*, if the bot supports inline queries. Returned only in getMe.
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}

// This object represents a chat.
type Chat struct {
	// Unique identifier for this chat. This number may have more than 32 significant
	// bits and some programming languages may have difficulty/silent defects in
	// interpreting it. But it has at most 52 significant bits, so a signed 64-bit
	// integer or double-precision float type are safe for storing this identifier.
	Id int64 `json:"id"`
	// Type of chat, can be either "private", "group", "supergroup" or "channel"
	Type string `json:"type"`
	// *Optional*. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`
	// *Optional*. Username, for private chats, supergroups and channels if available
	Username string `json:"username,omitempty"`
	// *Optional*. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`
	// *Optional*. Last name of the other party in a private chat
	LastName string `json:"last_name,omitempty"`
	// *Optional*. Chat photo. Returned only in getChat.
	Photo *ChatPhoto `json:"photo,omitempty"`
	// *Optional*. Bio of the other party in a private chat. Returned only in getChat.
	Bio string `json:"bio,omitempty"`
	// *Optional*. True, if privacy settings of the other party in the private chat
	// allows to use `tg://user?id=<user_id>` links only in chats with the user.
	// Returned only in getChat.
	HasPrivateForwards bool `json:"has_private_forwards,omitempty"`
	// *Optional*. Description, for groups, supergroups and channel chats. Returned
	// only in getChat.
	Description string `json:"description,omitempty"`
	// *Optional*. Primary invite link, for groups, supergroups and channel chats.
	// Returned only in getChat.
	InviteLink string `json:"invite_link,omitempty"`
	// *Optional*. The most recent pinned message (by sending date). Returned only in
	// getChat.
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// *Optional*. Default chat member permissions, for groups and supergroups.
	// Returned only in getChat.
	Permissions *ChatPermissions `json:"permissions,omitempty"`
	// *Optional*. For supergroups, the minimum allowed delay between consecutive
	// messages sent by each unpriviledged user; in seconds. Returned only in getChat.
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`
	// *Optional*. The time after which all messages sent to the chat will be
	// automatically deleted; in seconds. Returned only in getChat.
	MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`
	// *Optional*. True, if messages from the chat can't be forwarded to other chats.
	// Returned only in getChat.
	HasProtectedContent bool `json:"has_protected_content,omitempty"`
	// *Optional*. For supergroups, name of group sticker set. Returned only in
	// getChat.
	StickerSetName string `json:"sticker_set_name,omitempty"`
	// *Optional*. *True*, if the bot can change the group sticker set. Returned only
	// in getChat.
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
	// *Optional*. Unique identifier for the linked chat, i.e. the discussion group
	// identifier for a channel and vice versa; for supergroups and channel chats.
	// This identifier may be greater than 32 bits and some programming languages may
	// have difficulty/silent defects in interpreting it. But it is smaller than 52
	// bits, so a signed 64 bit integer or double-precision float type are safe for
	// storing this identifier. Returned only in getChat.
	LinkedChatId int64 `json:"linked_chat_id,omitempty"`
	// *Optional*. For supergroups, the location to which the supergroup is connected.
	// Returned only in getChat.
	Location *ChatLocation `json:"location,omitempty"`
}

// This object represents a message.
type Message struct {
	// Unique message identifier inside this chat
	MessageId int64 `json:"message_id"`
	// *Optional*. Sender of the message; empty for messages sent to channels. For
	// backward compatibility, the field contains a fake sender user in non-channel
	// chats, if the message was sent on behalf of a chat.
	From *User `json:"from,omitempty"`
	// *Optional*. Sender of the message, sent on behalf of a chat. For example, the
	// channel itself for channel posts, the supergroup itself for messages from
	// anonymous group administrators, the linked channel for messages automatically
	// forwarded to the discussion group.  For backward compatibility, the field
	// *from* contains a fake sender user in non-channel chats, if the message was
	// sent on behalf of a chat.
	SenderChat *Chat `json:"sender_chat,omitempty"`
	// Date the message was sent in Unix time
	Date int `json:"date"`
	// Conversation the message belongs to
	Chat *Chat `json:"chat"`
	// *Optional*. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from,omitempty"`
	// *Optional*. For messages forwarded from channels or from anonymous
	// administrators, information about the original sender chat
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
	// *Optional*. For messages forwarded from channels, identifier of the original
	// message in the channel
	ForwardFromMessageId int64 `json:"forward_from_message_id,omitempty"`
	// *Optional*. For forwarded messages that were originally sent in channels or by
	// an anonymous chat administrator, signature of the message sender if present
	ForwardSignature string `json:"forward_signature,omitempty"`
	// *Optional*. Sender's name for messages forwarded from users who disallow adding
	// a link to their account in forwarded messages
	ForwardSenderName string `json:"forward_sender_name,omitempty"`
	// *Optional*. For forwarded messages, date the original message was sent in Unix
	// time
	ForwardDate int `json:"forward_date,omitempty"`
	// *Optional*. True, if the message is a channel post that was automatically
	// forwarded to the connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`
	// *Optional*. For replies, the original message. Note that the Message object in
	// this field will not contain further *reply_to_message* fields even if it itself
	// is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	// *Optional*. Bot through which the message was sent
	ViaBot *User `json:"via_bot,omitempty"`
	// *Optional*. Date the message was last edited in Unix time
	EditDate int `json:"edit_date,omitempty"`
	// *Optional*. True, if the message can't be forwarded
	HasProtectedContent bool `json:"has_protected_content,omitempty"`
	// *Optional*. The unique identifier of a media message group this message belongs
	// to
	MediaGroupId string `json:"media_group_id,omitempty"`
	// *Optional*. Signature of the post author for messages in channels, or the
	// custom title of an anonymous group administrator
	AuthorSignature string `json:"author_signature,omitempty"`
	// *Optional*. For text messages, the actual UTF-8 text of the message, 0-4096
	// characters
	Text string `json:"text,omitempty"`
	// *Optional*. For text messages, special entities like usernames, URLs, bot
	// commands, etc. that appear in the text
	Entities []*MessageEntity `json:"entities,omitempty"`
	// *Optional*. Message is an animation, information about the animation. For
	// backward compatibility, when this field is set, the *document* field will also
	// be set
	Animation *Animation `json:"animation,omitempty"`
	// *Optional*. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`
	// *Optional*. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`
	// *Optional*. Message is a photo, available sizes of the photo
	Photo []*PhotoSize `json:"photo,omitempty"`
	// *Optional*. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`
	// *Optional*. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`
	// *Optional*. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`
	// *Optional*. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`
	// *Optional*. Caption for the animation, audio, document, photo, video or voice,
	// 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// *Optional*. For messages with a caption, special entities like usernames, URLs,
	// bot commands, etc. that appear in the caption
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`
	// *Optional*. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`
	// *Optional*. Message is a game, information about the game. More about games »
	Game *Game `json:"game,omitempty"`
	// *Optional*. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`
	// *Optional*. Message is a venue, information about the venue. For backward
	// compatibility, when this field is set, the *location* field will also be set
	Venue *Venue `json:"venue,omitempty"`
	// *Optional*. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`
	// *Optional*. New members that were added to the group or supergroup and
	// information about them (the bot itself may be one of these members)
	NewChatMembers []*User `json:"new_chat_members,omitempty"`
	// *Optional*. A member was removed from the group, information about them (this
	// member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	// *Optional*. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`
	// *Optional*. A chat photo was change to this value
	NewChatPhoto []*PhotoSize `json:"new_chat_photo,omitempty"`
	// *Optional*. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
	// *Optional*. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`
	// *Optional*. Service message: the supergroup has been created. This field can't
	// be received in a message coming through updates, because bot can't be a member
	// of a supergroup when it is created. It can only be found in reply_to_message if
	// someone replies to a very first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
	// *Optional*. Service message: the channel has been created. This field can't be
	// received in a message coming through updates, because bot can't be a member of
	// a channel when it is created. It can only be found in reply_to_message if
	// someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
	// *Optional*. Service message: auto-delete timer settings changed in the chat
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	// *Optional*. The group has been migrated to a supergroup with the specified
	// identifier. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
	// *Optional*. The supergroup has been migrated from a group with the specified
	// identifier. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	MigrateFromChatId int64 `json:"migrate_from_chat_id,omitempty"`
	// *Optional*. Specified message was pinned. Note that the Message object in this
	// field will not contain further *reply_to_message* fields even if it is itself a
	// reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// *Optional*. Message is an invoice for a payment, information about the invoice.
	// More about payments »
	Invoice *Invoice `json:"invoice,omitempty"`
	// *Optional*. Message is a service message about a successful payment,
	// information about the payment. More about payments »
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
	// *Optional*. The domain name of the website on which the user has logged in.
	// More about Telegram Login »
	ConnectedWebsite string `json:"connected_website,omitempty"`
	// *Optional*. Telegram Passport data
	PassportData *PassportData `json:"passport_data,omitempty"`
	// *Optional*. Service message. A user in the chat triggered another user's
	// proximity alert while sharing Live Location.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	// *Optional*. Service message: video chat scheduled
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`
	// *Optional*. Service message: video chat started
	VideoChatStarted json.RawMessage `json:"video_chat_started,omitempty"`
	// *Optional*. Service message: video chat ended
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`
	// *Optional*. Service message: new participants invited to a video chat
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	// *Optional*. Service message: data sent by a Web App
	WebAppData *WebAppData `json:"web_app_data,omitempty"`
	// *Optional*. Inline keyboard attached to the message. `login_url` buttons are
	// represented as ordinary `url` buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// This object represents a unique message identifier.
type MessageId struct {
	// Unique message identifier
	MessageId int64 `json:"message_id"`
}

// This object represents one special entity in a text message. For example,
// hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity. Currently, can be "mention" (`@username`), "hashtag"
	// (`#hashtag`), "cashtag" (`$USD`), "bot_command" (`/start@jobs_bot`), "url"
	// (`https://telegram.org`), "email" (`do-not-reply@telegram.org`), "phone_number"
	// (`+1-212-555-0123`), "bold" (**bold text**), "italic" (*italic text*),
	// "underline" (underlined text), "strikethrough" (strikethrough text), "spoiler"
	// (spoiler message), "code" (monowidth string), "pre" (monowidth block),
	// "text_link" (for clickable text URLs), "text_mention" (for users without
	// usernames)
	Type string `json:"type,omitempty"`
	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`
	// Length of the entity in UTF-16 code units
	Length int `json:"length"`
	// *Optional*. For "text_link" only, url that will be opened after user taps on
	// the text
	Url string `json:"url,omitempty"`
	// *Optional*. For "text_mention" only, the mentioned user
	User *User `json:"user,omitempty"`
	// *Optional*. For "pre" only, the programming language of the entity text
	Language string `json:"language,omitempty"`
}

// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and
	// for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Photo width
	Width int `json:"width"`
	// Photo height
	Height int `json:"height"`
	// *Optional*. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without
// sound).
type Animation struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and
	// for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Video width as defined by sender
	Width int `json:"width"`
	// Video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// *Optional*. Animation thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// *Optional*. Original animation filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// *Optional*. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// *Optional*. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// This object represents an audio file to be treated as music by the Telegram
// clients.
type Audio struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and
	// for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// *Optional*. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer,omitempty"`
	// *Optional*. Title of the audio as defined by sender or by audio tags
	Title string `json:"title,omitempty"`
	// *Optional*. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// *Optional*. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// *Optional*. File size in bytes
	FileSize int `json:"file_size,omitempty"`
	// *Optional*. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// This object represents a general file (as opposed to photos, voice messages and
// audio files).
type Document struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and
	// for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// *Optional*. Document thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// *Optional*. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// *Optional*. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// *Optional*. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// This object represents a video file.
type Video struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and
	// for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Video width as defined by sender
	Width int `json:"width"`
	// Video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// *Optional*. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// *Optional*. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// *Optional*. Mime type of a file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// *Optional*. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and
	// for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Video width and height (diameter of the video message) as defined by sender
	Length int `json:"length"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// *Optional*. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// *Optional*. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// This object represents a voice note.
type Voice struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and
	// for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// *Optional*. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// *Optional*. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// This object represents a phone contact.
type Contact struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// *Optional*. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// *Optional*. Contact's user identifier in Telegram. This number may have more
	// than 32 significant bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But it has at most 52 significant
	// bits, so a 64-bit integer or double-precision float type are safe for storing
	// this identifier.
	UserId int64 `json:"user_id,omitempty"`
	// *Optional*. Additional data about the contact in the form of a vCard
	Vcard string `json:"vcard,omitempty"`
}

// This object represents an animated emoji that displays a random value.
type Dice struct {
	// Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`
	// Value of the dice, 1-6 for "🎲", "🎯" and "🎳" base emoji, 1-5 for "🏀"
	// and "⚽" base emoji, 1-64 for "🎰" base emoji
	Value int `json:"value"`
}

// This object contains information about one answer option in a poll.
type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`
	// Number of users that voted for this option
	VoterCount int `json:"voter_count"`
}

// This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// Unique poll identifier
	PollId string `json:"poll_id"`
	// The user, who changed the answer to the poll
	User *User `json:"user"`
	// 0-based identifiers of answer options, chosen by the user. May be empty if the
	// user retracted their vote.
	OptionIds []int64 `json:"option_ids"`
}

// This object contains information about a poll.
type Poll struct {
	// Unique poll identifier
	Id string `json:"id"`
	// Poll question, 1-300 characters
	Question string `json:"question"`
	// List of poll options
	Options []*PollOption `json:"options"`
	// Total number of users that voted in the poll
	TotalVoterCount int `json:"total_voter_count"`
	// *True*, if the poll is closed
	IsClosed bool `json:"is_closed,omitempty"`
	// *True*, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// Poll type, currently can be "regular" or "quiz"
	Type string `json:"type"`
	// *True*, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`
	// *Optional*. 0-based identifier of the correct answer option. Available only for
	// polls in the quiz mode, which are closed, or was sent (not forwarded) by the
	// bot or to the private chat with the bot.
	CorrectOptionId int64 `json:"correct_option_id,omitempty"`
	// *Optional*. Text that is shown when a user chooses an incorrect answer or taps
	// on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation,omitempty"`
	// *Optional*. Special entities like usernames, URLs, bot commands, etc. that
	// appear in the *explanation*
	ExplanationEntities []*MessageEntity `json:"explanation_entities,omitempty"`
	// *Optional*. Amount of time in seconds the poll will be active after creation
	OpenPeriod int `json:"open_period,omitempty"`
	// *Optional*. Point in time (Unix timestamp) when the poll will be automatically
	// closed
	CloseDate int `json:"close_date,omitempty"`
}

// This object represents a point on the map.
type Location struct {
	// Longitude as defined by sender
	Longitude float32 `json:"longitude"`
	// Latitude as defined by sender
	Latitude float32 `json:"latitude"`
	// *Optional*. The radius of uncertainty for the location, measured in meters;
	// 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`
	// *Optional*. Time relative to the message sending date, during which the
	// location can be updated; in seconds. For active live locations only.
	LivePeriod int `json:"live_period,omitempty"`
	// *Optional*. The direction in which user is moving, in degrees; 1-360. For
	// active live locations only.
	Heading int `json:"heading,omitempty"`
	// *Optional*. Maximum distance for proximity alerts about approaching another
	// chat member, in meters. For sent live locations only.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

// This object represents a venue.
type Venue struct {
	// Venue location. Can't be a live location
	Location *Location `json:"location"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// *Optional*. Foursquare identifier of the venue
	FoursquareId string `json:"foursquare_id,omitempty"`
	// *Optional*. Foursquare type of the venue. (For example,
	// "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	FoursquareType string `json:"foursquare_type,omitempty"`
	// *Optional*. Google Places identifier of the venue
	GooglePlaceId string `json:"google_place_id,omitempty"`
	// *Optional*. Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// Contains data sent from a Web App to the bot.
type WebAppData struct {
	// The data. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`
	// Text of the *web_app* keyboard button, from which the Web App was opened. Be
	// aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text,omitempty"`
}

// This object represents the content of a service message, sent whenever a user
// in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	// User that triggered the alert
	Traveler *User `json:"traveler"`
	// User that set the alert
	Watcher *User `json:"watcher"`
	// The distance between the users
	Distance int `json:"distance"`
}

// This object represents a service message about a change in auto-delete timer
// settings.
type MessageAutoDeleteTimerChanged struct {
	// New auto-delete time for messages in the chat; in seconds
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// This object represents a service message about a video chat scheduled in the
// chat.
type VideoChatScheduled struct {
	// Point in time (Unix timestamp) when the video chat is supposed to be started by
	// a chat administrator
	StartDate int `json:"start_date"`
}

// This object represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	// Video chat duration in seconds
	Duration int `json:"duration"`
}

// This object represents a service message about new members invited to a video
// chat.
type VideoChatParticipantsInvited struct {
	// New members that were invited to the video chat
	Users []*User `json:"users"`
}

// This object represent a user's profile pictures.
type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`
	// Requested profile pictures (in up to 4 sizes each)
	Photos [][]*PhotoSize `json:"photos"`
}

// This object represents a file ready to be downloaded. The file can be
// downloaded via the link `https://api.telegram.org/file/bot<token>/<file_path>`.
// It is guaranteed that the link will be valid for at least 1 hour. When the link
// expires, a new one can be requested by calling getFile.
//
// Maximum file size to download is 20 MB
type File struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and
	// for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// *Optional*. File size in bytes, if known
	FileSize int `json:"file_size,omitempty"`
	// *Optional*. File path. Use
	// `https://api.telegram.org/file/bot<token>/<file_path>` to get the file.
	FilePath string `json:"file_path,omitempty"`
}

// Contains information about a Web App.
type WebAppInfo struct {
	// An HTTPS URL of a Web App to be opened with additional data as specified in
	// Initializing Web Apps
	Url string `json:"url"`
}

// This object represents a custom keyboard with reply options (see Introduction
// to bots for details and examples).
type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]*KeyboardButton `json:"keyboard"`
	// *Optional*. Requests clients to resize the keyboard vertically for optimal fit
	// (e.g., make the keyboard smaller if there are just two rows of buttons).
	// Defaults to *false*, in which case the custom keyboard is always of the same
	// height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
	// *Optional*. Requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display
	// the usual letter-keyboard in the chat – the user can press a special button
	// in the input field to see the custom keyboard again. Defaults to *false*.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	// *Optional*. The placeholder to be shown in the input field when the keyboard is
	// active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	// *Optional*. Use this parameter if you want to show the keyboard to specific
	// users only. Targets: 1) users that are @mentioned in the *text* of the Message
	// object; 2) if the bot's message is a reply (has *reply_to_message_id*), sender
	// of the original message.
	// 
	// *Example:* A user requests to change the bot's
	// language, bot replies to the request with a keyboard to select the new
	// language. Other users in the group don't see the keyboard.
	Selective bool `json:"selective,omitempty"`
}

// This object represents one button of the reply keyboard. For simple text
// buttons *String* can be used instead of this object to specify text of the
// button. Optional fields *web_app*, *request_contact*, *request_location*, and
// *request_poll* are mutually exclusive.
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used, it will be sent as
	// a message when the button is pressed
	Text string `json:"text"`
	// *Optional*. If *True*, the user's phone number will be sent as a contact when
	// the button is pressed. Available in private chats only.
	RequestContact bool `json:"request_contact,omitempty"`
	// *Optional*. If *True*, the user's current location will be sent when the button
	// is pressed. Available in private chats only.
	RequestLocation bool `json:"request_location,omitempty"`
	// *Optional*. If specified, the user will be asked to create a poll and send it
	// to the bot when the button is pressed. Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
	// *Optional*. If specified, the described Web App will be launched when the
	// button is pressed. The Web App will be able to send a "web_app_data" service
	// message. Available in private chats only.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// This object represents type of a poll, which is allowed to be created and sent
// when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// *Optional*. If *quiz* is passed, the user will be allowed to create only polls
	// in the quiz mode. If *regular* is passed, only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type.
	Type string `json:"type,omitempty"`
}

// Upon receiving a message with this object, Telegram clients will remove the
// current custom keyboard and display the default letter-keyboard. By default,
// custom keyboards are displayed until a new keyboard is sent by a bot. An
// exception is made for one-time keyboards that are hidden immediately after the
// user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	// Requests clients to remove the custom keyboard (user will not be able to summon
	// this keyboard; if you want to hide the keyboard from sight but keep it
	// accessible, use *one_time_keyboard* in ReplyKeyboardMarkup)
	RemoveKeyboard bool `json:"remove_keyboard,omitempty"`
	// *Optional*. Use this parameter if you want to remove the keyboard for specific
	// users only. Targets: 1) users that are @mentioned in the *text* of the Message
	// object; 2) if the bot's message is a reply (has *reply_to_message_id*), sender
	// of the original message.
	// 
	// *Example:* A user votes in a poll, bot returns
	// confirmation message in reply to the vote and removes the keyboard for that
	// user, while still showing the keyboard with poll options to users who haven't
	// voted yet.
	Selective bool `json:"selective,omitempty"`
}

// This object represents an inline keyboard that appears right next to the
// message it belongs to.
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of InlineKeyboardButton
	// objects
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// This object represents one button of an inline keyboard. You **must** use
// exactly one of the optional fields.
type InlineKeyboardButton struct {
	// Label text on the button
	Text string `json:"text"`
	// *Optional*. HTTP or tg:// url to be opened when the button is pressed. Links
	// `tg://user?id=<user_id>` can be used to mention a user by their ID without
	// using a username, if this is allowed by their privacy settings.
	Url string `json:"url,omitempty"`
	// *Optional*. Data to be sent in a callback query to the bot when button is
	// pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`
	// *Optional*. Description of the Web App that will be launched when the user
	// presses the button. The Web App will be able to send an arbitrary message on
	// behalf of the user using the method answerWebAppQuery. Available only in
	// private chats between a user and the bot.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
	// *Optional*. An HTTP URL used to automatically authorize the user. Can be used
	// as a replacement for the Telegram Login Widget.
	LoginUrl *LoginUrl `json:"login_url,omitempty"`
	// *Optional*. If set, pressing the button will prompt the user to select one of
	// their chats, open that chat and insert the bot's username and the specified
	// inline query in the input field. Can be empty, in which case just the bot's
	// username will be inserted.
	// 
	// **Note:** This offers an easy way for users to
	// start using your bot in inline mode when they are currently in a private chat
	// with it. Especially useful when combined with *switch_pm…* actions – in
	// this case the user will be automatically returned to the chat they switched
	// from, skipping the chat selection screen.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`
	// *Optional*. If set, pressing the button will insert the bot's username and the
	// specified inline query in the current chat's input field. Can be empty, in
	// which case only the bot's username will be inserted.
	// 
	// This offers a quick way
	// for the user to open your bot in inline mode in the same chat – good for
	// selecting something from multiple options.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
	// *Optional*. Description of the game that will be launched when the user presses
	// the button.
	// 
	// **NOTE:** This type of button **must** always be the first button
	// in the first row.
	CallbackGame json.RawMessage `json:"callback_game,omitempty"`
	// *Optional*. Specify *True*, to send a Pay button.
	// 
	// **NOTE:** This type of
	// button **must** always be the first button in the first row and can only be
	// used in invoice messages.
	Pay bool `json:"pay,omitempty"`
}

// This object represents a parameter of the inline keyboard button used to
// automatically authorize a user. Serves as a great replacement for the Telegram
// Login Widget when the user is coming from Telegram. All the user needs to do is
// tap/click a button and confirm that they want to log in:
//
// Sample bot: @discussbot
type LoginUrl struct {
	// An HTTP URL to be opened with user authorization data added to the query string
	// when the button is pressed. If the user refuses to provide authorization data,
	// the original URL without information about the user will be opened. The data
	// added is the same as described in Receiving authorization data.
	// 
	// **NOTE:** You
	// **must** always check the hash of the received data to verify the
	// authentication and the integrity of the data as described in Checking
	// authorization.
	Url string `json:"url"`
	// *Optional*. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`
	// *Optional*. Username of a bot, which will be used for user authorization. See
	// Setting up a bot for more details. If not specified, the current bot's username
	// will be assumed. The *url*'s domain must be the same as the domain linked with
	// the bot. See Linking your domain to the bot for more details.
	BotUsername string `json:"bot_username,omitempty"`
	// *Optional*. Pass *True* to request the permission for your bot to send messages
	// to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// This object represents an incoming callback query from a callback button in an
// inline keyboard. If the button that originated the query was attached to a
// message sent by the bot, the field *message* will be present. If the button was
// attached to a message sent via the bot (in inline mode), the field
// *inline_message_id* will be present. Exactly one of the fields *data* or
// *game_short_name* will be present.
//
// **NOTE:** After the user presses a callback button, Telegram clients will
// display a progress bar until you call answerCallbackQuery. It is, therefore,
// necessary to react by calling answerCallbackQuery even if no notification to
// the user is needed (e.g., without specifying any of the optional parameters).
type CallbackQuery struct {
	// Unique identifier for this query
	Id string `json:"id"`
	// Sender
	From *User `json:"from"`
	// *Optional*. Message with the callback button that originated the query. Note
	// that message content and message date will not be available if the message is
	// too old
	Message *Message `json:"message,omitempty"`
	// *Optional*. Identifier of the message sent via the bot in inline mode, that
	// originated the query.
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// Global identifier, uniquely corresponding to the chat to which the message with
	// the callback button was sent. Useful for high scores in games.
	ChatInstance string `json:"chat_instance"`
	// *Optional*. Data associated with the callback button. Be aware that a bad
	// client can send arbitrary data in this field.
	Data string `json:"data,omitempty"`
	// *Optional*. Short name of a Game to be returned, serves as the unique
	// identifier for the game
	GameShortName string `json:"game_short_name,omitempty"`
}

// Upon receiving a message with this object, Telegram clients will display a
// reply interface to the user (act as if the user has selected the bot's message
// and tapped 'Reply'). This can be extremely useful if you want to create
// user-friendly step-by-step interfaces without having to sacrifice privacy mode.
//
// **Example:** A poll bot for groups runs in privacy mode (only receives
// commands, replies to its messages and mentions). There could be two ways to
// create a new poll:
// 
// 
//   - Explain the user how to send a command with parameters
// (e.g. /newpoll question answer1 answer2). May be appealing for hardcore users
// but lacks modern day polish.
// 
//   - Guide the user through a step-by-step
// process. 'Please send me your question', 'Cool, now let's add the first answer
// option', 'Great. Keep adding answer options, then send /done when you're
// ready'.
// 
// The last option is definitely more attractive. And if you use
// ForceReply in your bot's questions, it will receive the user's answers even if
// it only receives replies, commands and mentions — without any extra work for
// the user.
type ForceReply struct {
	// Shows reply interface to the user, as if they manually selected the bot's
	// message and tapped 'Reply'
	ForceReply bool `json:"force_reply"`
	// *Optional*. The placeholder to be shown in the input field when the reply is
	// active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	// *Optional*. Use this parameter if you want to force reply from specific users
	// only. Targets: 1) users that are @mentioned in the *text* of the Message
	// object; 2) if the bot's message is a reply (has *reply_to_message_id*), sender
	// of the original message.
	Selective bool `json:"selective,omitempty"`
}

// This object represents a chat photo.
type ChatPhoto struct {
	// File identifier of small (160x160) chat photo. This file_id can be used only
	// for photo download and only for as long as the photo is not changed.
	SmallFileId string `json:"small_file_id"`
	// Unique file identifier of small (160x160) chat photo, which is supposed to be
	// the same over time and for different bots. Can't be used to download or reuse
	// the file.
	SmallFileUniqueId string `json:"small_file_unique_id"`
	// File identifier of big (640x640) chat photo. This file_id can be used only for
	// photo download and only for as long as the photo is not changed.
	BigFileId string `json:"big_file_id"`
	// Unique file identifier of big (640x640) chat photo, which is supposed to be the
	// same over time and for different bots. Can't be used to download or reuse the
	// file.
	BigFileUniqueId string `json:"big_file_unique_id"`
}

// Represents an invite link for a chat.
type ChatInviteLink struct {
	// The invite link. If the link was created by another chat administrator, then
	// the second part of the link will be replaced with "…".
	InviteLink string `json:"invite_link"`
	// Creator of the link
	Creator *User `json:"creator"`
	// *True*, if users joining the chat via the link need to be approved by chat
	// administrators
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
	// *True*, if the link is primary
	IsPrimary bool `json:"is_primary,omitempty"`
	// *True*, if the link is revoked
	IsRevoked bool `json:"is_revoked,omitempty"`
	// *Optional*. Invite link name
	Name string `json:"name,omitempty"`
	// *Optional*. Point in time (Unix timestamp) when the link will expire or has
	// been expired
	ExpireDate int `json:"expire_date,omitempty"`
	// *Optional*. Maximum number of users that can be members of the chat
	// simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit int `json:"member_limit,omitempty"`
	// *Optional*. Number of pending join requests created using this link
	PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`
}

// Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	// *True*, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// *True*, if the administrator can access the chat event log, chat statistics,
	// message statistics in channels, see channel members, see anonymous
	// administrators in supergroups and ignore slow mode. Implied by any other
	// administrator privilege
	CanManageChat bool `json:"can_manage_chat,omitempty"`
	// *True*, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
	// *True*, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
	// *True*, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	// *True*, if the administrator can add new administrators with a subset of their
	// own privileges or demote administrators that he has promoted, directly or
	// indirectly (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
	// *True*, if the user is allowed to change the chat title, photo and other
	// settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// *True*, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// *Optional*. *True*, if the administrator can post in the channel; channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	// *Optional*. *True*, if the administrator can edit messages of other users and
	// can pin messages; channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	// *Optional*. *True*, if the user is allowed to pin messages; groups and
	// supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// Represents a chat member that owns the chat and has all administrator
// privileges.
type ChatMemberOwner struct {
	// The member's status in the chat, always "creator"
	Status string `json:"status"`
	// Information about the user
	User *User `json:"user"`
	// *True*, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// *Optional*. Custom title for this user
	CustomTitle string `json:"custom_title,omitempty"`
}

// Represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
	// The member's status in the chat, always "administrator"
	Status string `json:"status"`
	// Information about the user
	User *User `json:"user"`
	// *True*, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited,omitempty"`
	// *True*, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// *True*, if the administrator can access the chat event log, chat statistics,
	// message statistics in channels, see channel members, see anonymous
	// administrators in supergroups and ignore slow mode. Implied by any other
	// administrator privilege
	CanManageChat bool `json:"can_manage_chat,omitempty"`
	// *True*, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
	// *True*, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
	// *True*, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	// *True*, if the administrator can add new administrators with a subset of their
	// own privileges or demote administrators that he has promoted, directly or
	// indirectly (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
	// *True*, if the user is allowed to change the chat title, photo and other
	// settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// *True*, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// *Optional*. *True*, if the administrator can post in the channel; channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	// *Optional*. *True*, if the administrator can edit messages of other users and
	// can pin messages; channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	// *Optional*. *True*, if the user is allowed to pin messages; groups and
	// supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	// *Optional*. Custom title for this user
	CustomTitle string `json:"custom_title,omitempty"`
}

// Represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	// The member's status in the chat, always "member"
	Status string `json:"status"`
	// Information about the user
	User *User `json:"user"`
}

// Represents a chat member that is under certain restrictions in the chat.
// Supergroups only.
type ChatMemberRestricted struct {
	// The member's status in the chat, always "restricted"
	Status string `json:"status"`
	// Information about the user
	User *User `json:"user"`
	// *True*, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member,omitempty"`
	// *True*, if the user is allowed to change the chat title, photo and other
	// settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// *True*, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// *True*, if the user is allowed to pin messages
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	// *True*, if the user is allowed to send text messages, contacts, locations and
	// venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	// *True*, if the user is allowed to send audios, documents, photos, videos, video
	// notes and voice notes
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	// *True*, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls,omitempty"`
	// *True*, if the user is allowed to send animations, games, stickers and use
	// inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	// *True*, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	// Date when restrictions will be lifted for this user; unix time. If 0, then the
	// user is restricted forever
	UntilDate int `json:"until_date"`
}

// Represents a chat member that isn't currently a member of the chat, but may
// join it themselves.
type ChatMemberLeft struct {
	// The member's status in the chat, always "left"
	Status string `json:"status"`
	// Information about the user
	User *User `json:"user"`
}

// Represents a chat member that was banned in the chat and can't return to the
// chat or view chat messages.
type ChatMemberBanned struct {
	// The member's status in the chat, always "kicked"
	Status string `json:"status"`
	// Information about the user
	User *User `json:"user"`
	// Date when restrictions will be lifted for this user; unix time. If 0, then the
	// user is banned forever
	UntilDate int `json:"until_date"`
}

// This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	// Chat the user belongs to
	Chat *Chat `json:"chat"`
	// Performer of the action, which resulted in the change
	From *User `json:"from"`
	// Date the change was done in Unix time
	Date int `json:"date"`
	// Previous information about the chat member
	OldChatMember json.RawMessage `json:"old_chat_member"`
	// New information about the chat member
	NewChatMember json.RawMessage `json:"new_chat_member"`
	// *Optional*. Chat invite link, which was used by the user to join the chat; for
	// joining by invite link events only.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// Represents a join request sent to a chat.
type ChatJoinRequest struct {
	// Chat to which the request was sent
	Chat *Chat `json:"chat"`
	// User that sent the join request
	From *User `json:"from"`
	// Date the request was sent in Unix time
	Date int `json:"date"`
	// *Optional*. Bio of the user.
	Bio string `json:"bio,omitempty"`
	// *Optional*. Chat invite link that was used by the user to send the join request
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// *Optional*. *True*, if the user is allowed to send text messages, contacts,
	// locations and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	// *Optional*. *True*, if the user is allowed to send audios, documents, photos,
	// videos, video notes and voice notes, implies can_send_messages
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	// *Optional*. *True*, if the user is allowed to send polls, implies
	// can_send_messages
	CanSendPolls bool `json:"can_send_polls,omitempty"`
	// *Optional*. *True*, if the user is allowed to send animations, games, stickers
	// and use inline bots, implies can_send_media_messages
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	// *Optional*. *True*, if the user is allowed to add web page previews to their
	// messages, implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	// *Optional*. *True*, if the user is allowed to change the chat title, photo and
	// other settings. Ignored in public supergroups
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// *Optional*. *True*, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// *Optional*. *True*, if the user is allowed to pin messages. Ignored in public
	// supergroups
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// Represents a location to which a chat is connected.
type ChatLocation struct {
	// The location to which the supergroup is connected. Can't be a live location.
	Location *Location `json:"location"`
	// Location address; 1-64 characters, as defined by the chat owner
	Address string `json:"address"`
}

// This object represents a bot command.
type BotCommand struct {
	// Text of the command; 1-32 characters. Can contain only lowercase English
	// letters, digits and underscores.
	Command string `json:"command"`
	// Description of the command; 1-256 characters.
	Description string `json:"description"`
}

// Represents the default scope of bot commands. Default commands are used if no
// commands with a narrower scope are specified for the user.
type BotCommandScopeDefault struct {
	// Scope type, must be *default*
	Type string `json:"type,omitempty"`
}

// Represents the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	// Scope type, must be *all_private_chats*
	Type string `json:"type,omitempty"`
}

// Represents the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	// Scope type, must be *all_group_chats*
	Type string `json:"type,omitempty"`
}

// Represents the scope of bot commands, covering all group and supergroup chat
// administrators.
type BotCommandScopeAllChatAdministrators struct {
	// Scope type, must be *all_chat_administrators*
	Type string `json:"type,omitempty"`
}

// Represents the scope of bot commands, covering a specific chat.
type BotCommandScopeChat struct {
	// Scope type, must be *chat*
	Type string `json:"type,omitempty"`
	// Unique identifier for the target chat or username of the target supergroup (in
	// the format `@supergroupusername`)
	ChatId interface{} `json:"chat_id"`
}

// Represents the scope of bot commands, covering all administrators of a specific
// group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	// Scope type, must be *chat_administrators*
	Type string `json:"type,omitempty"`
	// Unique identifier for the target chat or username of the target supergroup (in
	// the format `@supergroupusername`)
	ChatId interface{} `json:"chat_id"`
}

// Represents the scope of bot commands, covering a specific member of a group or
// supergroup chat.
type BotCommandScopeChatMember struct {
	// Scope type, must be *chat_member*
	Type string `json:"type,omitempty"`
	// Unique identifier for the target chat or username of the target supergroup (in
	// the format `@supergroupusername`)
	ChatId interface{} `json:"chat_id"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id"`
}

// Represents a menu button, which opens the bot's list of commands.
type MenuButtonCommands struct {
	// Type of the button, must be *commands*
	Type string `json:"type,omitempty"`
}

// Represents a menu button, which launches a Web App.
type MenuButtonWebApp struct {
	// Type of the button, must be *web_app*
	Type string `json:"type,omitempty"`
	// Text on the button
	Text string `json:"text"`
	// Description of the Web App that will be launched when the user presses the
	// button. The Web App will be able to send an arbitrary message on behalf of the
	// user using the method answerWebAppQuery.
	WebApp *WebAppInfo `json:"web_app"`
}

// Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {
	// Type of the button, must be *default*
	Type string `json:"type,omitempty"`
}

// Contains information about why a request was unsuccessful.
type ResponseParameters struct {
	// *Optional*. The group has been migrated to a supergroup with the specified
	// identifier. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
	// *Optional*. In case of exceeding flood control, the number of seconds left to
	// wait before the request can be repeated
	RetryAfter int `json:"retry_after,omitempty"`
}

// Represents a photo to be sent.
type InputMediaPhoto struct {
	// Type of the result, must be *photo*
	Type string `json:"type,omitempty"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass "attach://<file_attach_name>" to upload a new one using
	// multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// *Optional*. Caption of the photo to be sent, 0-1024 characters after entities
	// parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the photo caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
}

// Represents a video to be sent.
type InputMediaVideo struct {
	// Type of the result, must be *video*
	Type string `json:"type,omitempty"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass "attach://<file_attach_name>" to upload a new one using
	// multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation
	// for the file is supported server-side. The thumbnail should be in JPEG format
	// and less than 200 kB in size. A thumbnail's width and height should not exceed
	// 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails
	// can't be reused and can be only uploaded as a new file, so you can pass
	// "attach://<file_attach_name>" if the thumbnail was uploaded using
	// multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb json.RawMessage `json:"thumb,omitempty"`
	// *Optional*. Caption of the video to be sent, 0-1024 characters after entities
	// parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the video caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Video width
	Width int `json:"width,omitempty"`
	// *Optional*. Video height
	Height int `json:"height,omitempty"`
	// *Optional*. Video duration in seconds
	Duration int `json:"duration,omitempty"`
	// *Optional*. Pass *True*, if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to
// be sent.
type InputMediaAnimation struct {
	// Type of the result, must be *animation*
	Type string `json:"type,omitempty"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass "attach://<file_attach_name>" to upload a new one using
	// multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation
	// for the file is supported server-side. The thumbnail should be in JPEG format
	// and less than 200 kB in size. A thumbnail's width and height should not exceed
	// 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails
	// can't be reused and can be only uploaded as a new file, so you can pass
	// "attach://<file_attach_name>" if the thumbnail was uploaded using
	// multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb json.RawMessage `json:"thumb,omitempty"`
	// *Optional*. Caption of the animation to be sent, 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the animation caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Animation width
	Width int `json:"width,omitempty"`
	// *Optional*. Animation height
	Height int `json:"height,omitempty"`
	// *Optional*. Animation duration in seconds
	Duration int `json:"duration,omitempty"`
}

// Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	// Type of the result, must be *audio*
	Type string `json:"type,omitempty"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass "attach://<file_attach_name>" to upload a new one using
	// multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation
	// for the file is supported server-side. The thumbnail should be in JPEG format
	// and less than 200 kB in size. A thumbnail's width and height should not exceed
	// 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails
	// can't be reused and can be only uploaded as a new file, so you can pass
	// "attach://<file_attach_name>" if the thumbnail was uploaded using
	// multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb json.RawMessage `json:"thumb,omitempty"`
	// *Optional*. Caption of the audio to be sent, 0-1024 characters after entities
	// parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the audio caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`
	// *Optional*. Performer of the audio
	Performer string `json:"performer,omitempty"`
	// *Optional*. Title of the audio
	Title string `json:"title,omitempty"`
}

// Represents a general file to be sent.
type InputMediaDocument struct {
	// Type of the result, must be *document*
	Type string `json:"type,omitempty"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers
	// (recommended), pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass "attach://<file_attach_name>" to upload a new one using
	// multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// *Optional*. Thumbnail of the file sent; can be ignored if thumbnail generation
	// for the file is supported server-side. The thumbnail should be in JPEG format
	// and less than 200 kB in size. A thumbnail's width and height should not exceed
	// 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails
	// can't be reused and can be only uploaded as a new file, so you can pass
	// "attach://<file_attach_name>" if the thumbnail was uploaded using
	// multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb json.RawMessage `json:"thumb,omitempty"`
	// *Optional*. Caption of the document to be sent, 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the document caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Disables automatic server-side content type detection for files
	// uploaded using multipart/form-data. Always *True*, if the document is sent as
	// part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}

// This object represents a sticker.
type Sticker struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and
	// for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Sticker width
	Width int `json:"width"`
	// Sticker height
	Height int `json:"height"`
	// *True*, if the sticker is animated
	IsAnimated bool `json:"is_animated,omitempty"`
	// *True*, if the sticker is a video sticker
	IsVideo bool `json:"is_video,omitempty"`
	// *Optional*. Sticker thumbnail in the .WEBP or .JPG format
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// *Optional*. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`
	// *Optional*. Name of the sticker set to which the sticker belongs
	SetName string `json:"set_name,omitempty"`
	// *Optional*. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// *Optional*. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// This object represents a sticker set.
type StickerSet struct {
	// Sticker set name
	Name string `json:"name"`
	// Sticker set title
	Title string `json:"title"`
	// *True*, if the sticker set contains animated stickers
	IsAnimated bool `json:"is_animated,omitempty"`
	// *True*, if the sticker set contains video stickers
	IsVideo bool `json:"is_video,omitempty"`
	// *True*, if the sticker set contains masks
	ContainsMasks bool `json:"contains_masks,omitempty"`
	// List of all set stickers
	Stickers []*Sticker `json:"stickers"`
	// *Optional*. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// This object describes the position on faces where a mask should be placed by
// default.
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed. One of
	// "forehead", "eyes", "mouth", or "chin".
	Point string `json:"point"`
	// Shift by X-axis measured in widths of the mask scaled to the face size, from
	// left to right. For example, choosing -1.0 will place mask just to the left of
	// the default mask position.
	XShift float32 `json:"x_shift"`
	// Shift by Y-axis measured in heights of the mask scaled to the face size, from
	// top to bottom. For example, 1.0 will place the mask just below the default mask
	// position.
	YShift float32 `json:"y_shift"`
	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float32 `json:"scale"`
}

// This object represents an incoming inline query. When the user sends an empty
// query, your bot could return some default or trending results.
type InlineQuery struct {
	// Unique identifier for this query
	Id string `json:"id"`
	// Sender
	From *User `json:"from"`
	// Text of the query (up to 256 characters)
	Query string `json:"query"`
	// Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`
	// *Optional*. Type of the chat, from which the inline query was sent. Can be
	// either "sender" for a private chat with the inline query sender, "private",
	// "group", "supergroup", or "channel". The chat type should be always known for
	// requests sent from official clients and most third-party clients, unless the
	// request was sent from a secret chat
	ChatType string `json:"chat_type,omitempty"`
	// *Optional*. Sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`
}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	// Type of the result, must be *article*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Title of the result
	Title string `json:"title"`
	// Content of the message to be sent
	InputMessageContent json.RawMessage `json:"input_message_content"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. URL of the result
	Url string `json:"url,omitempty"`
	// *Optional*. Pass *True*, if you don't want the URL to be shown in the message
	HideUrl bool `json:"hide_url,omitempty"`
	// *Optional*. Short description of the result
	Description string `json:"description,omitempty"`
	// *Optional*. Url of the thumbnail for the result
	ThumbUrl string `json:"thumb_url,omitempty"`
	// *Optional*. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// *Optional*. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// Represents a link to a photo. By default, this photo will be sent by the user
// with optional caption. Alternatively, you can use *input_message_content* to
// send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	// Type of the result, must be *photo*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL of the photo. Photo must be in **JPEG** format. Photo size must not
	// exceed 5MB
	PhotoUrl string `json:"photo_url"`
	// URL of the thumbnail for the photo
	ThumbUrl string `json:"thumb_url"`
	// *Optional*. Width of the photo
	PhotoWidth int `json:"photo_width,omitempty"`
	// *Optional*. Height of the photo
	PhotoHeight int `json:"photo_height,omitempty"`
	// *Optional*. Title for the result
	Title string `json:"title,omitempty"`
	// *Optional*. Short description of the result
	Description string `json:"description,omitempty"`
	// *Optional*. Caption of the photo to be sent, 0-1024 characters after entities
	// parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the photo caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the photo
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to an animated GIF file. By default, this animated GIF file
// will be sent by the user with optional caption. Alternatively, you can use
// *input_message_content* to send a message with the specified content instead of
// the animation.
type InlineQueryResultGif struct {
	// Type of the result, must be *gif*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the GIF file. File size must not exceed 1MB
	GifUrl string `json:"gif_url"`
	// *Optional*. Width of the GIF
	GifWidth int `json:"gif_width,omitempty"`
	// *Optional*. Height of the GIF
	GifHeight int `json:"gif_height,omitempty"`
	// *Optional*. Duration of the GIF in seconds
	GifDuration int `json:"gif_duration,omitempty"`
	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbUrl string `json:"thumb_url"`
	// *Optional*. MIME type of the thumbnail, must be one of "image/jpeg",
	// "image/gif", or "video/mp4". Defaults to "image/jpeg"
	ThumbMimeType string `json:"thumb_mime_type,omitempty"`
	// *Optional*. Title for the result
	Title string `json:"title,omitempty"`
	// *Optional*. Caption of the GIF file to be sent, 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the caption. See formatting options
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the GIF animation
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional
// caption. Alternatively, you can use *input_message_content* to send a message
// with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	// Type of the result, must be *mpeg4_gif*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the MP4 file. File size must not exceed 1MB
	Mpeg4Url string `json:"mpeg4_url"`
	// *Optional*. Video width
	Mpeg4Width int `json:"mpeg4_width,omitempty"`
	// *Optional*. Video height
	Mpeg4Height int `json:"mpeg4_height,omitempty"`
	// *Optional*. Video duration in seconds
	Mpeg4Duration int `json:"mpeg4_duration,omitempty"`
	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbUrl string `json:"thumb_url"`
	// *Optional*. MIME type of the thumbnail, must be one of "image/jpeg",
	// "image/gif", or "video/mp4". Defaults to "image/jpeg"
	ThumbMimeType string `json:"thumb_mime_type,omitempty"`
	// *Optional*. Title for the result
	Title string `json:"title,omitempty"`
	// *Optional*. Caption of the MPEG-4 file to be sent, 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the caption. See formatting options
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the video animation
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to a page containing an embedded video player or a video
// file. By default, this video file will be sent by the user with an optional
// caption. Alternatively, you can use *input_message_content* to send a message
// with the specified content instead of the video.
//
// If an InlineQueryResultVideo message contains an embedded video (e.g.,
// YouTube), you **must** replace its content using *input_message_content*.
type InlineQueryResultVideo struct {
	// Type of the result, must be *video*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the embedded video player or video file
	VideoUrl string `json:"video_url"`
	// Mime type of the content of video url, "text/html" or "video/mp4"
	MimeType string `json:"mime_type"`
	// URL of the thumbnail (JPEG only) for the video
	ThumbUrl string `json:"thumb_url"`
	// Title for the result
	Title string `json:"title"`
	// *Optional*. Caption of the video to be sent, 0-1024 characters after entities
	// parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the video caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Video width
	VideoWidth int `json:"video_width,omitempty"`
	// *Optional*. Video height
	VideoHeight int `json:"video_height,omitempty"`
	// *Optional*. Video duration in seconds
	VideoDuration int `json:"video_duration,omitempty"`
	// *Optional*. Short description of the result
	Description string `json:"description,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the video. This field
	// is **required** if InlineQueryResultVideo is used to send an HTML-page as a
	// result (e.g., a YouTube video).
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to an MP3 audio file. By default, this audio file will be
// sent by the user. Alternatively, you can use *input_message_content* to send a
// message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	// Type of the result, must be *audio*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the audio file
	AudioUrl string `json:"audio_url"`
	// Title
	Title string `json:"title"`
	// *Optional*. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the audio caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Performer
	Performer string `json:"performer,omitempty"`
	// *Optional*. Audio duration in seconds
	AudioDuration int `json:"audio_duration,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the audio
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to a voice recording in an .OGG container encoded with OPUS.
// By default, this voice recording will be sent by the user. Alternatively, you
// can use *input_message_content* to send a message with the specified content
// instead of the the voice message.
type InlineQueryResultVoice struct {
	// Type of the result, must be *voice*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the voice recording
	VoiceUrl string `json:"voice_url"`
	// Recording title
	Title string `json:"title"`
	// *Optional*. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the voice message caption. See
	// formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Recording duration in seconds
	VoiceDuration int `json:"voice_duration,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the voice recording
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to a file. By default, this file will be sent by the user
// with an optional caption. Alternatively, you can use *input_message_content* to
// send a message with the specified content instead of the file. Currently, only
// **.PDF** and **.ZIP** files can be sent using this method.
type InlineQueryResultDocument struct {
	// Type of the result, must be *document*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// Title for the result
	Title string `json:"title"`
	// *Optional*. Caption of the document to be sent, 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the document caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// A valid URL for the file
	DocumentUrl string `json:"document_url"`
	// Mime type of the content of the file, either "application/pdf" or
	// "application/zip"
	MimeType string `json:"mime_type"`
	// *Optional*. Short description of the result
	Description string `json:"description,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the file
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
	// *Optional*. URL of the thumbnail (JPEG only) for the file
	ThumbUrl string `json:"thumb_url,omitempty"`
	// *Optional*. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// *Optional*. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// Represents a location on a map. By default, the location will be sent by the
// user. Alternatively, you can use *input_message_content* to send a message with
// the specified content instead of the location.
type InlineQueryResultLocation struct {
	// Type of the result, must be *location*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Location latitude in degrees
	Latitude float32 `json:"latitude"`
	// Location longitude in degrees
	Longitude float32 `json:"longitude"`
	// Location title
	Title string `json:"title"`
	// *Optional*. The radius of uncertainty for the location, measured in meters;
	// 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`
	// *Optional*. Period in seconds for which the location can be updated, should be
	// between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`
	// *Optional*. For live locations, a direction in which the user is moving, in
	// degrees. Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`
	// *Optional*. For live locations, a maximum distance for proximity alerts about
	// approaching another chat member, in meters. Must be between 1 and 100000 if
	// specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the location
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
	// *Optional*. Url of the thumbnail for the result
	ThumbUrl string `json:"thumb_url,omitempty"`
	// *Optional*. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// *Optional*. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// Represents a venue. By default, the venue will be sent by the user.
// Alternatively, you can use *input_message_content* to send a message with the
// specified content instead of the venue.
type InlineQueryResultVenue struct {
	// Type of the result, must be *venue*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Latitude of the venue location in degrees
	Latitude float32 `json:"latitude"`
	// Longitude of the venue location in degrees
	Longitude float32 `json:"longitude"`
	// Title of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// *Optional*. Foursquare identifier of the venue if known
	FoursquareId string `json:"foursquare_id,omitempty"`
	// *Optional*. Foursquare type of the venue, if known. (For example,
	// "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	FoursquareType string `json:"foursquare_type,omitempty"`
	// *Optional*. Google Places identifier of the venue
	GooglePlaceId string `json:"google_place_id,omitempty"`
	// *Optional*. Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the venue
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
	// *Optional*. Url of the thumbnail for the result
	ThumbUrl string `json:"thumb_url,omitempty"`
	// *Optional*. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// *Optional*. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// Represents a contact with a phone number. By default, this contact will be sent
// by the user. Alternatively, you can use *input_message_content* to send a
// message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	// Type of the result, must be *contact*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// *Optional*. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// *Optional*. Additional data about the contact in the form of a vCard, 0-2048
	// bytes
	Vcard string `json:"vcard,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the contact
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
	// *Optional*. Url of the thumbnail for the result
	ThumbUrl string `json:"thumb_url,omitempty"`
	// *Optional*. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// *Optional*. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// Represents a Game.
type InlineQueryResultGame struct {
	// Type of the result, must be *game*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// Short name of the game
	GameShortName string `json:"game_short_name"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Represents a link to a photo stored on the Telegram servers. By default, this
// photo will be sent by the user with an optional caption. Alternatively, you can
// use *input_message_content* to send a message with the specified content
// instead of the photo.
type InlineQueryResultCachedPhoto struct {
	// Type of the result, must be *photo*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier of the photo
	PhotoFileId string `json:"photo_file_id"`
	// *Optional*. Title for the result
	Title string `json:"title,omitempty"`
	// *Optional*. Short description of the result
	Description string `json:"description,omitempty"`
	// *Optional*. Caption of the photo to be sent, 0-1024 characters after entities
	// parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the photo caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the photo
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to an animated GIF file stored on the Telegram servers. By
// default, this animated GIF file will be sent by the user with an optional
// caption. Alternatively, you can use *input_message_content* to send a message
// with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	// Type of the result, must be *gif*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the GIF file
	GifFileId string `json:"gif_file_id"`
	// *Optional*. Title for the result
	Title string `json:"title,omitempty"`
	// *Optional*. Caption of the GIF file to be sent, 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the caption. See formatting options
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the GIF animation
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound)
// stored on the Telegram servers. By default, this animated MPEG-4 file will be
// sent by the user with an optional caption. Alternatively, you can use
// *input_message_content* to send a message with the specified content instead of
// the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	// Type of the result, must be *mpeg4_gif*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the MP4 file
	Mpeg4FileId string `json:"mpeg4_file_id"`
	// *Optional*. Title for the result
	Title string `json:"title,omitempty"`
	// *Optional*. Caption of the MPEG-4 file to be sent, 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the caption. See formatting options
	// for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the video animation
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to a sticker stored on the Telegram servers. By default, this
// sticker will be sent by the user. Alternatively, you can use
// *input_message_content* to send a message with the specified content instead of
// the sticker.
type InlineQueryResultCachedSticker struct {
	// Type of the result, must be *sticker*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier of the sticker
	StickerFileId string `json:"sticker_file_id"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the sticker
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to a file stored on the Telegram servers. By default, this
// file will be sent by the user with an optional caption. Alternatively, you can
// use *input_message_content* to send a message with the specified content
// instead of the file.
type InlineQueryResultCachedDocument struct {
	// Type of the result, must be *document*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// Title for the result
	Title string `json:"title"`
	// A valid file identifier for the file
	DocumentFileId string `json:"document_file_id"`
	// *Optional*. Short description of the result
	Description string `json:"description,omitempty"`
	// *Optional*. Caption of the document to be sent, 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the document caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the file
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to a video file stored on the Telegram servers. By default,
// this video file will be sent by the user with an optional caption.
// Alternatively, you can use *input_message_content* to send a message with the
// specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	// Type of the result, must be *video*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the video file
	VideoFileId string `json:"video_file_id"`
	// Title for the result
	Title string `json:"title"`
	// *Optional*. Short description of the result
	Description string `json:"description,omitempty"`
	// *Optional*. Caption of the video to be sent, 0-1024 characters after entities
	// parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the video caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the video
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to a voice message stored on the Telegram servers. By
// default, this voice message will be sent by the user. Alternatively, you can
// use *input_message_content* to send a message with the specified content
// instead of the voice message.
type InlineQueryResultCachedVoice struct {
	// Type of the result, must be *voice*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the voice message
	VoiceFileId string `json:"voice_file_id"`
	// Voice message title
	Title string `json:"title"`
	// *Optional*. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the voice message caption. See
	// formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the voice message
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents a link to an MP3 audio file stored on the Telegram servers. By
// default, this audio file will be sent by the user. Alternatively, you can use
// *input_message_content* to send a message with the specified content instead of
// the audio.
type InlineQueryResultCachedAudio struct {
	// Type of the result, must be *audio*
	Type string `json:"type,omitempty"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the audio file
	AudioFileId string `json:"audio_file_id"`
	// *Optional*. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// *Optional*. Mode for parsing entities in the audio caption. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in the caption, which can be
	// specified instead of *parse_mode*
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// *Optional*. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// *Optional*. Content of the message to be sent instead of the audio
	InputMessageContent json.RawMessage `json:"input_message_content,omitempty"`
}

// Represents the content of a text message to be sent as the result of an inline
// query.
type InputTextMessageContent struct {
	// Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`
	// *Optional*. Mode for parsing entities in the message text. See formatting
	// options for more details.
	ParseMode string `json:"parse_mode,omitempty"`
	// *Optional*. List of special entities that appear in message text, which can be
	// specified instead of *parse_mode*
	Entities []*MessageEntity `json:"entities,omitempty"`
	// *Optional*. Disables link previews for links in the sent message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

// Represents the content of a location message to be sent as the result of an
// inline query.
type InputLocationMessageContent struct {
	// Latitude of the location in degrees
	Latitude float32 `json:"latitude"`
	// Longitude of the location in degrees
	Longitude float32 `json:"longitude"`
	// *Optional*. The radius of uncertainty for the location, measured in meters;
	// 0-1500
	HorizontalAccuracy float32 `json:"horizontal_accuracy,omitempty"`
	// *Optional*. Period in seconds for which the location can be updated, should be
	// between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`
	// *Optional*. For live locations, a direction in which the user is moving, in
	// degrees. Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`
	// *Optional*. For live locations, a maximum distance for proximity alerts about
	// approaching another chat member, in meters. Must be between 1 and 100000 if
	// specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

// Represents the content of a venue message to be sent as the result of an inline
// query.
type InputVenueMessageContent struct {
	// Latitude of the venue in degrees
	Latitude float32 `json:"latitude"`
	// Longitude of the venue in degrees
	Longitude float32 `json:"longitude"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// *Optional*. Foursquare identifier of the venue, if known
	FoursquareId string `json:"foursquare_id,omitempty"`
	// *Optional*. Foursquare type of the venue, if known. (For example,
	// "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	FoursquareType string `json:"foursquare_type,omitempty"`
	// *Optional*. Google Places identifier of the venue
	GooglePlaceId string `json:"google_place_id,omitempty"`
	// *Optional*. Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// Represents the content of a contact message to be sent as the result of an
// inline query.
type InputContactMessageContent struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// *Optional*. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// *Optional*. Additional data about the contact in the form of a vCard, 0-2048
	// bytes
	Vcard string `json:"vcard,omitempty"`
}

// Represents the content of an invoice message to be sent as the result of an
// inline query.
type InputInvoiceMessageContent struct {
	// Product name, 1-32 characters
	Title string `json:"title"`
	// Product description, 1-255 characters
	Description string `json:"description"`
	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the
	// user, use for your internal processes.
	Payload string `json:"payload"`
	// Payment provider token, obtained via Botfather
	ProviderToken string `json:"provider_token"`
	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency"`
	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax,
	// discount, delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices"`
	// *Optional*. The maximum accepted amount for tips in the *smallest units* of the
	// currency (integer, **not** float/double). For example, for a maximum tip of
	// `US$ 1.45` pass `max_tip_amount = 145`. See the *exp* parameter in
	// currencies.json, it shows the number of digits past the decimal point for each
	// currency (2 for the majority of currencies). Defaults to 0
	MaxTipAmount int `json:"max_tip_amount,omitempty"`
	// *Optional*. A JSON-serialized array of suggested amounts of tip in the
	// *smallest units* of the currency (integer, **not** float/double). At most 4
	// suggested tip amounts can be specified. The suggested tip amounts must be
	// positive, passed in a strictly increased order and must not exceed
	// *max_tip_amount*.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`
	// *Optional*. A JSON-serialized object for data about the invoice, which will be
	// shared with the payment provider. A detailed description of the required fields
	// should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`
	// *Optional*. URL of the product photo for the invoice. Can be a photo of the
	// goods or a marketing image for a service. People like it better when they see
	// what they are paying for.
	PhotoUrl string `json:"photo_url,omitempty"`
	// *Optional*. Photo size
	PhotoSize int `json:"photo_size,omitempty"`
	// *Optional*. Photo width
	PhotoWidth int `json:"photo_width,omitempty"`
	// *Optional*. Photo height
	PhotoHeight int `json:"photo_height,omitempty"`
	// *Optional*. Pass *True*, if you require the user's full name to complete the
	// order
	NeedName bool `json:"need_name,omitempty"`
	// *Optional*. Pass *True*, if you require the user's phone number to complete the
	// order
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
	// *Optional*. Pass *True*, if you require the user's email address to complete
	// the order
	NeedEmail bool `json:"need_email,omitempty"`
	// *Optional*. Pass *True*, if you require the user's shipping address to complete
	// the order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
	// *Optional*. Pass *True*, if user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
	// *Optional*. Pass *True*, if user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
	// *Optional*. Pass *True*, if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"`
}

// Represents a result of an inline query that was chosen by the user and sent to
// their chat partner.
type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ResultId string `json:"result_id"`
	// The user that chose the result
	From *User `json:"from"`
	// *Optional*. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`
	// *Optional*. Identifier of the sent inline message. Available only if there is
	// an inline keyboard attached to the message. Will be also received in callback
	// queries and can be used to edit the message.
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// The query that was used to obtain the result
	Query string `json:"query"`
}

// Contains information about an inline message sent by a Web App on behalf of a
// user.
type SentWebAppMessage struct {
	// *Optional*. Identifier of the sent inline message. Available only if there is
	// an inline keyboard attached to the message.
	InlineMessageId string `json:"inline_message_id,omitempty"`
}

// This object represents a portion of the price for goods or services.
type LabeledPrice struct {
	// Portion label
	Label string `json:"label"`
	// Price of the product in the *smallest units* of the currency (integer, **not**
	// float/double). For example, for a price of `US$ 1.45` pass `amount = 145`. See
	// the *exp* parameter in currencies.json, it shows the number of digits past the
	// decimal point for each currency (2 for the majority of currencies).
	Amount int `json:"amount,omitempty"`
}

// This object contains basic information about an invoice.
type Invoice struct {
	// Product name
	Title string `json:"title"`
	// Product description
	Description string `json:"description"`
	// Unique bot deep-linking parameter that can be used to generate this invoice
	StartParameter string `json:"start_parameter"`
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Total price in the *smallest units* of the currency (integer, **not**
	// float/double). For example, for a price of `US$ 1.45` pass `amount = 145`. See
	// the *exp* parameter in currencies.json, it shows the number of digits past the
	// decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount,omitempty"`
}

// This object represents a shipping address.
type ShippingAddress struct {
	// ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	// State, if applicable
	State string `json:"state"`
	// City
	City string `json:"city"`
	// First line for the address
	StreetLine1 string `json:"street_line1"`
	// Second line for the address
	StreetLine2 string `json:"street_line2"`
	// Address post code
	PostCode string `json:"post_code"`
}

// This object represents information about an order.
type OrderInfo struct {
	// *Optional*. User name
	Name string `json:"name,omitempty"`
	// *Optional*. User's phone number
	PhoneNumber string `json:"phone_number,omitempty"`
	// *Optional*. User email
	Email string `json:"email,omitempty"`
	// *Optional*. User shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// This object represents one shipping option.
type ShippingOption struct {
	// Shipping option identifier
	Id string `json:"id"`
	// Option title
	Title string `json:"title"`
	// List of price portions
	Prices []*LabeledPrice `json:"prices"`
}

// This object contains basic information about a successful payment.
type SuccessfulPayment struct {
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Total price in the *smallest units* of the currency (integer, **not**
	// float/double). For example, for a price of `US$ 1.45` pass `amount = 145`. See
	// the *exp* parameter in currencies.json, it shows the number of digits past the
	// decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount,omitempty"`
	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// *Optional*. Identifier of the shipping option chosen by the user
	ShippingOptionId string `json:"shipping_option_id,omitempty"`
	// *Optional*. Order info provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	// Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	// Provider payment identifier
	ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
}

// This object contains information about an incoming shipping query.
type ShippingQuery struct {
	// Unique query identifier
	Id string `json:"id"`
	// User who sent the query
	From *User `json:"from"`
	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// User specified shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	// Unique query identifier
	Id string `json:"id"`
	// User who sent the query
	From *User `json:"from"`
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Total price in the *smallest units* of the currency (integer, **not**
	// float/double). For example, for a price of `US$ 1.45` pass `amount = 145`. See
	// the *exp* parameter in currencies.json, it shows the number of digits past the
	// decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount,omitempty"`
	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// *Optional*. Identifier of the shipping option chosen by the user
	ShippingOptionId string `json:"shipping_option_id,omitempty"`
	// *Optional*. Order info provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

// Contains information about Telegram Passport data shared with the bot by the
// user.
type PassportData struct {
	// Array with information about documents and other Telegram Passport elements
	// that was shared with the bot
	Data []*EncryptedPassportElement `json:"data"`
	// Encrypted credentials required to decrypt the data
	Credentials *EncryptedCredentials `json:"credentials"`
}

// This object represents a file uploaded to Telegram Passport. Currently all
// Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and
	// for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// File size in bytes
	FileSize int `json:"file_size"`
	// Unix time when the file was uploaded
	FileDate int `json:"file_date"`
}

// Contains information about documents or other Telegram Passport elements shared
// with the bot by the user.
type EncryptedPassportElement struct {
	// Element type. One of "personal_details", "passport", "driver_license",
	// "identity_card", "internal_passport", "address", "utility_bill",
	// "bank_statement", "rental_agreement", "passport_registration",
	// "temporary_registration", "phone_number", "email".
	Type string `json:"type"`
	// *Optional*. Base64-encoded encrypted Telegram Passport element data provided by
	// the user, available for "personal_details", "passport", "driver_license",
	// "identity_card", "internal_passport" and "address" types. Can be decrypted and
	// verified using the accompanying EncryptedCredentials.
	Data string `json:"data,omitempty"`
	// *Optional*. User's verified phone number, available only for "phone_number" type
	PhoneNumber string `json:"phone_number,omitempty"`
	// *Optional*. User's verified email address, available only for "email" type
	Email string `json:"email,omitempty"`
	// *Optional*. Array of encrypted files with documents provided by the user,
	// available for "utility_bill", "bank_statement", "rental_agreement",
	// "passport_registration" and "temporary_registration" types. Files can be
	// decrypted and verified using the accompanying EncryptedCredentials.
	Files []*PassportFile `json:"files,omitempty"`
	// *Optional*. Encrypted file with the front side of the document, provided by the
	// user. Available for "passport", "driver_license", "identity_card" and
	// "internal_passport". The file can be decrypted and verified using the
	// accompanying EncryptedCredentials.
	FrontSide *PassportFile `json:"front_side,omitempty"`
	// *Optional*. Encrypted file with the reverse side of the document, provided by
	// the user. Available for "driver_license" and "identity_card". The file can be
	// decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`
	// *Optional*. Encrypted file with the selfie of the user holding a document,
	// provided by the user; available for "passport", "driver_license",
	// "identity_card" and "internal_passport". The file can be decrypted and verified
	// using the accompanying EncryptedCredentials.
	Selfie *PassportFile `json:"selfie,omitempty"`
	// *Optional*. Array of encrypted files with translated versions of documents
	// provided by the user. Available if requested for "passport", "driver_license",
	// "identity_card", "internal_passport", "utility_bill", "bank_statement",
	// "rental_agreement", "passport_registration" and "temporary_registration" types.
	// Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []*PassportFile `json:"translation,omitempty"`
	// Base64-encoded element hash for using in PassportElementErrorUnspecified
	Hash string `json:"hash"`
}

// Contains data required for decrypting and authenticating
// EncryptedPassportElement. See the Telegram Passport Documentation for a
// complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
	// Base64-encoded encrypted JSON-serialized data with unique user's payload, data
	// hashes and secrets required for EncryptedPassportElement decryption and
	// authentication
	Data string `json:"data"`
	// Base64-encoded data hash for data authentication
	Hash string `json:"hash"`
	// Base64-encoded secret, encrypted with the bot's public RSA key, required for
	// data decryption
	Secret string `json:"secret"`
}

// Represents an issue in one of the data fields that was provided by the user.
// The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	// Error source, must be *data*
	Source string `json:"source,omitempty"`
	// The section of the user's Telegram Passport which has the error, one of
	// "personal_details", "passport", "driver_license", "identity_card",
	// "internal_passport", "address"
	Type string `json:"type"`
	// Name of the data field which has the error
	FieldName string `json:"field_name"`
	// Base64-encoded data hash
	DataHash string `json:"data_hash"`
	// Error message
	Message string `json:"message"`
}

// Represents an issue with the front side of a document. The error is considered
// resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	// Error source, must be *front_side*
	Source string `json:"source,omitempty"`
	// The section of the user's Telegram Passport which has the issue, one of
	// "passport", "driver_license", "identity_card", "internal_passport"
	Type string `json:"type"`
	// Base64-encoded hash of the file with the front side of the document
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

// Represents an issue with the reverse side of a document. The error is
// considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	// Error source, must be *reverse_side*
	Source string `json:"source,omitempty"`
	// The section of the user's Telegram Passport which has the issue, one of
	// "driver_license", "identity_card"
	Type string `json:"type"`
	// Base64-encoded hash of the file with the reverse side of the document
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

// Represents an issue with the selfie with a document. The error is considered
// resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	// Error source, must be *selfie*
	Source string `json:"source,omitempty"`
	// The section of the user's Telegram Passport which has the issue, one of
	// "passport", "driver_license", "identity_card", "internal_passport"
	Type string `json:"type"`
	// Base64-encoded hash of the file with the selfie
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

// Represents an issue with a document scan. The error is considered resolved when
// the file with the document scan changes.
type PassportElementErrorFile struct {
	// Error source, must be *file*
	Source string `json:"source,omitempty"`
	// The section of the user's Telegram Passport which has the issue, one of
	// "utility_bill", "bank_statement", "rental_agreement", "passport_registration",
	// "temporary_registration"
	Type string `json:"type"`
	// Base64-encoded file hash
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

// Represents an issue with a list of scans. The error is considered resolved when
// the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	// Error source, must be *files*
	Source string `json:"source,omitempty"`
	// The section of the user's Telegram Passport which has the issue, one of
	// "utility_bill", "bank_statement", "rental_agreement", "passport_registration",
	// "temporary_registration"
	Type string `json:"type"`
	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`
	// Error message
	Message string `json:"message"`
}

// Represents an issue with one of the files that constitute the translation of a
// document. The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	// Error source, must be *translation_file*
	Source string `json:"source,omitempty"`
	// Type of element of the user's Telegram Passport which has the issue, one of
	// "passport", "driver_license", "identity_card", "internal_passport",
	// "utility_bill", "bank_statement", "rental_agreement", "passport_registration",
	// "temporary_registration"
	Type string `json:"type"`
	// Base64-encoded file hash
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

// Represents an issue with the translated version of a document. The error is
// considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	// Error source, must be *translation_files*
	Source string `json:"source,omitempty"`
	// Type of element of the user's Telegram Passport which has the issue, one of
	// "passport", "driver_license", "identity_card", "internal_passport",
	// "utility_bill", "bank_statement", "rental_agreement", "passport_registration",
	// "temporary_registration"
	Type string `json:"type"`
	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`
	// Error message
	Message string `json:"message"`
}

// Represents an issue in an unspecified place. The error is considered resolved
// when new data is added.
type PassportElementErrorUnspecified struct {
	// Error source, must be *unspecified*
	Source string `json:"source,omitempty"`
	// Type of element of the user's Telegram Passport which has the issue
	Type string `json:"type"`
	// Base64-encoded element hash
	ElementHash string `json:"element_hash"`
	// Error message
	Message string `json:"message"`
}

// This object represents a game. Use BotFather to create and edit games, their
// short names will act as unique identifiers.
type Game struct {
	// Title of the game
	Title string `json:"title"`
	// Description of the game
	Description string `json:"description"`
	// Photo that will be displayed in the game message in chats.
	Photo []*PhotoSize `json:"photo"`
	// *Optional*. Brief description of the game or high scores included in the game
	// message. Can be automatically edited to include current high scores for the
	// game when the bot calls setGameScore, or manually edited using editMessageText.
	// 0-4096 characters.
	Text string `json:"text,omitempty"`
	// *Optional*. Special entities that appear in *text*, such as usernames, URLs,
	// bot commands, etc.
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	// *Optional*. Animation that will be displayed in the game message in chats.
	// Upload via BotFather
	Animation *Animation `json:"animation,omitempty"`
}

// This object represents one row of the high scores table for a game.
type GameHighScore struct {
	// Position in high score table for the game
	Position int `json:"position"`
	// User
	User *User `json:"user"`
	// Score
	Score int `json:"score"`
}