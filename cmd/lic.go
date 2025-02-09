package cmd

import (
	"os"
	"text/template"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cbrfCmd represents the cbrf command.
var licCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
	Use:   "lic",
	Short: "Generate new WTFPL license for you project",
	Long:  `Just for fun! Generate new WTFPL license for LICENSE file with your email, full name and current year.`,
	Run: func(_ *cobra.Command, _ []string) {
		email := viper.GetString("email")
		fullName := viper.GetString("full_name")
		year, _, _ := time.Now().Date()
		tmpl, err := template.New("wtfpl").Parse(licenseTpl)
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(os.Stdout, data{email, fullName, year})
		if err != nil {
			panic(err)
		}
	},
}

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	rootCmd.AddCommand(licCmd)
}

type data struct {
	Email    string
	FullName string
	Year     int
}

//nolint:gochecknoglobals,nolintlint // This is normal.
var licenseTpl = `            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) {{ .Year }} {{ .FullName }} <{{ .Email }}>

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.

`
