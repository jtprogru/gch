package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command.
var updateCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
	Use:   "update",
	Short: "Update gch",
	Long:  `Check available update and upgrade gch.`,
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := context.Background()
		latestVersion, downloadURL, err := getLatestRelease(ctx, "jtprogru/gch")
		if err != nil {
			fmt.Printf("Error checking for updates: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Current version: %s\n", cmd.Version)
		fmt.Printf("Latest version: %s\n", latestVersion)

		if latestVersion == cmd.Version {
			fmt.Println("You are already using the latest version.")
			return
		}

		if update {
			fmt.Println("Updating to the latest version...")
			err := downloadAndInstall(downloadURL)
			if err != nil {
				fmt.Printf("Error updating gch: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Update successful!")
		} else {
			fmt.Println("Run the command with --upgrade/-u to update to the latest version.")
		}
	},
}

var update bool //nolint:gochecknoglobals,nolintlint // This is normal.

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolVarP(&update, "upgrade", "u", false, "Upgrade gch")
}

// getLatestRelease fetches the latest release information from GitHub.
func getLatestRelease(ctx context.Context, repo string) (string, string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to fetch release info: %s", resp.Status)
	}

	var release struct {
		TagName string `json:"tag_name"` //nolint:tagliatelle // This is predefined by the API.
		Assets  []struct {
			BrowserDownloadURL string `json:"browser_download_url"` //nolint:tagliatelle // This is predefined by the API.
		} `json:"assets"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	err = json.Unmarshal(body, &release)
	if err != nil {
		return "", "", err
	}

	if len(release.Assets) == 0 {
		return "", "", errors.New("no assets found for the latest release")
	}

	return release.TagName, release.Assets[0].BrowserDownloadURL, nil
}

// downloadAndInstall downloads the latest release and installs it.
func downloadAndInstall(downloadURL string) error {
	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Создаем HTTP-запрос с контекстом
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, downloadURL, nil)
	if err != nil {
		return err
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	// Создаем временный файл
	tmpFile, err := os.CreateTemp("", "gch-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return err
	}

	// Делаем файл исполняемым
	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		return err
	}

	// Заменяем текущий бинарный файл
	execPath, err := os.Executable()
	if err != nil {
		return err
	}

	err = os.Rename(tmpFile.Name(), execPath)
	if err != nil {
		return err
	}

	return nil
}
