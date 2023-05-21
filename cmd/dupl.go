/*
Copyright © 2023 Michael <jtprogru@gmail> Savin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
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
