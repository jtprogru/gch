package cas

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	CasApiUrl = "https://api.cas.chat/check"
)

type CasClient struct {
	verbose    bool
	httpClient *http.Client
	logger     *log.Logger
}

func New(timeout int, verbose bool) *CasClient {
	return &CasClient{
		verbose: verbose,
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
		logger: log.New(os.Stderr, "gch-cas-client: ", log.LstdFlags),
	}
}

func (r *CasClient) logf(format string, v ...interface{}) {
	if r.verbose {
		r.logger.Printf(format, v...)
	}
}

func (r *CasClient) Check(userId uint64) (bool, error) {
	var casResponse CasResponse

	url := fmt.Sprintf("%s?user_id=%d", CasApiUrl, userId)
	r.logf("Request URL: %s", url)
	req, err := http.NewRequest("GET", url, nil)
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
		r.logf("User is in the CAS list")
		return true, nil
	} else {
		r.logf("User is not in the CAS list")
		return false, nil
	}
}

type CasResponse struct {
	Ok bool `json:"ok"`
}

type CasResponseTrue struct {
	Ok     bool      `json:"ok"`
	Result CasResult `json:"result"`
}

type CasResult struct {
	Offenses  int       `json:"offenses"`
	Message   []string  `json:"message"`
	TimeAdded time.Time `json:"time_added"`
}

type CasResponseFalse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}
