package cmd

import (
	"os"
	"text/template"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// cbrfCmd represents the cbrf command
	licCmd = &cobra.Command{
		Use:   "lic",
		Short: "Generate new WTFPL license for you project",
		Long:  `Just for fun! Generate new WTFPL license for LICENSE file with your email, full name and current year.`,
		Run: func(cmd *cobra.Command, args []string) {

			Email := viper.GetString("email")
			FullName := viper.GetString("full_name")
			Year, _, _ := time.Now().Date()
			tmpl, err := template.New("wtfpl").Parse(licenseTpl)
			if err != nil {
				panic(err)
			}

			err = tmpl.Execute(os.Stdout, data{Email, FullName, Year})
			if err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(licCmd)

	// licCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all rates")
}

type data struct {
	Email    string
	FullName string
	Year     int
}

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
