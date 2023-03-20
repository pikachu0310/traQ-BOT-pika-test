package commands

import (
	"bytes"
	"encoding/json"
	"errors"
	"example-bot/api"
	"example-bot/util"
	"fmt"
	"github.com/traPtitech/go-traq"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
)

type OpenaiRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenaiResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

var apiKey = util.GetApiKey()

var (
	JsonError  = errors.New("Error:invalid character '<' looking for beginning of value")
	TokenError = errors.New("Error: TokenOver")
)

const model = "gpt-3.5-turbo"
const openaiURL = "https://api.openai.com/v1/chat/completions"

var blobs = [...]string{":blob_bongo:", ":blob_crazy_happy:", ":blob_grin:", ":blob_hype:", ":blob_love:", ":blob_lurk:", ":blob_pyon:", ":blob_pyon_inverse:", ":blob_slide:", ":blob_snowball_1:", ":blob_snowball_2:", ":blob_speedy_roll:", ":blob_speedy_roll_inverse:", ":blob_thinking:", ":blob_thinking_fast:", ":blob_thinking_portal:", ":blob_thinking_upsidedown:", ":blob_thonkang:", ":blob_thumbs_up:", ":blobblewobble:", ":blobenjoy:", ":blobglitch:", ":blobbass:", ":blobjam:", ":blobkeyboard:", ":bloblamp:", ":blobmaracas:", ":blobmicrophone:", ":blobthinksmart:", ":blobwobwork:", ":conga_party_thinking_blob:", ":Hyperblob:", ":party_blob:", ":partyparrot_blob:", ":partyparrot_blob_cat:"}

var Messages []Message = make([]Message, 0)
var Responses []OpenaiResponse = make([]OpenaiResponse, 0)

func ChatGPT(args ArgsV2) {
	msg := api.PostMessage(args.ChannelID, blobs[rand.Intn(len(blobs))]+":loading:")
	api.EditMessage(msg.Id, PostApiAndGetResponseTextAndRetryWhenError(msg, args.MessageText))
}

func ChatGPTReset(args ArgsV2) {
	msg := api.PostMessage(args.ChannelID, ":blobnom::loading:")
	Messages = make([]Message, 0)
	api.EditMessage(msg.Id, PostApiAndGetResponseTextAndRetryWhenError(msg, "ユーザーに向けて、<今までの会話履歴を削除し、リセットしました>という旨の文を返してください 謝る必要はありません ダブルクォーテーションも必要ありません"))
	Messages = make([]Message, 0)
	Responses = make([]OpenaiResponse, 0)
	return
}

func ChatGPTDebug(args ArgsV2) {
	returnString := "```\n"
	for _, m := range Messages {
		chatText := regexp.MustCompile("```").ReplaceAllString(m.Content, "")
		if len(chatText) >= 30 {
			returnString += m.Role + ": " + chatText[:30] + "...\n"
		} else {
			returnString += m.Role + ": " + chatText + "\n"
		}
	}
	returnString += "```\n```\n"
	prices := float32(0)
	for _, r := range Responses {
		prices += float32(r.Usage.TotalTokens) * (131.34 / 1000) * 0.002
	}
	r := Responses[len(Responses)-1]
	price := float32(r.Usage.TotalTokens) * (131.34 / 1000) * 0.002
	returnString += fmt.Sprintf("PromptTokens: %d\nCompletionTokens: %d\nTotalTokens: %d\n最後の一回で使った金額: %.2f円\n最後にリセットされてから使った合計金額:  %.2f円\n", r.Usage.PromptTokens, r.Usage.CompletionTokens, r.Usage.TotalTokens, price, prices)
	returnString += "```"
	api.PostMessage(args.ChannelID, returnString)
}

func PostApiAndGetResponseTextAndRetryWhenError(msg *traq.Message, input string) string {
	responseText, err := PostApiAndGetResponseText(input)
	if err != nil {
		if err == JsonError {
			api.EditMessage(msg.Id, "Error:"+fmt.Sprint(err)+"\nRETRYING :thonk_sweat: :loading::loading::loading:")
			responseText2, err2 := PostApiAndGetResponseText(input)
			if err2 != nil {
				return "Error:" + fmt.Sprint(err) + "\nError:" + fmt.Sprint(err2)
			}
			return responseText2
		} else {
			return "Error:" + fmt.Sprint(err)
		}
	}
	return responseText
}

func PostApiAndGetResponseText(input string) (string, error) {
	response, err := getOpenaiResponse(input)
	if err != nil {
		fmt.Println("Error:", err)
		return response.Result(), err
	}
	Responses = append(Responses, response)
	return response.Result(), nil
}

func (response OpenaiResponse) Result() string {
	response.AddText()
	return response.Text()
}

func (response OpenaiResponse) Text() string {
	if len(response.Choices) >= 1 {
		return response.Choices[0].Message.Content
	}
	return "Error: ResponseText nil"
}

func (response OpenaiResponse) AddText() {
	Messages = append(Messages, Message{
		Role:    "assistant",
		Content: response.Text(),
	})
}

func getOpenaiResponse(inputMessage string) (OpenaiResponse, error) {
	Messages = append(Messages, Message{
		Role:    "user",
		Content: inputMessage,
	})

	requestBody := OpenaiRequest{
		Model:    model,
		Messages: Messages,
	}

	return openaiRequest(requestBody)
}

func openaiRequest(requestBody OpenaiRequest) (OpenaiResponse, error) {
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return OpenaiResponse{}, err
	}

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		return OpenaiResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return OpenaiResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return OpenaiResponse{}, err
	}

	var response OpenaiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("%#v", body)
		return OpenaiResponse{}, JsonError
	}

	return response, nil
}
