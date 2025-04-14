package context

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetContext(prompt string, url string) (string, error) {
	payload := map[string]string{"question": prompt}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Answer string `json:"answer"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("failed to decode JSON response: %w", err)
	}

	return response.Answer, nil
}
