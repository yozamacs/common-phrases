package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

type TextValidationSuite struct{}

var _ = Suite(&TextValidationSuite{})

func Test(t *testing.T) { TestingT(t) }

var RelativelyLongTextRaw = `
Hey everyone! I just wanted to make an announcement that we are no longer
going to tell your friends that you have unfriended them. This is a great improvement
so that way your friends don't get really jealous and all of that. I just wanted you
to enjoy this new feature!
- Chao
`

var expectedCleanText = []string{
	"hey", "everyone", "i", "just", "wanted", "to", "make", "an", "announcement",
	"that", "we", "are", "no", "longer", "going", "to", "tell", "your", "friends", "that", "you",
	"have", "unfriended", "them", "this", "is", "a", "great", "improvement", "so", "that", "way",
	"your", "friends", "don't", "get", "really", "jealous", "and", "all", "of", "that", "i", "just",
	"wanted", "you", "to", "enjoy", "this", "new", "feature", "chao",
}

var repeatedPhrases = []string{
	"is", "repeated", "is", "repeated", "is", "repeated", "is", "repeated", "is", "tons",
}

var expectedRepeatedPhrases = PhraseList{
	Phrase{count: 4, phrase: "is repeated is"}, Phrase{count: 3, phrase: "repeated is repeated"}, Phrase{count: 1, phrase: "repeated is tons"},
}

func (tv *TextValidationSuite) TestTextCleaning(c *C) {
	actualCleanText := cleanText(RelativelyLongTextRaw)
	c.Assert(actualCleanText, DeepEquals, expectedCleanText)
}

func (tv *TextValidationSuite) TestFindCommonPhrases(c *C) {
	actualRepeatedPhrases := findCommonPhrases(repeatedPhrases)
	c.Assert(actualRepeatedPhrases, DeepEquals, expectedRepeatedPhrases)
}
