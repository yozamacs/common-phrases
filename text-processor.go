package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type Phrase struct {
	count  int
	phrase string
}

type PhraseList []Phrase

func (pl PhraseList) Len() int           { return len(pl) }
func (pl PhraseList) Swap(i, j int)      { pl[i], pl[j] = pl[j], pl[i] }
func (pl PhraseList) Less(i, j int) bool { return pl[i].count < pl[j].count }

func getRawText(fileNames []string) string {
	buffer := bytes.NewBufferString("")
	for _, fileName := range fileNames {
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
			continue
		}
		_, err = buffer.Write(fileBytes)
		if err != nil {
			fmt.Printf("error writing to buffer: %v", err)
			continue
		}
	}
	return buffer.String()
}

func cleanText(rawText string) []string {
	lowercaseText := strings.ToLower(rawText)
	newlinesRemoved := strings.Replace(lowercaseText, "\n", " ", -1)
	punctuationRegex := regexp.MustCompile(`[\w']+`)
	cleanTextList := punctuationRegex.FindAllString(newlinesRemoved, -1)
	return cleanTextList
}

func findCommonPhrases(text []string) PhraseList {
	phraseMap := make(map[string]int)
	for i := 0; i < len(text)-2; i++ {
		phrase := text[i] + " " + text[i+1] + " " + text[i+2]
		if count, ok := phraseMap[phrase]; !ok {
			phraseMap[phrase] = 1
		} else {
			phraseMap[phrase] = count + 1
		}
	}

	var phraseList PhraseList
	for phrase, count := range phraseMap {
		tempPhrase := Phrase{
			count:  count,
			phrase: phrase,
		}
		phraseList = append(phraseList, tempPhrase)
	}
	sort.Sort(sort.Reverse(phraseList))
	return phraseList
}
