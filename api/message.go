package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"example-bot/util"

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

	res, err := bot.API().
		MessageApi.EditMessage(context.Background(), messageID).PostMessageRequest(traq.PostMessageRequest{
		Content: content,
	}).Execute()
	if err != nil {
		res2, err2 := io.ReadAll(res.Body)
		if err2 != nil {
			return err2
		}
		return fmt.Errorf("%w: %s", err, string(res2))
	}
	return nil
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

func GetMessagesFromUser(userID string, limit int, offset int, before time.Time) (*traq.MessageSearchResult, error) {
	bot := util.GetBot()

	messages, res, err := bot.API().
		MessageApi.
		SearchMessages(context.Background()).From(userID).Limit(int32(limit)).Offset(int32(offset)).Before(before).
		Execute()
	if err != nil {
		res2, err2 := io.ReadAll(res.Body)
		if err2 != nil {
			return nil, err2
		}
		return nil, fmt.Errorf("%w: %s", err, string(res2))
	}
	return messages, err
}

func GetMessagesFromPeriod(after time.Time, before time.Time, limit int, offset int) (*traq.MessageSearchResult, error) {
	bot := util.GetBot()

	messages, res, err := bot.API().
		MessageApi.
		SearchMessages(context.Background()).Limit(int32(limit)).Offset(int32(offset)).Before(before).After(after).
		Execute()
	if err != nil {
		res2, err2 := io.ReadAll(res.Body)
		if err2 != nil {
			return nil, err2
		}
		return nil, fmt.Errorf("%w: %s", err, string(res2))
	}
	return messages, err
}

func DeleteMessage(messageID string) error {
	bot := util.GetBot()

	res, err := bot.API().
		MessageApi.DeleteMessage(context.Background(), messageID).
		Execute()
	if err != nil {
		res2, err2 := io.ReadAll(res.Body)
		if err2 != nil {
			return err2
		}
		return fmt.Errorf("%w: %s", err, string(res2))
	}
	return err
}
