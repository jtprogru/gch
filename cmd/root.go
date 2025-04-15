package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string   //nolint:gochecknoglobals,nolintlint // This is normal.
var showVersion bool //nolint:gochecknoglobals,nolintlint // This is normal.

var version = "dev"

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
	Use:   "gch",
	Short: "Go CLI Helper",
	Long: `Go CLI Helper this is a simple CLI utility that helps
make my life easier and will be gradually supplemented with various functionality.

Now gch is not able to do so much, but I use it every day.

Complete documentation is available at https://github.com/jtprogru/gch/wiki`,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if showVersion {
		fmt.Println("gch version:", version)
		os.Exit(0)
	}
	// Execute the root command.
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1) //nolint:revive,nolintlint // This is normal.
	}
}

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gch.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show version")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gch" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gch")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match.

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		_, _ = fmt.Printf("Error loading config: %s", err) //nolint:errcheck,nolintlint // Ignore errors.
	}
}
