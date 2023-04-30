/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/corona10/goimagehash"
	"github.com/spf13/cobra"
)

var (
	// duplCmd represents the dupl command
	duplCmd = &cobra.Command{
		Use:   "dupl",
		Short: "Show all duplicates JPG and PNG in folder",
		Long:  `Show all duplicates JPG and PNG in folder.`,
		Run: func(cmd *cobra.Command, args []string) {
			duplicates, err := findDuplicates(imgPath)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			for original, duplicatesList := range duplicates {
				fmt.Printf("Оригинал: %s\n", original)
				fmt.Printf("Дубликаты: %s\n", duplicatesList)
			}
		},
	}
	imgPath string
)

func init() {
	rootCmd.AddCommand(duplCmd)

	duplCmd.Flags().StringVarP(&imgPath, "imgPath", "p", ".", "Path to find image duplicates")
}

// TODO: Потестировать алгоримт на живых примерах
func findDuplicates(directory string) (map[string][]string, error) {
	duplicates := make(map[string][]string)
	hashes := make(map[string]string)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("os open err: %v\n", err.Error())
			return err
		}
		defer file.Close()
		fmt.Printf("open file: %v\n", file.Name())

		// FIXME: Исправить часто возникающую ошибку "image decode err: image: unknown format"
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Printf("image decode err: %v\n", err.Error())
			return nil
		}
		fmt.Printf("image decoded: %v\n", file.Name())

		imgHash, err := goimagehash.PerceptionHash(img)
		if err != nil {
			fmt.Printf("image perception hash err: %v\n", err.Error())
			return err
		}

		hashStr := imgHash.ToString()
		if original, exists := hashes[hashStr]; exists {
			duplicates[original] = append(duplicates[original], path)
		} else {
			hashes[hashStr] = path
		}
		fmt.Printf("image processed: %v\n", file.Name())

		return nil
	})

	return duplicates, err
}
