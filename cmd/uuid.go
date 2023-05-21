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

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var (
	uuidCmd = &cobra.Command{
		Use:   "uuid",
		Short: "Generate UUID string",
		Long: `A UUID is a 16 byte (128 bit) array.
UUIDs may be used as keys to maps or compared directly.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(GetUUID())
		},
	}
	v4UUID    bool
	v5UUID    bool
	nilUUID   bool
	namespace = uuid.NameSpaceURL
	url       = []byte("https://jtprog.ru")
)

func init() {
	rootCmd.AddCommand(uuidCmd)
	uuidCmd.Flags().BoolVarP(&v4UUID, "v4", "4", true, "Generate UUID4")
	uuidCmd.Flags().BoolVarP(&v5UUID, "v5", "5", false, "Generate UUID5")
	uuidCmd.Flags().BoolVarP(&nilUUID, "v0", "0", false, "Generate UUIDnil")

}

func GetUUID() string {
	switch {
	case v4UUID:
		return uuid.New().String()
	case v5UUID:
		return uuid.NewMD5(namespace, url).String()
	case nilUUID:
		return "00000000-0000-0000-0000-000000000000"
	default:
		return "00000000-0000-0000-0000-000000000000"
	}
}
