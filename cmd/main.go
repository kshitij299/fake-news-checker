package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kshitij299/fake-news-checker/internal/pkg/NewsScanner"
)

var (
	newsScanner NewsScanner.INewsScanner
	reader      *bufio.Reader = bufio.NewReader(os.Stdin)
)

func init() {
	newsScanner = NewsScanner.NewGoogleScanner()
}

func main() {
	fmt.Println("please provide the API key:")
	for {
		key, _ := reader.ReadString('\n')
		// convert CRLF to LF
		key = strings.Replace(key, "\n", "", -1)
		isValid := isValidApiKey(key)
		if !isValid {
			fmt.Println("api key is invalid, please retry with a valid key:")
			continue
		} else {
			newsScanner.SetApiKey(key)
			break
		}
	}
	var (
		err    error
		isFake bool
	)
	for {
		fmt.Println("please enter your query:")
		news, _ := reader.ReadString('\n')
		// convert CRLF to LF
		news = strings.Replace(news, "\n", "", -1)
		isValid := isValidNews(news)
		if !isValid {
			fmt.Println("news is invalid, please retry with a valid news")
			continue
		}
		isFake, err = newsScanner.IsFake(news)
		if err != nil {
			fmt.Println("could not check status, error:", err.Error())
		} else {
			if isFake {
				fmt.Println("news is fake")
			} else {
				fmt.Println("news looks genuine")
			}
		}
	}
}

//user input validations
//isValidApiKey validates API key
func isValidApiKey(key string) bool {
	//only basic check for now
	if key == "" {
		return false
	}
	return true
}

//isValidNews validates query news
func isValidNews(news string) bool {
	//only basic check for now
	if news == "" {
		return false
	}
	return true
}
