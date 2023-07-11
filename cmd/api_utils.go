/*
Copyright © 2023 WABEL GROUP <m.lesage@wabelgroup.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	isValidAPIKey = "/is-valid-api-key"
	getCosts      = "/costs"
	getServices   = "/services"
	getTokens     = "/tokens"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [apikey]",
	Short: "Initialize the Wabel toolkit with your API key from wabeltools.com",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		configPath := filepath.Join(home, configFileName)

		// Make a request to the API to check if the API key is valid
		client := &http.Client{}
		url := fmt.Sprintf("%s%s", baseURL, isValidAPIKey)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("X-API-KEY", args[0])
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error querying Wabel Tools API: %s", err)
			os.Exit(1)
		}

		// Check if the API key is valid
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error validating API key: %s", resp.Status)
			os.Exit(1)
		}

		// Save the API key to the config file
		if err = os.WriteFile(configPath, []byte(args[0]), 0600); err != nil {
			fmt.Printf("Error saving API key: %s", err)
			os.Exit(1)
		}
		fmt.Println("API key saved.")
	},
}

var costsCmd = &cobra.Command{
	Use:   "costs",
	Short: "Costs commands lists the costs of each service",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Make a request to the API to check if the API key is valid
		client := &http.Client{}
		url := fmt.Sprintf("%s%s", baseURL, getCosts)
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error querying Wabel Tools API: %s", err)
			os.Exit(1)
		}

		// Check if the API key is valid
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error getting costs, bad status code: %s", resp.Status)
			os.Exit(1)
		}

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %s", err)
			os.Exit(1)
		}

		fmt.Println(string(body))
	},
}

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Services command prints the list of available services",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Make a request to the API to check if the API key is valid
		client := &http.Client{}
		url := fmt.Sprintf("%s%s", baseURL, getServices)
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error querying Wabel Tools API: %s", err)
			os.Exit(1)
		}

		// Check if the API key is valid
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error getting services, bad status code: %s", resp.Status)
			os.Exit(1)
		}

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %s", err)
			os.Exit(1)
		}

		fmt.Println(string(body))
	},
}

var tokensCmd = &cobra.Command{
	Use:   "tokens",
	Short: "Tokens commands allows you to see your remaining tokens for the day",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Make a request to the API to check if the API key is valid
		client := &http.Client{}
		url := fmt.Sprintf("%s%s", baseURL, getTokens)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("X-API-KEY", args[0])
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error querying Wabel Tools API: %s", err)
			os.Exit(1)
		}

		// Check if the API key is valid
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error validating API key: %s", resp.Status)
			os.Exit(1)
		}

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %s", err)
			os.Exit(1)
		}

		fmt.Println(string(body))
	},
}
