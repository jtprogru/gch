package cas

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	baseURL    string
	verbose    bool
	httpClient *http.Client
	timeout    time.Duration
	logger     *log.Logger
}

func New(timeout int, verbose bool) *Client {
	return &Client{
		baseURL: "https://api.cas.chat/check",
		verbose: verbose,
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
		timeout: time.Duration(timeout) * time.Second,
		logger:  log.New(os.Stderr, "gch-cas-client: ", log.LstdFlags),
	}
}

func (r *Client) logf(format string, v ...any) {
	if r.verbose {
		r.logger.Printf(format, v...)
	}
}

func (r *Client) Check(userID uint64) (bool, error) {
	var casResponse Response
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	url := fmt.Sprintf("%s?user_id=%d", r.baseURL, userID)
	r.logf("Request URL: %s", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		r.logf("Error while creating request: %v", err)
		return false, err
	}

	req.Header.Set("User-Agent", "gch-cas-client")
	resp, err := r.httpClient.Do(req)
	if err != nil {
		r.logf("Error while sending request: %v", err)
		return false, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&casResponse)
	if err != nil {
		r.logf("Error while decoding default response: %v", err)
		return false, err
	}

	if casResponse.Ok {
		r.logf("BAN! User is in the CAS list")
		return true, nil
	}
	r.logf("CLEAN! User is not in the CAS list")
	return false, nil
}

type Response struct {
	Ok bool `json:"ok"`
}

type ResponseTrue struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	Offenses  int       `json:"offenses"`
	Message   []string  `json:"message"`
	TimeAdded time.Time `json:"time_added"`
}

type ResponseFalse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}
