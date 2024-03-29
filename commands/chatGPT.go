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
	"strings"
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
	Model   string   `json:"model"`
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
const model2 = "gpt-4"

var gpt4 = false
var lastGpt4 = false

const openaiURL = "https://api.openai.com/v1/chat/completions"

var blobs = [...]string{":blob_bongo:", ":blob_crazy_happy:", ":blob_grin:", ":blob_hype:", ":blob_love:", ":blob_lurk:", ":blob_pyon:", ":blob_pyon_inverse:", ":blob_slide:", ":blob_snowball_1:", ":blob_snowball_2:", ":blob_speedy_roll:", ":blob_speedy_roll_inverse:", ":blob_thinking:", ":blob_thinking_fast:", ":blob_thinking_portal:", ":blob_thinking_upsidedown:", ":blob_thonkang:", ":blob_thumbs_up:", ":blobblewobble:", ":blobenjoy:", ":blobglitch:", ":blobbass:", ":blobjam:", ":blobkeyboard:", ":bloblamp:", ":blobmaracas:", ":blobmicrophone:", ":blobthinksmart:", ":blobwobwork:", ":conga_party_thinking_blob:", ":Hyperblob:", ":party_blob:", ":partyparrot_blob:", ":partyparrot_blob_cat:"}

var Messages []Message = make([]Message, 0)
var Responses []OpenaiResponse = make([]OpenaiResponse, 0)
var SystemRoleMessage string = "あなたはサークルである東京工業大学デジタル創作同好会traPの部内SNS、traQのユーザーを楽しませる娯楽用途のBOTの中に作られた、openaiのモデルgpt-3.5-turboを用いた対話型AIです。身内しかいないSNSで、ユーザーに緩く接してください。そして、ユーザーの言う事に出来る限り従うようにしてください。"

func ChatGPT(args ArgsV2) {
	msg := api.PostMessage(args.ChannelID, blobs[rand.Intn(len(blobs))]+":loading:")
	response, err := PostApiAndGetResponseAndRetryWhenError(msg, args.MessageText)
	if err != nil {
		api.EditMessage(msg.Id, "Error: "+fmt.Sprint(err))
	}
	api.EditMessage(msg.Id, response.Text())
	lastGpt4 = false
}

func ChatGPT4(args ArgsV2) {
	msg := api.PostMessage(args.ChannelID, blobs[rand.Intn(len(blobs))]+":loading:(gpt-4)")
	gpt4 = true
	response, err := PostApiAndGetResponseAndRetryWhenError(msg, args.MessageText)
	if err != nil {
		api.EditMessage(msg.Id, "Error: "+fmt.Sprint(err))
	}
	api.EditMessage(msg.Id, response.Text())
	lastGpt4 = true
}

func ChatGPTReset(args ArgsV2) {
	msg := api.PostMessage(args.ChannelID, ":blobnom::loading:")
	Messages = make([]Message, 0)
	response, err := PostApiAndGetResponseAndRetryWhenError(msg, "ユーザーに向けて、<今までの会話履歴を削除し、リセットしました>という旨の文を返してください 謝る必要はありません ダブルクォーテーションも必要ありません")
	if err != nil {
		api.EditMessage(msg.Id, "Error: "+fmt.Sprint(err))
	}
	api.EditMessage(msg.Id, response.Text())
	Messages = make([]Message, 0)
	Responses = make([]OpenaiResponse, 0)
	return
}

func Sum(arr []float32) float32 {
	var res float32 = 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
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
	var prices []float32
	for _, r := range Responses {
		if strings.Contains(r.Model, "gpt-4") {
			prices = append(prices, float32(r.Usage.PromptTokens)*(132.29/1000)*0.03+float32(r.Usage.CompletionTokens)*(132.29/1000)*0.06)
			continue
		} else if strings.Contains(r.Model, "gpt-3.5") {
			prices = append(prices, float32(r.Usage.TotalTokens)*(132.29/1000)*0.002)
			continue
		}
	}
	if len(Responses) == 0 {
		api.PostMessage(args.ChannelID, returnString)
		return
	}
	r := Responses[len(Responses)-1]
	var price float32
	if strings.Contains(r.Model, "gpt-4") {
		price = float32(r.Usage.PromptTokens)*(132.29/1000)*0.03 + float32(r.Usage.CompletionTokens)*(132.29/1000)*0.06
	} else if strings.Contains(r.Model, "gpt-3.5") {
		price = float32(r.Usage.TotalTokens) * (132.29 / 1000) * 0.002
	}
	returnString += fmt.Sprintf("PromptTokens: %d\nCompletionTokens: %d\nTotalTokens: %d\n最後の一回で使った金額: %.2f円\n最後にリセットされてから使った合計金額:  %.2f円\n", r.Usage.PromptTokens, r.Usage.CompletionTokens, r.Usage.TotalTokens, price, Sum(prices))
	returnString += "```"
	api.PostMessage(args.ChannelID, returnString)
}

func ChatGPTChangeFirstSystemMessage(args ArgsV2) {
	SystemRoleMessage = args.MessageText
	api.PostMessage(args.ChannelID, "FirstSystemMessageを変更しました。/gptsys showで確認できます。\nFirstSystemMessageとは、常に履歴の一番最初に入り、最初にgptに情報や状況を説明するのに使用する文字列です")
}

func ChatGPTShowFirstSystemMessage(args ArgsV2) {
	api.PostMessage(args.ChannelID, SystemRoleMessage)
}

func PostApiAndGetResponseAndRetryWhenError(msg *traq.Message, input string) (OpenaiResponse, error) {
	response, err := PostApiAndGetResponse(input)
	for i := 0; overTokenCheck(response) && i <= 4; i++ {
		api.EditMessage(msg.Id, "Clearing old history and retrying.["+fmt.Sprintf("%d", i+1)+"] :loading:")
		if len(Messages) >= 5 {
			Messages = Messages[4:]
			Messages = Messages[:len(Messages)-1]
		} else if len(Messages) >= 2 {
			Messages = Messages[1:]
			Messages = Messages[:len(Messages)-1]
		} else if len(Messages) >= 1 {
			Messages = Messages[1:]
		}
		response, err = PostApiAndGetResponse(input)
		if err != nil {
			api.EditMessage(msg.Id, "Error:"+fmt.Sprint(err)+"\nRETRYING :thonk_sweat: :loading::loading::loading:")
			return Retry(msg, input, err)
		}
	}
	if err != nil {
		api.EditMessage(msg.Id, "Error:"+fmt.Sprint(err)+"\nRETRYING :thonk_sweat: :loading::loading::loading:")
		return Retry(msg, input, err)
	}
	return response, nil
}

func Retry(msg *traq.Message, input string, err error) (OpenaiResponse, error) {
	response, err2 := PostApiAndGetResponse(input)
	if err2 != nil {
		return response, errors.New(fmt.Sprint(err) + "\nError:" + fmt.Sprint(err2))
	}
	return response, nil
}

// overToken -> true
func overTokenCheck(response OpenaiResponse) bool {
	if len(response.Choices) == 0 {
		return true
	}
	if response.Choices[0].FinishReason == "length" {
		return true
	}
	return false
}

func PostApiAndGetResponse(input string) (OpenaiResponse, error) {
	updateFirstSystemRoleMessage(SystemRoleMessage)
	response, err := getOpenaiResponse(input)
	if err != nil {
		fmt.Println("Error:", err)
		return response, err
	}
	Responses = append(Responses, response)
	return response, nil
}

func (response OpenaiResponse) Text() string {
	if len(response.Choices) >= 1 {
		response.AddText()
		return response.Choices[0].Message.Content
	}
	return "Error: ResponseText nil"
}

func (response OpenaiResponse) AddText() {
	Messages = append(Messages, Message{
		Role:    "assistant",
		Content: response.Choices[0].Message.Content,
	})
}

func getOpenaiResponse(inputMessage string) (OpenaiResponse, error) {
	Messages = append(Messages, Message{
		Role:    "user",
		Content: inputMessage,
	})

	requestBody := OpenaiRequest{}
	if gpt4 {
		requestBody = OpenaiRequest{
			Model:    model2,
			Messages: Messages,
		}
		gpt4 = false
	} else {
		requestBody = OpenaiRequest{
			Model:    model,
			Messages: Messages,
		}
	}

	return openaiRequest(requestBody)
}

func addSystemRoleMessage(systemMessage string) {
	Messages = append(Messages, Message{
		Role:    "system",
		Content: systemMessage,
	})
}

func addFirstSystemRoleMessageIfNotExist(systemMessage string) {
	for _, m := range Messages {
		if m.Role == "system" {
			return
		}
	}
	systemM := Message{
		Role:    "system",
		Content: systemMessage,
	}
	Messages = append([]Message{systemM}, Messages...)
}

func updateFirstSystemRoleMessage(systemMessage string) {
	addFirstSystemRoleMessageIfNotExist(systemMessage)
	systemM := Message{
		Role:    "system",
		Content: systemMessage,
	}
	Messages[0] = systemM
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
