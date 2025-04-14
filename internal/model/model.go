package model

import (
	"LawPrompt/config"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type PromptRequest struct {
	Question string `json:"question"`
}

type ModelResponse struct {
	Response string `json:"response"`
}

func AskModel(prompt string) (string, error) {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		return "", err
	}

	client := resty.New()
	var result ModelResponse

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]any{
			"model":  cfg.Env.Model,
			"prompt": prompt,
			"stream": false,
		}).
		SetResult(&result).
		Post(cfg.Env.OllamaUrl)

	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}

	return result.Response, nil
}
