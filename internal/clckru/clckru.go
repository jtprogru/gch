package clckru

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	// Endpoint for the URL shortening service
	endpoint = "https://clck.ru/--"
)

func Shorten(longUrl string, utmTags string) (string, error) {
	fullURL := longUrl
	if utmTags != "" {
		fullURL = fmt.Sprintf("%s?%s", longUrl, utmTags)
	}
	data := url.Values{}
	data.Set("url", fullURL)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

	// Возвращаем результат в виде строки
	return string(body), nil
}
