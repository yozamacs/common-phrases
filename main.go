package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var rawText string

	fileNamesStdin, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fileNamesStdin.Size() > 0 {
		fileNameBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		rawText = string(fileNameBytes)
	} else {
		fileNames := os.Args[1:]
		rawText = getRawText(fileNames)
	}

	results := findCommonPhrases(rawText)

	for phrase, count := range results {
		fmt.Printf("%v - %s", count, phrase)
	}
}
