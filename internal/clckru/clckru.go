package clckru

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	// Endpoint for the URL shortening service.
	endpoint = "https://clck.ru/--"
)

// Shorten shortens a long URL by adding UTM tags if specified.
func Shorten(longURL, utmTags string) (string, error) {
	// Forming the full URL with UTM tags if they are provided.
	fullURL := longURL
	if utmTags != "" {
		fullURL = fmt.Sprintf("%s?%s", longURL, utmTags)
	}

	// Preparing the data for a POST request.
	data := url.Values{}
	data.Set("url", fullURL)

	// Using the Background context to create the request.
	// This solution allows for future extensions (e.g., timeout management).
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Setting headers for the request.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Sending the HTTP request using a client.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Checking the response status.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	// Reading the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Returning the result as a string.
	return string(body), nil
}
