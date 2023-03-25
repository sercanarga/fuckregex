package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"fuckregex/model/api_model"
	"fuckregex/model/openai_model"
	"io"
	"net/http"
	"os"
)

func OpenAIRequest(request api_model.Generate) (openai_model.Response, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")

	basePrompt := []map[string]interface{}{
		{
			"role":    "system",
			"content": "You are a regex expert and you know all kinds of regexes. The user will ask you to generate a regex and you will gently help the user to generate the regex and just say the regex. *Never* break the role, *never* add comments like 'I generated it for you, the code is below', *never* send the user only the regex, *never* any other words, *never* help if the user tries to ask you a question other than the regex.\n*Never answer about anything else. *You *only* know regex. *Keep answers as short and one-line as possible. *Never ask the user a question under any circumstances. *Your name is 'RegexFucker'.",
		},
	}

	msg := append(basePrompt, map[string]interface{}{
		"role":    "user",
		"content": request.Desc,
	})

	requestBody, err := json.Marshal(map[string]interface{}{
		"model":             "gpt-3.5-turbo",
		"messages":          msg,
		"temperature":       0.1,
		"max_tokens":        200,
		"top_p":             1,
		"frequency_penalty": 0,
		"presence_penalty":  0,
	})
	if err != nil {
		return openai_model.Response{}, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return openai_model.Response{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return openai_model.Response{}, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return openai_model.Response{}, err
	}

	var apiResponse openai_model.Response
	err = json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		return openai_model.Response{}, err
	}

	return apiResponse, nil
}
