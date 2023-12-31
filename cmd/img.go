/*
Copyright © 2023 WABEL GROUP <m.lesage@wabelgroup.com>
*/
package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	local  = "/local"
	remote = "/remote"
)

// imgCmd represents the image command
var imgCmd = &cobra.Command{
	Use:   "image [local|remote]",
	Short: "Image commands stand for image processing tasks: resizing, compressing, OCR, etc.",
	Long: `Image commands stand for image processing tasks: resizing, compressing, OCR, etc.

For example:

wabeltools image --resize 100x100 "image.jpg"
wabeltools image local --compress "path/to/image.jpg"
wabeltools image remote --many --compress --urls="https://www.wabel.com/images/logo.png,https://www.wabel.com/images/logo.png"`,
	Args: cobra.MinimumNArgs(1),
}

func init() {
	imgLocalCmd.Flags().IntP("quality", "q", 50, "Quality of the processed image")
	imgRemoteCmd.Flags().IntP("quality", "q", 50, "Quality of the processed image")

	imgCmd.AddCommand(imgLocalCmd, imgRemoteCmd)
}

// imgLocalCmd represents the image local command
var imgLocalCmd = &cobra.Command{
	Use:   "local [image_path]",
	Short: "Local commands stand for image processing tasks on local images: resizing, compressing, OCR, etc.",
	Long: `Local commands stand for image processing tasks on local images: resizing, compressing, OCR, etc.

For example:

wabeltools image local --resize 100x100 "image.jpg"
wabeltools image local --compress "path/to/image.jpg" "path/to/other/image.jpg"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the user provided a list of URLs
		if len(args) == 0 {
			fmt.Println("Please provide at list one URL as argument\n" +
				"Example: wabeltools image local --compress \"/path/to/my_image.png\"")
			os.Exit(1)
		}

		// Create a buffer to store our request body
		var requestBody bytes.Buffer

		// Create a multipart writer
		multiPartWriter := multipart.NewWriter(&requestBody)

		// Loop over the files and add each one to the multipart form data
		for _, path := range args {
			file, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()

			// Get the basename of the file path (i.e., the file name)
			fileName := filepath.Base(path)

			// Create a new form-data header with the file name
			part, err := multiPartWriter.CreateFormFile("img", fileName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Copy the file into the form-data part
			if _, err := io.Copy(part, file); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// We have to close the multipart writer after we added all the files
		if err := multiPartWriter.Close(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Now we can create the request with our body, note that we're also setting the Content-Type header here
		var manyOrNot string
		if len(args) > 1 {
			manyOrNot = "/many"
		}
		url := fmt.Sprintf("%s%s%s", imgURL, local, manyOrNot)
		req, err := http.NewRequest("POST", url, &requestBody)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.Header.Add("X-API-KEY", viper.GetString("apikey"))
		req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

		// Add query parameters
		quality, _ := cmd.Flags().GetInt("quality")
		params := map[string]string{
			"quality": strconv.Itoa(quality),
		}
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()

		// Finally, we can send our request
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer response.Body.Close()

		// Read body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Save the processed zip file
		if len(args) > 1 {
			filename := "processed_" + strconv.Itoa(len(args)) + "_images.zip"
			err = os.WriteFile(filename, body, 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Processed images saved as '%s'\n", filename)
			return
		}

		// Save the processed image
		filename := "processed_" + filepath.Base(args[0])
		err = os.WriteFile(filename, body, 0666)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Processed image saved as '%s'\n", filename)
	},
}

// imgRemoteCmd represents the image remote command
var imgRemoteCmd = &cobra.Command{
	Use:   "remote [urls]",
	Short: "Remote commands stand for image processing tasks on remote images: resizing, compressing, OCR, etc.",
	Long: `Remote commands stand for image processing tasks on remote images: resizing, compressing, OCR, etc.

For example:

wabeltools image remote --resize 100x100 "https://www.wabel.com/images/logo.png"
wabeltools image remote --compress "https://www.wabel.com/images/logo.png" "https://www.wabel.com/images/logo.png"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the user provided a list of URLs
		if len(args) == 0 {
			fmt.Println("Please provide at list one URL as argument\n" +
				"Example: wabeltools image remote --compress \"https://www.wabel.com/images/logo.png\"")
			os.Exit(1)
		}

		// Now we can create the request
		var manyOrNot string
		if len(args) > 1 {
			manyOrNot = "/many"
		}
		url := fmt.Sprintf("%s%s%s", imgURL, remote, manyOrNot)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.Header.Add("X-API-KEY", viper.GetString("apikey"))

		// Add query parameters
		quality, _ := cmd.Flags().GetInt("quality")
		params := map[string]string{
			"quality": strconv.Itoa(quality),
		}
		if len(args) > 1 {
			params["urls"] = strings.Join(args, ",")
		} else {
			params["url"] = args[0]
		}
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()

		// Finally, we can send our request
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer response.Body.Close()

		// Read body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Save the processed zip file
		if len(args) > 1 {
			filename := "processed_" + strconv.Itoa(len(args)) + "_images.zip"
			err = os.WriteFile(filename, body, 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Processed images saved as '%s'\n", filename)
			return
		}

		// Save the processed image
		filename := "processed_" + filepath.Base(args[0])
		err = os.WriteFile(filename, body, 0666)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Processed image saved as '%s'\n", filename)
	},
}
