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

	cleanText := cleanText(rawText)
	phraseList := findCommonPhrases(cleanText)

	for _,phrase:=range phraseList {
		fmt.Printf("%v - %s\n", phrase.count, phrase.phrase)
	}

}
