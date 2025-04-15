package yandexgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	endpoint       = "https://300.ya.ru/api/sharing-url"
	defaultTimeout = time.Second * 15
)

// Response represents the expected JSON response from the API.
type Response struct {
	Status     string `json:"status"`
	SharingURL string `json:"sharing_url"` //nolint:tagliatelle // This is predefined by the API
}

var ErrTokenNotSet = errors.New("environment variable TOKEN_300_YA_RU is not set")

// Brief sends a POST request with the given URL to the API and retrieves a sharing URL.
func Brief(url string) (string, error) {
	// Ensure the token is set in environment variables.
	token := os.Getenv("TOKEN_300_YA_RU")
	if token == "" {
		return "", ErrTokenNotSet
	}

	httpClient := &http.Client{}

	// Use a context with timeout for the request.
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Prepare the request payload.
	payload := map[string]string{
		"article_url": url,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create the HTTP request.
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth "+token)

	// Send the HTTP request.
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Ensure the response status is HTTP 200 (OK).
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	// Parse the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check for success status in the response.
	if apiResponse.Status != "success" {
		return "", fmt.Errorf("API error: response status is '%s'", apiResponse.Status)
	}

	return apiResponse.SharingURL, nil
}
