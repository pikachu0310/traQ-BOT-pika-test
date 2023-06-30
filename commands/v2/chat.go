package v2

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"

	"github.com/sashabaranov/go-openai"

	"example-bot/api"
	"example-bot/util"
)

var apiKey = util.GetApiKey()

type Models int

const (
	GPT3dot5Turbo Models = iota
	GPT4
)

type FinishReason int

const (
	stop FinishReason = iota
	length
	errorHappen
)

var blobs = [...]string{":blob_bongo:", ":blob_crazy_happy:", ":blob_grin:", ":blob_hype:", ":blob_love:", ":blob_lurk:", ":blob_pyon:", ":blob_pyon_inverse:", ":blob_slide:", ":blob_snowball_1:", ":blob_snowball_2:", ":blob_speedy_roll:", ":blob_speedy_roll_inverse:", ":blob_thinking:", ":blob_thinking_fast:", ":blob_thinking_portal:", ":blob_thinking_upsidedown:", ":blob_thonkang:", ":blob_thumbs_up:", ":blobblewobble:", ":blobenjoy:", ":blobglitch:", ":blobbass:", ":blobjam:", ":blobkeyboard:", ":bloblamp:", ":blobmaracas:", ":blobmicrophone:", ":blobthinksmart:", ":blobwobwork:", ":conga_party_thinking_blob:", ":Hyperblob:", ":party_blob:", ":partyparrot_blob:", ":partyparrot_blob_cat:"}
var amazed = [...]string{":amazed_fuzzy:", ":amazed_amazed_fuzzy:", ":amazed_god_enel:", ":amazed_hamutaro:"}
var blobsAndAmazed = append(blobs[:], amazed[:]...)
var warnings = [...]string{":warning:", ":ikura-hamu_shooting_warning:"}

type Message = openai.ChatCompletionMessage

var Messages = make([]Message, 0)
var SystemRoleMessage = "あなたはサークルである東京工業大学デジタル創作同好会traPの部内SNS、traQのユーザーを楽しませる娯楽用途のBOTの中に作られた、openaiのモデルgpt-3.5-turboを用いた対話型AIです。身内しかいないSNSで、ユーザーに緩く接してください。そして、ユーザーの言う事に出来る限り従うようにしてください。"

const GptSystemString = "FirstSystemMessageを変更しました。/gptsys showで確認できます。\nFirstSystemMessageとは、常に履歴の一番最初に入り、最初にgptに情報や状況を説明するのに使用する文字列です"

// func main() {
// 	addMessageAsUser("hello!")
// 	OpenAIStream(Messages, GPT3dot5Turbo, func(responseMessage string) {
// 		fmt.Println(responseMessage)
// 	})
// 	addMessageAsUser("I'm hungry!")
// 	OpenAIStream(Messages, GPT3dot5Turbo, func(responseMessage string) {
// 		fmt.Println(responseMessage)
// 	})
// }

func OpenAIStream(messages []Message, openaiModel Models, do func(string)) (finishReason FinishReason, err error) {
	c := openai.NewClient(apiKey)
	ctx := context.Background()

	var model string
	switch openaiModel {
	case GPT3dot5Turbo:
		model = openai.GPT3Dot5Turbo
	case GPT4:
		model = openai.GPT4
	default:
		model = openai.GPT3Dot5Turbo
	}

	req := openai.ChatCompletionRequest{
		Model:    model,
		Messages: messages,
		Stream:   true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	var responseMessage string
	for {
		response, err := stream.Recv()

		if err != nil {
			do(responseMessage + warnings[rand.Intn(len(warnings))] + ":blobglitch: Error: " + fmt.Sprint(err))
			finishReason = errorHappen
			break
		}
		if errors.Is(err, io.EOF) {
			err = errors.New("stream closed")
			finishReason = errorHappen
			break
		}

		if response.Choices[0].FinishReason == "stop" {
			do(responseMessage)
			finishReason = stop
			break
		} else if response.Choices[0].FinishReason == "length" {
			do(responseMessage + "\n" + amazed[rand.Intn(len(amazed))] + "トークン(履歴を含む文字数)が上限に達したので履歴の最初のメッセージを削除して続きを出力します:loading:")
			finishReason = length
			break
		}

		responseMessage += response.Choices[0].Delta.Content
		do(blobsAndAmazed[rand.Intn(len(blobs))] + responseMessage + ":loading:")
	}
	addMessageAsAssistant(responseMessage)
	return
}

func Chat(channelID, newMessage string, openaiModel Models) {
	addMessageAsUser(newMessage)
	updateSystemRoleMessage(SystemRoleMessage)
	postMessage, err := api.PostMessageWithErr(channelID, blobs[rand.Intn(len(blobs))]+":loading:")
	if err != nil {
		fmt.Println(err)
	}
	finishReason, err := OpenAIStream(Messages, openaiModel, func(responseMessage string) {
		api.EditMessage(postMessage.Id, responseMessage)
	})
	if err != nil {
		fmt.Println(err)
	}

	// finishReasonがlength以外になるまで、最大5回まで繰り返す
	for i := 0; i < 5 && finishReason == length; i++ {
		nowPostMessage := postMessage.Content
		Messages = Messages[1:]
		addMessageAsUser("先ほどのあなたのメッセージが途中で途切れてしまっているので、続きだけを出力してください。")
		finishReason, err = OpenAIStream(Messages, openaiModel, func(responseMessage string) {
			api.EditMessage(postMessage.Id, nowPostMessage+"\n"+responseMessage)
		})
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func ChatChangeSystemMessage(channelID, message string) {
	SystemRoleMessage = message
	api.PostMessage(channelID, GptSystemString)
}

func ChatShowSystemMessage(channelID string) {
	api.PostMessage(channelID, SystemRoleMessage)
}

func ChatReset(channelID string) {
	msg := api.PostMessage(channelID, ":blobnom::loading:")
	Messages = make([]Message, 0)
	err := api.EditMessageWithErr(msg.Id, ":done:")
	if err != nil {
		api.EditMessage(msg.Id, "Error: "+fmt.Sprint(err))
	}
	return
}

func addMessageAsUser(message string) {
	Messages = append(Messages, Message{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
}

func addMessageAsAssistant(message string) {
	Messages = append(Messages, Message{
		Role:    openai.ChatMessageRoleAssistant,
		Content: message,
	})
}

func addMessageAsSystem(message string) {
	Messages = append(Messages, Message{
		Role:    openai.ChatMessageRoleSystem,
		Content: message,
	})
}

func addSystemMessageIfNotExist(message string) {
	for _, m := range Messages {
		if m.Role == "system" {
			return
		}
	}
	Messages = append([]Message{{
		Role:    openai.ChatMessageRoleSystem,
		Content: message,
	}}, Messages...)
}

func updateSystemRoleMessage(message string) {
	addSystemMessageIfNotExist(message)
	Messages[0] = Message{
		Role:    "system",
		Content: message,
	}
}
