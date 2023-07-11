/*
Copyright © 2023 WABEL GROUP <m.lesage@wabelgroup.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	baseURL        = "https://api.wabeltools.com/v1"
	imgURL         = baseURL + "/img"
	nlpURL         = baseURL + "/nlp"
	configFileName = ".wabeltools"
)

var rootCmd = &cobra.Command{
	Use:   "wabeltools",
	Short: "CLI for the Wabel Tools service",
	Long: `wabeltools is a CLI application for interacting with the Wabel Tools service.
You can use it to use various tools provided by the service like image processing, NLP, etc.

For example:

As first step, you need to initialize the tool with your API key:
wabeltools init [apikey]

Then you can use the various commands:

In image processing:
wabeltools image local --resize 100x100 "image.jpg"
wabeltools image local --compress "path/to/image.jpg"
wabeltools image remote --many --compress --urls="https://www.wabel.com/images/logo.png,https://www.wabel.com/images/logo.png"

In NLP:
wabeltools nlp sentiment "I love Wabel Tools!"
wabeltools nlp stemming --lang=fr "J'adore Wabel Tools!"
wabeltools nlp stopwords --lang=en "Je vis à Paris, en France."`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Name() == "init" {
			return
		}
		if cmd.Name() == "help" {
			return
		}
		home, _ := os.UserHomeDir()
		configPath := filepath.Join(home, configFileName)
		apikey, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Println("API key is not set. Run 'wabeltools init' to set it.")
			os.Exit(1)
		}
		viper.Set("apikey", string(apikey)) // Set apikey in viper for other commands to use
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigName(configFileName)
	viper.AddConfigPath("$HOME")
	viper.ReadInConfig()

	// Add subcommands
	rootCmd.AddCommand(initCmd, tokensCmd, costsCmd, servicesCmd, imgCmd, nlpCmd)
}
