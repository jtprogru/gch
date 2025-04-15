package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	// uuidCmd represents the uuid command.
	uuidCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
		Use:   "uuid",
		Short: "Generate UUID string",
		Long: `A UUID is a 16 byte (128 bit) array.
UUIDs may be used as keys to maps or compared directly.`,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(GetUUID())
		},
	}
	v4UUID    bool                          //nolint:gochecknoglobals,nolintlint // This is normal.
	v5UUID    bool                          //nolint:gochecknoglobals,nolintlint // This is normal.
	nilUUID   bool                          //nolint:gochecknoglobals,nolintlint // This is normal.
	namespace = uuid.NameSpaceURL           //nolint:gochecknoglobals,nolintlint // This is normal.
	url       = []byte("https://jtprog.ru") //nolint:gochecknoglobals,nolintlint // This is normal.
)

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
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
