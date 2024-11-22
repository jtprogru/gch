package yandexgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	endpoint = "https://300.ya.ru/api/sharing-url"
)

func Breif(url string) (string, error) {
	// Формирование тела запроса
	payload := map[string]string{
		"article_url": url,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Создание HTTP-запроса
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")

	token := os.Getenv("TOKEN_300_YA_RU")

	req.Header.Set("Authorization", fmt.Sprintf("OAuth %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	temporaryStruct := make(map[string]string)

	json.Unmarshal(body, &temporaryStruct)
	if temporaryStruct["status"] != "success" {
		return "", fmt.Errorf("error: %s", temporaryStruct["status"])
	}

	// Возвращаем результат в виде строки
	return temporaryStruct["sharing_url"], nil
}

type YandexGPTResponse struct {
	Status     string `json:"status"`
	SharingUrl string `json:"sharing_url"`
}
