package gch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/mattn/go-isatty"
	"gopkg.in/yaml.v3"
)

var gitDescribeSuffixRE = regexp.MustCompile(`\d+-\d+-g[a-f0-9]{8}$`)

// ReleaseInfo stores information about a release
type ReleaseInfo struct {
	Version     string    `json:"tag_name"`
	URL         string    `json:"html_url"`
	PublishedAt time.Time `json:"published_at"`
}

type StateEntry struct {
	CheckedForUpdateAt time.Time   `yaml:"checked_for_update_at"`
	LatestRelease      ReleaseInfo `yaml:"latest_release"`
}

// ShouldCheckForUpdate decides whether we check for updates for the GitHub CLI based on user preferences and current execution context.
func ShouldCheckForUpdate() bool {
	if os.Getenv("GCH_NO_UPDATE_NOTIFIER") != "" {
		return false
	}
	if os.Getenv("CODESPACES") != "" {
		return false
	}
	return !IsCI() && IsTerminal(os.Stdout) && IsTerminal(os.Stderr)
}

// CheckForUpdate checks whether an update exists for the GitHub CLI based on recency of last check within past 24 hours.
func CheckForUpdate(ctx context.Context, client *http.Client, stateFilePath, repo, currentVersion string) (*ReleaseInfo, error) {
	stateEntry, _ := getStateEntry(stateFilePath)
	if stateEntry != nil && time.Since(stateEntry.CheckedForUpdateAt).Hours() < 24 {
		return nil, nil
	}

	releaseInfo, err := getLatestReleaseInfo(ctx, client, repo)
	if err != nil {
		return nil, err
	}

	err = setStateEntry(stateFilePath, time.Now(), *releaseInfo)
	if err != nil {
		return nil, err
	}

	if versionGreaterThan(releaseInfo.Version, currentVersion) {
		return releaseInfo, nil
	}

	return nil, nil
}

func getLatestReleaseInfo(ctx context.Context, client *http.Client, repo string) (*ReleaseInfo, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo), nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		res.Body.Close()
	}()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected HTTP %d", res.StatusCode)
	}
	dec := json.NewDecoder(res.Body)
	var latestRelease ReleaseInfo
	if err := dec.Decode(&latestRelease); err != nil {
		return nil, err
	}
	return &latestRelease, nil
}

func getStateEntry(stateFilePath string) (*StateEntry, error) {
	content, err := os.ReadFile(stateFilePath)
	if err != nil {
		return nil, err
	}

	var stateEntry StateEntry
	err = yaml.Unmarshal(content, &stateEntry)
	if err != nil {
		return nil, err
	}

	return &stateEntry, nil
}

func setStateEntry(stateFilePath string, t time.Time, r ReleaseInfo) error {
	data := StateEntry{CheckedForUpdateAt: t, LatestRelease: r}
	content, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(stateFilePath), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(stateFilePath, content, 0600)
	return err
}

func versionGreaterThan(v, w string) bool {
	w = gitDescribeSuffixRE.ReplaceAllStringFunc(w, func(m string) string {
		idx := strings.IndexRune(m, '-')
		n, _ := strconv.Atoi(m[0:idx])
		return fmt.Sprintf("%d-pre.0", n+1)
	})

	vv, ve := version.NewVersion(v)
	vw, we := version.NewVersion(w)

	return ve == nil && we == nil && vv.GreaterThan(vw)
}

// IsTerminal determines if a file descriptor is an interactive terminal / TTY.
func IsTerminal(f *os.File) bool {
	return isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
}

// IsCI determines if the current execution context is within a known CI/CD system.
// This is based on https://github.com/watson/ci-info/blob/HEAD/index.js.
func IsCI() bool {
	return os.Getenv("CI") != "" || // GitHub Actions, Travis CI, CircleCI, Cirrus CI, GitLab CI, AppVeyor, CodeShip, dsari
		os.Getenv("BUILD_NUMBER") != "" || // Jenkins, TeamCity
		os.Getenv("RUN_ID") != "" // TaskCluster, dsari
}

type GithubReleaseResponse struct {
	URL             string         `json:"url"`
	AssetsURL       string         `json:"assets_url"`
	UploadURL       string         `json:"upload_url"`
	HTMLURL         string         `json:"html_url"`
	ID              int            `json:"id"`
	Author          GithubAuthor   `json:"author"`
	NodeID          string         `json:"node_id"`
	TagName         string         `json:"tag_name"`
	TargetCommitish string         `json:"target_commitish"`
	Name            string         `json:"name"`
	Draft           bool           `json:"draft"`
	Prerelease      bool           `json:"prerelease"`
	CreatedAt       time.Time      `json:"created_at"`
	PublishedAt     time.Time      `json:"published_at"`
	Assets          []GithubAssets `json:"assets"`
	TarballURL      string         `json:"tarball_url"`
	ZipballURL      string         `json:"zipball_url"`
	Body            string         `json:"body"`
}

type GithubAuthor struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	UserViewType      string `json:"user_view_type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type GithubUploader struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	UserViewType      string `json:"user_view_type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type GithubAssets struct {
	URL                string         `json:"url"`
	ID                 int            `json:"id"`
	NodeID             string         `json:"node_id"`
	Name               string         `json:"name"`
	Label              string         `json:"label"`
	Uploader           GithubUploader `json:"uploader"`
	ContentType        string         `json:"content_type"`
	State              string         `json:"state"`
	Size               int            `json:"size"`
	DownloadCount      int            `json:"download_count"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	BrowserDownloadURL string         `json:"browser_download_url"`
}
