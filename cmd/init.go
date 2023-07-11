package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [apikey]",
	Short: "Initialize the tool with your API key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		configPath := filepath.Join(home, configFileName)
		err := os.WriteFile(configPath, []byte(args[0]), 0600)
		if err != nil {
			fmt.Printf("Error saving API key: %s", err)
			os.Exit(1)
		}
		fmt.Println("API key saved.")
	},
}
