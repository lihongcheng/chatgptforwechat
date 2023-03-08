package gtp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/869413421/wechatbot/config"
	"io/ioutil"
	"log"
	"net/http"
)

const BASEURL = "https://api.openai.com/v1/"

// ChatGPTResponseBody 响应体

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
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

type ChatGPTResponseBody struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// ChatGPTRequestBody 请求体

type SendMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type ChatGPTRequestBody struct {
	Model            string        `json:"model"`
	Messages         []SendMessage `json:"messages"`
	Temperature      float64       `json:"temperature"`
	MaxTokens        int           `json:"max_tokens"`
	TopP             float64       `json:"top_p"`
	FrequencyPenalty float64       `json:"frequency_penalty"`
	PresencePenalty  float64       `json:"presence_penalty"`
}

// Completions gtp文本模型回复
// curl https://api.openai.com/v1/completions
// -H "Content-Type: application/json"
// -H "Authorization: Bearer your chatGPT key"
// -d '{"model": "text-davinci-003", "prompt": "give me good song", "temperature": 0, "max_tokens": 7}'
func Completions(msg string) (string, error) {
	sendMsg := SendMessage{
		Role:    "user",
		Content: msg,
	}
	messages := []SendMessage{sendMsg}
	requestBody := ChatGPTRequestBody{
		Model:            "gpt-3.5-turbo",
		Messages:         messages,
		Temperature:      0.7,
		MaxTokens:        4000,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	log.Printf("request gtp json string : %v", string(requestData))
	req, err := http.NewRequest("POST", BASEURL+"chat/completions", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	apiKey := config.LoadConfig().ApiKey
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("gtp api status code not equals 200,code is %d", response.StatusCode))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}

	var reply string
	if len(gptResponseBody.Choices) > 0 {
		reply = gptResponseBody.Choices[0].Message.Content
	}
	log.Printf("gpt response text: %s \n", reply)
	return reply, nil
}
