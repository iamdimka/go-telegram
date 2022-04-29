package telegram

import "encoding/json"

func (b *Bot) GetUpdates(request *GetUpdatesRequest) (result []*Update, err error) {
	err = b.request("getUpdates", request, &result)
	return
}

func (b *Bot) SetWebhook(request *SetWebhookRequest) (result bool, err error) {
	err = b.request("setWebhook", request, &result)
	return
}

func (b *Bot) DeleteWebhook(request *DeleteWebhookRequest) (result bool, err error) {
	err = b.request("deleteWebhook", request, &result)
	return
}

func (b *Bot) GetWebhookInfo() (result *WebhookInfo, err error) {
	err = b.request("getWebhookInfo", nil, result)
	return
}

func (b *Bot) GetMe() (result *User, err error) {
	err = b.request("getMe", nil, result)
	return
}

func (b *Bot) LogOut() (result bool, err error) {
	err = b.request("logOut", nil, &result)
	return
}

func (b *Bot) Close() (result bool, err error) {
	err = b.request("close", nil, &result)
	return
}

func (b *Bot) SendMessage(request *SendMessageRequest) (result *Message, err error) {
	err = b.request("sendMessage", request, result)
	return
}

func (b *Bot) ForwardMessage(request *ForwardMessageRequest) (result *Message, err error) {
	err = b.request("forwardMessage", request, result)
	return
}

func (b *Bot) CopyMessage(request *CopyMessageRequest) (result *MessageId, err error) {
	err = b.request("copyMessage", request, result)
	return
}

func (b *Bot) SendPhoto(request *SendPhotoRequest) (result *Message, err error) {
	err = b.request("sendPhoto", request, result)
	return
}

func (b *Bot) SendAudio(request *SendAudioRequest) (result *Message, err error) {
	err = b.request("sendAudio", request, result)
	return
}

func (b *Bot) SendDocument(request *SendDocumentRequest) (result *Message, err error) {
	err = b.request("sendDocument", request, result)
	return
}

func (b *Bot) SendVideo(request *SendVideoRequest) (result *Message, err error) {
	err = b.request("sendVideo", request, result)
	return
}

func (b *Bot) SendAnimation(request *SendAnimationRequest) (result *Message, err error) {
	err = b.request("sendAnimation", request, result)
	return
}

func (b *Bot) SendVoice(request *SendVoiceRequest) (result *Message, err error) {
	err = b.request("sendVoice", request, result)
	return
}

func (b *Bot) SendVideoNote(request *SendVideoNoteRequest) (result *Message, err error) {
	err = b.request("sendVideoNote", request, result)
	return
}

func (b *Bot) SendMediaGroup(request *SendMediaGroupRequest) (result []json.RawMessage, err error) {
	err = b.request("sendMediaGroup", request, &result)
	return
}

func (b *Bot) SendLocation(request *SendLocationRequest) (result *Message, err error) {
	err = b.request("sendLocation", request, result)
	return
}

func (b *Bot) EditMessageLiveLocation(request *EditMessageLiveLocationRequest) (result *Message, err error) {
	err = b.request("editMessageLiveLocation", request, result)
	return
}

func (b *Bot) StopMessageLiveLocation(request *StopMessageLiveLocationRequest) (result *Message, err error) {
	err = b.request("stopMessageLiveLocation", request, result)
	return
}

func (b *Bot) SendVenue(request *SendVenueRequest) (result *Message, err error) {
	err = b.request("sendVenue", request, result)
	return
}

func (b *Bot) SendContact(request *SendContactRequest) (result *Message, err error) {
	err = b.request("sendContact", request, result)
	return
}

func (b *Bot) SendPoll(request *SendPollRequest) (result *Message, err error) {
	err = b.request("sendPoll", request, result)
	return
}

func (b *Bot) SendDice(request *SendDiceRequest) (result *Message, err error) {
	err = b.request("sendDice", request, result)
	return
}

func (b *Bot) SendChatAction(request *SendChatActionRequest) (result bool, err error) {
	err = b.request("sendChatAction", request, &result)
	return
}

func (b *Bot) GetUserProfilePhotos(request *GetUserProfilePhotosRequest) (result *UserProfilePhotos, err error) {
	err = b.request("getUserProfilePhotos", request, result)
	return
}

func (b *Bot) GetFile(request *GetFileRequest) (result *File, err error) {
	err = b.request("getFile", request, result)
	return
}

func (b *Bot) BanChatMember(request *BanChatMemberRequest) (result bool, err error) {
	err = b.request("banChatMember", request, &result)
	return
}

func (b *Bot) UnbanChatMember(request *UnbanChatMemberRequest) (result bool, err error) {
	err = b.request("unbanChatMember", request, &result)
	return
}

func (b *Bot) RestrictChatMember(request *RestrictChatMemberRequest) (result bool, err error) {
	err = b.request("restrictChatMember", request, &result)
	return
}

func (b *Bot) PromoteChatMember(request *PromoteChatMemberRequest) (result bool, err error) {
	err = b.request("promoteChatMember", request, &result)
	return
}

func (b *Bot) SetChatAdministratorCustomTitle(request *SetChatAdministratorCustomTitleRequest) (result bool, err error) {
	err = b.request("setChatAdministratorCustomTitle", request, &result)
	return
}

func (b *Bot) BanChatSenderChat(request *BanChatSenderChatRequest) (result bool, err error) {
	err = b.request("banChatSenderChat", request, &result)
	return
}

func (b *Bot) UnbanChatSenderChat(request *UnbanChatSenderChatRequest) (result bool, err error) {
	err = b.request("unbanChatSenderChat", request, &result)
	return
}

func (b *Bot) SetChatPermissions(request *SetChatPermissionsRequest) (result bool, err error) {
	err = b.request("setChatPermissions", request, &result)
	return
}

func (b *Bot) ExportChatInviteLink(request *ExportChatInviteLinkRequest) (result string, err error) {
	err = b.request("exportChatInviteLink", request, &result)
	return
}

func (b *Bot) CreateChatInviteLink(request *CreateChatInviteLinkRequest) (result *ChatInviteLink, err error) {
	err = b.request("createChatInviteLink", request, result)
	return
}

func (b *Bot) EditChatInviteLink(request *EditChatInviteLinkRequest) (result *ChatInviteLink, err error) {
	err = b.request("editChatInviteLink", request, result)
	return
}

func (b *Bot) RevokeChatInviteLink(request *RevokeChatInviteLinkRequest) (result *ChatInviteLink, err error) {
	err = b.request("revokeChatInviteLink", request, result)
	return
}

func (b *Bot) ApproveChatJoinRequest(request *ApproveChatJoinRequestRequest) (result bool, err error) {
	err = b.request("approveChatJoinRequest", request, &result)
	return
}

func (b *Bot) DeclineChatJoinRequest(request *DeclineChatJoinRequestRequest) (result bool, err error) {
	err = b.request("declineChatJoinRequest", request, &result)
	return
}

func (b *Bot) SetChatPhoto(request *SetChatPhotoRequest) (result bool, err error) {
	err = b.request("setChatPhoto", request, &result)
	return
}

func (b *Bot) DeleteChatPhoto(request *DeleteChatPhotoRequest) (result bool, err error) {
	err = b.request("deleteChatPhoto", request, &result)
	return
}

func (b *Bot) SetChatTitle(request *SetChatTitleRequest) (result bool, err error) {
	err = b.request("setChatTitle", request, &result)
	return
}

func (b *Bot) SetChatDescription(request *SetChatDescriptionRequest) (result bool, err error) {
	err = b.request("setChatDescription", request, &result)
	return
}

func (b *Bot) PinChatMessage(request *PinChatMessageRequest) (result bool, err error) {
	err = b.request("pinChatMessage", request, &result)
	return
}

func (b *Bot) UnpinChatMessage(request *UnpinChatMessageRequest) (result bool, err error) {
	err = b.request("unpinChatMessage", request, &result)
	return
}

func (b *Bot) UnpinAllChatMessages(request *UnpinAllChatMessagesRequest) (result bool, err error) {
	err = b.request("unpinAllChatMessages", request, &result)
	return
}

func (b *Bot) LeaveChat(request *LeaveChatRequest) (result bool, err error) {
	err = b.request("leaveChat", request, &result)
	return
}

func (b *Bot) GetChat(request *GetChatRequest) (result *Chat, err error) {
	err = b.request("getChat", request, result)
	return
}

func (b *Bot) GetChatAdministrators(request *GetChatAdministratorsRequest) (result []json.RawMessage, err error) {
	err = b.request("getChatAdministrators", request, &result)
	return
}

func (b *Bot) GetChatMemberCount(request *GetChatMemberCountRequest) (result int, err error) {
	err = b.request("getChatMemberCount", request, &result)
	return
}

func (b *Bot) GetChatMember(request *GetChatMemberRequest) (result json.RawMessage, err error) {
	err = b.request("getChatMember", request, &result)
	return
}

func (b *Bot) SetChatStickerSet(request *SetChatStickerSetRequest) (result bool, err error) {
	err = b.request("setChatStickerSet", request, &result)
	return
}

func (b *Bot) DeleteChatStickerSet(request *DeleteChatStickerSetRequest) (result bool, err error) {
	err = b.request("deleteChatStickerSet", request, &result)
	return
}

func (b *Bot) AnswerCallbackQuery(request *AnswerCallbackQueryRequest) (result bool, err error) {
	err = b.request("answerCallbackQuery", request, &result)
	return
}

func (b *Bot) SetMyCommands(request *SetMyCommandsRequest) (result bool, err error) {
	err = b.request("setMyCommands", request, &result)
	return
}

func (b *Bot) DeleteMyCommands(request *DeleteMyCommandsRequest) (result bool, err error) {
	err = b.request("deleteMyCommands", request, &result)
	return
}

func (b *Bot) GetMyCommands(request *GetMyCommandsRequest) (result []*BotCommand, err error) {
	err = b.request("getMyCommands", request, &result)
	return
}

func (b *Bot) SetChatMenuButton(request *SetChatMenuButtonRequest) (result bool, err error) {
	err = b.request("setChatMenuButton", request, &result)
	return
}

func (b *Bot) GetChatMenuButton(request *GetChatMenuButtonRequest) (result json.RawMessage, err error) {
	err = b.request("getChatMenuButton", request, &result)
	return
}

func (b *Bot) SetMyDefaultAdministratorRights(request *SetMyDefaultAdministratorRightsRequest) (result bool, err error) {
	err = b.request("setMyDefaultAdministratorRights", request, &result)
	return
}

func (b *Bot) GetMyDefaultAdministratorRights(request *GetMyDefaultAdministratorRightsRequest) (result *ChatAdministratorRights, err error) {
	err = b.request("getMyDefaultAdministratorRights", request, result)
	return
}

func (b *Bot) EditMessageText(request *EditMessageTextRequest) (result *Message, err error) {
	err = b.request("editMessageText", request, result)
	return
}

func (b *Bot) EditMessageCaption(request *EditMessageCaptionRequest) (result *Message, err error) {
	err = b.request("editMessageCaption", request, result)
	return
}

func (b *Bot) EditMessageMedia(request *EditMessageMediaRequest) (result *Message, err error) {
	err = b.request("editMessageMedia", request, result)
	return
}

func (b *Bot) EditMessageReplyMarkup(request *EditMessageReplyMarkupRequest) (result *Message, err error) {
	err = b.request("editMessageReplyMarkup", request, result)
	return
}

func (b *Bot) StopPoll(request *StopPollRequest) (result *Poll, err error) {
	err = b.request("stopPoll", request, result)
	return
}

func (b *Bot) DeleteMessage(request *DeleteMessageRequest) (result bool, err error) {
	err = b.request("deleteMessage", request, &result)
	return
}

func (b *Bot) SendSticker(request *SendStickerRequest) (result *Message, err error) {
	err = b.request("sendSticker", request, result)
	return
}

func (b *Bot) GetStickerSet(request *GetStickerSetRequest) (result *StickerSet, err error) {
	err = b.request("getStickerSet", request, result)
	return
}

func (b *Bot) UploadStickerFile(request *UploadStickerFileRequest) (result *File, err error) {
	err = b.request("uploadStickerFile", request, result)
	return
}

func (b *Bot) CreateNewStickerSet(request *CreateNewStickerSetRequest) (result bool, err error) {
	err = b.request("createNewStickerSet", request, &result)
	return
}

func (b *Bot) AddStickerToSet(request *AddStickerToSetRequest) (result bool, err error) {
	err = b.request("addStickerToSet", request, &result)
	return
}

func (b *Bot) SetStickerPositionInSet(request *SetStickerPositionInSetRequest) (result bool, err error) {
	err = b.request("setStickerPositionInSet", request, &result)
	return
}

func (b *Bot) DeleteStickerFromSet(request *DeleteStickerFromSetRequest) (result bool, err error) {
	err = b.request("deleteStickerFromSet", request, &result)
	return
}

func (b *Bot) SetStickerSetThumb(request *SetStickerSetThumbRequest) (result bool, err error) {
	err = b.request("setStickerSetThumb", request, &result)
	return
}

func (b *Bot) AnswerInlineQuery(request *AnswerInlineQueryRequest) (result bool, err error) {
	err = b.request("answerInlineQuery", request, &result)
	return
}

func (b *Bot) AnswerWebAppQuery(request *AnswerWebAppQueryRequest) (result *SentWebAppMessage, err error) {
	err = b.request("answerWebAppQuery", request, result)
	return
}

func (b *Bot) SendInvoice(request *SendInvoiceRequest) (result *Message, err error) {
	err = b.request("sendInvoice", request, result)
	return
}

func (b *Bot) AnswerShippingQuery(request *AnswerShippingQueryRequest) (result bool, err error) {
	err = b.request("answerShippingQuery", request, &result)
	return
}

func (b *Bot) AnswerPreCheckoutQuery(request *AnswerPreCheckoutQueryRequest) (result bool, err error) {
	err = b.request("answerPreCheckoutQuery", request, &result)
	return
}

func (b *Bot) SetPassportDataErrors(request *SetPassportDataErrorsRequest) (result bool, err error) {
	err = b.request("setPassportDataErrors", request, &result)
	return
}

func (b *Bot) SendGame(request *SendGameRequest) (result *Message, err error) {
	err = b.request("sendGame", request, result)
	return
}

func (b *Bot) SetGameScore(request *SetGameScoreRequest) (result *Message, err error) {
	err = b.request("setGameScore", request, result)
	return
}

func (b *Bot) GetGameHighScores(request *GetGameHighScoresRequest) (result []*GameHighScore, err error) {
	err = b.request("getGameHighScores", request, &result)
	return
}