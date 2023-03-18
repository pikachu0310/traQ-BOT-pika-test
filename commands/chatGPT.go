package commands

import (
	"bytes"
	"encoding/json"
	"example-bot/api"
	"example-bot/util"
	"fmt"
	"io"
	"io/ioutil"
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

const model = "gpt-3.5-turbo"
const openaiURL = "https://api.openai.com/v1/chat/completions"

var InputMessages []Message = make([]Message, 0)

func ChatGPT(args ArgsV2) {
	chatgptStampMatch := regexp.MustCompile(`(new(chat|gpt|chatgpt)|(chat|gpt|chatgpt)new|new (chat|gpt|chatgpt)|(chat|gpt|chatgpt) new)`)
	if chatgptStampMatch.MatchString(args.MessageText) {
		InputMessages = make([]Message, 0)
		api.PostMessage(args.ChannelID, PostApiAndGetResponseText("今までの会話履歴を削除し、リセットしました"))
		return
	}
	api.PostMessage(args.ChannelID, PostApiAndGetResponseText(args.MessageText))
}

func PostApiAndGetResponseText(input string) string {
	response, err := getOpenaiResponse(input)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error:" + fmt.Sprint(err)
	}
	return response.Result()
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
	InputMessages = append(InputMessages, Message{
		Role:    "assistant",
		Content: response.Text(),
	})
}

func getOpenaiResponse(inputMessage string) (OpenaiResponse, error) {
	InputMessages = append(InputMessages, Message{
		Role:    "user",
		Content: inputMessage,
	})

	requestBody := OpenaiRequest{
		Model:    model,
		Messages: InputMessages,
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
		return OpenaiResponse{}, err
	}

	return response, nil
}
