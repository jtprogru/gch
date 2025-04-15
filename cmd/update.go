package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command.
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gch",
	Long:  `Check available update and upgrade gch.`,
	Run: func(cmd *cobra.Command, args []string) {
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

var update bool

func init() {
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
		TagName string `json:"tag_name"`
		Assets  []struct {
			BrowserDownloadURL string `json:"browser_download_url"`
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
		return "", "", fmt.Errorf("no assets found for the latest release")
	}

	return release.TagName, release.Assets[0].BrowserDownloadURL, nil
}

// downloadAndInstall downloads the latest release and installs it.
func downloadAndInstall(downloadURL string) error {
	// Download the file
	resp, err := http.Get(downloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "gch-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return err
	}

	// Make the file executable
	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		return err
	}

	// Replace the current binary
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
