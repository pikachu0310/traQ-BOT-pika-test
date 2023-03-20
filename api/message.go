package api

import (
	"context"
	"example-bot/util"
	"log"

	"github.com/traPtitech/go-traq"
)

type TraqMessage traq.Message

func (message *TraqMessage) Edit(content string) error {
	bot := util.GetBot()

	_, err := bot.API().
		MessageApi.EditMessage(context.Background(), message.Id).PostMessageRequest(traq.PostMessageRequest{
		Content: content,
	}).Execute()
	return err
}

func PostMessage(channelID string, content string) *traq.Message {

	bot := util.GetBot()

	message, _, err := bot.API().
		MessageApi.
		PostMessage(context.Background(), channelID).
		PostMessageRequest(traq.PostMessageRequest{
			Content: content,
		}).
		Execute()
	if err != nil {
		log.Println(err)
	}
	return message
}

func PostMessageWithErr(channelID string, content string) (*traq.Message, error) {

	bot := util.GetBot()

	message, _, err := bot.API().
		MessageApi.
		PostMessage(context.Background(), channelID).
		PostMessageRequest(traq.PostMessageRequest{
			Content: content,
		}).
		Execute()
	if err != nil {
		log.Println(err)
	}
	return message, err
}

func EditMessage(messageID string, content string) {

	bot := util.GetBot()

	_, err := bot.API().
		MessageApi.EditMessage(context.Background(), messageID).PostMessageRequest(traq.PostMessageRequest{
		Content: content,
	}).Execute()
	if err != nil {
		log.Println(err)
	}
}

func EditMessageWithErr(messageID string, content string) error {

	bot := util.GetBot()

	_, err := bot.API().
		MessageApi.EditMessage(context.Background(), messageID).PostMessageRequest(traq.PostMessageRequest{
		Content: content,
	}).Execute()
	return err
}

func GetMessage(messageID string) *traq.Message {

	bot := util.GetBot()

	message, _, err := bot.API().
		MessageApi.
		GetMessage(context.Background(), messageID).
		Execute()
	if err != nil {
		log.Println(err)
	}
	return message
}

func AddMessage(messageId string, content string) {
	messageContent := GetMessage(messageId).Content
	EditMessage(messageId, messageContent+content)
}

func AddMessageWithNewLine(messageId string, content string) {
	messageContent := GetMessage(messageId).Content
	EditMessage(messageId, messageContent+"\n"+content)
}

func GetMessages(text string) *traq.MessageSearchResult {

	bot := util.GetBot()

	messages, _, err := bot.API().
		MessageApi.
		SearchMessages(context.Background()).Word(text).
		Execute()
	if err != nil {
		log.Println(err)
	}
	return messages
}
