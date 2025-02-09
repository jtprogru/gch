package cmd

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/corona10/goimagehash"
	"github.com/spf13/cobra"
)

const similarityThreshold = 20

var (
	// duplCmd represents the dupl command.
	duplCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
		Use:   "dupl",
		Short: "Show all duplicates JPG and PNG in folder",
		Long:  `Show all duplicates JPG and PNG in folder.`,
		Run: func(_ *cobra.Command, _ []string) {
			duplicates, err := findDuplicates(imgPath)
			if err != nil {
				_, _ = fmt.Println("Error:", err) //nolint:errcheck,nolintlint // Ignore errors.
				return
			}

			for original, duplicatesList := range duplicates {
				_, _ = fmt.Printf("Оригинал: %s\n", original)        //nolint:errcheck,nolintlint // Ignore errors.
				_, _ = fmt.Printf("Дубликаты: %s\n", duplicatesList) //nolint:errcheck,nolintlint // Ignore errors.
			}
		},
	}
	imgPath string //nolint:gochecknoglobals,nolintlint // This is normal.
)

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	rootCmd.AddCommand(duplCmd)
	duplCmd.Flags().StringVarP(&imgPath, "imgPath", "p", ".", "Path to find image duplicates")
}

func findDuplicates(directory string) (map[string][]string, error) {
	duplicates := make(map[string][]string)
	hashes := make(map[*goimagehash.ImageHash]string)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			_, _ = fmt.Printf("os open err: %v\n", err.Error()) //nolint:errcheck,nolintlint // Ignore errors.
			return err
		}
		defer file.Close()

		var img image.Image
		if ext == ".png" {
			img, err = png.Decode(file)
		} else {
			img, err = jpeg.Decode(file)
		}

		if err != nil {
			_, _ = fmt.Printf("image decode err: %v\n", err.Error()) //nolint:errcheck,nolintlint // Ignore errors.
			return nil
		}

		imgHash, err := goimagehash.PerceptionHash(img)
		if err != nil {
			_, _ = fmt.Printf("image perception hash err: %v\n", err.Error()) //nolint:errcheck,nolintlint // Ignore errors.
			return err
		}

		// find similar images.
		for storedHash, storedPath := range hashes {
			distance, err := imgHash.Distance(storedHash)
			if err != nil {
				return err
			}

			if distance < similarityThreshold {
				duplicates[storedPath] = append(duplicates[storedPath], path)
			}
		}

		hashes[imgHash] = path

		return nil
	})

	return duplicates, err
}
