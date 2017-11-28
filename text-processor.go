package main

import (
	"regexp"
	"strings"
)

func cleanText(rawText string) []string {
	lowercaseText := strings.ToLower(rawText)
	newlinesRemoved := strings.Replace(lowercaseText, "\n", " ", -1)
	punctuationRegex := regexp.MustCompile(`[\w']+`)
	cleanTextList := punctuationRegex.FindAllString(newlinesRemoved, -1)
	return cleanTextList
}

func findCommonPhrases(text []string) map[string]int {
	return map[string]int{text[0]: 1}
}
