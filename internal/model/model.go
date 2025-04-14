package model

import (
	"LawPrompt/config"
	context2 "LawPrompt/internal/context"
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

	context, err := context2.GetContext(prompt, cfg.Env.SearchUrl)
	if err != nil {
		return "Failed to get Context", err
	}

	client := resty.New()
	var result ModelResponse

	promptContext := fmt.Sprintf(`Ты — юридический помощник, обученный на кодексах и других нормативных актах Республики Казахстан. Тебе нужно дать ответ от приведенного ниже правильного ответа в уважительной форме без '\n'.

		Ответ:
		"""
		%s
		"""

		Вопрос:
		%s`, context, prompt)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]any{
			"model":  cfg.Env.Model,
			"prompt": promptContext,
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
