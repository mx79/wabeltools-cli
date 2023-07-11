/*
Copyright © 2023 WABEL GROUP <m.lesage@wabelgroup.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

const (
	ner        = "ner"
	posTagging = "pos-tagging"
	sentiment  = "sentiment"
	segmenter  = "segmenter"
	rake       = "rake"
	stemming   = "stemming"
	stopwords  = "stopwords"
	wer        = "wer"
)

// nlpCmd represents the nlp command
var nlpCmd = &cobra.Command{
	Use:   "nlp",
	Short: "Natural Language Processing",
	Long: `Natural Language Processing commands stand for natural language processing tasks: text analysis, sentiment, rake score, stemming, stopwords, etc.

For example:

wabeltools nlp sentiment "I love you"
wabeltools nlp rake "I love you"
wabeltools nlp stemming "I like to eat apples"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Missing required subcommand.\n" +
			"Please use one of the following subcommands: ner, pos-tagging, sentiment, segmenter, rake, stemming, stopwords, wer")
	},
}

func init() {
	nlpCmd.AddCommand(nlpNerCmd, nlpPosTaggingCmd, nlpSentimentCmd, nlpSegmenter, nlpRakeCmd, nlpStemmingCmd, nlpStopwordsCmd, nlpWerCmd)
}

// nlpRequest is a helper function to make a request to the NLP API
func nlpRequest(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-KEY", viper.GetString("apikey"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// nlpNerCmd represents the nlp ner command
var nlpNerCmd = &cobra.Command{
	Use:   "ner",
	Short: "Named entity recognition",
	Long: `Named entity recognition is the task of extracting named entities from the text and classifying them into pre-defined categories.

For example:

wabeltools nlp ner "My string to apply NER on"
=> {}`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Please provide one string to analyze with Named entity recognition")
		}
		url := fmt.Sprintf("%s/%s?text=%s", nlpURL, ner, args[0])
		res, err := nlpRequest(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(res))
	},
}

// nlpPosTaggingCmd represents the nlp posTagging command
var nlpPosTaggingCmd = &cobra.Command{
	Use:   "pos-tagging",
	Short: "Pos tagging",
	Long: `Pos tagging is the task of marking up a word in a text as corresponding to a particular part of speech, based on both its definition and its context.

For example:

wabeltools nlp pos-tagging "My string to apply pos tagging on"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Please provide one string to apply pos tagging on")
		}
		url := fmt.Sprintf("%s/%s?text=%s", nlpURL, posTagging, args[0])
		res, err := nlpRequest(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(res))
	},
}

// nlpSentimentCmd represents the nlp sentiment command
var nlpSentimentCmd = &cobra.Command{
	Use:   "sentiment",
	Short: "Sentiment analysis",
	Long: `Sentiment analysis is the task of analyzing a string of text to determine the sentiment or opinion of the text.

For example:

wabeltools nlp sentiment "My string to analyze sentiment on"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Please provide one string to analyze sentiment on")
		}
		url := fmt.Sprintf("%s/%s?text=%s", nlpURL, sentiment, args[0])
		res, err := nlpRequest(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(res))
	},
}

// nlpSegmenter represents the nlp segmenter command
var nlpSegmenter = &cobra.Command{
	Use:   "segmenter",
	Short: "Segmentation",
	Long: `Segmentation is the task of dividing a string of written language into its component parts (segments).

For example:

wabeltools nlp segmenter "My string to segment"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Please provide one string to segment")
		}
		url := fmt.Sprintf("%s/%s?text=%s", nlpURL, segmenter, args[0])
		res, err := nlpRequest(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(res))
	},
}

// nlpRakeCmd represents the nlp rake command
var nlpRakeCmd = &cobra.Command{
	Use:   "rake",
	Short: "Rake score",
	Long: `Ranking of keywords extracted from a text using the RAKE algorithm.

For example:

waebeltools nlp rake "My string to get rake score"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Please provide one string to get rake score")
		}
		url := fmt.Sprintf("%s/%s?text=%s", nlpURL, rake, args[0])
		res, err := nlpRequest(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(res))
	},
}

// nlpStemmingCmd represents the nlp stemming command
var nlpStemmingCmd = &cobra.Command{
	Use:   "stemming",
	Short: "Stemming",
	Long: `Stemming is the process of reducing inflected (or sometimes derived) words to their word stem, base or root form—generally a written word form.

For example:

wabeltools nlp stemming "My string to stem"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Please provide one string to stem")
		}
		url := fmt.Sprintf("%s/%s?text=%s", nlpURL, stemming, args[0])
		res, err := nlpRequest(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(res))
	},
}

// nlpStopwordsCmd represents the nlp stopwords command
var nlpStopwordsCmd = &cobra.Command{
	Use:   "stopwords",
	Short: "Stopwords removal",
	Long: `Stopwords removal is the task of removing stopwords (redundant words of a language) from a string of text. These words usually don't add any value to the text.

For example:

wabeltools nlp stopwords "My string to remove stopwords from"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Please provide one string to remove stopwords from")
		}
		url := fmt.Sprintf("%s/%s?text=%s", nlpURL, stopwords, args[0])
		res, err := nlpRequest(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(res))
	},
}

// nlpWerCmd represents the nlp wer command
var nlpWerCmd = &cobra.Command{
	Use:   "wer",
	Short: "Word error rate",
	Long: `Word error rate is the task of comparing two strings and calculating the word error rate between them.

For example:

wabeltools nlp wer "My string to compare" "My string to compare with"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Please provide two strings to compare with word error rate")
		}
		url := fmt.Sprintf("%s/%s?text1=%s&text2=%s", nlpURL, wer, args[0], args[1])
		res, err := nlpRequest(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(res))
	},
}
