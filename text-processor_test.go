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
so that way your friends don't get really jealous and all of that. I hope you
enjoy this new feature!
- Chao
`

var expectedCleanText = []string{
	"hey", "everyone", "i", "just", "wanted", "to", "make", "an", "announcement",
	"that", "we", "are", "no", "longer", "going", "to", "tell", "your", "friends", "that", "you",
	"have", "unfriended", "them", "this", "is", "a", "great", "improvement", "so", "that", "way",
	"your", "friends", "don't", "get", "really", "jealous", "and", "all", "of", "that", "i", "hope",
	"you", "enjoy", "this", "new", "feature", "chao",
}

func (tv *TextValidationSuite) TestTextCleaning(c *C) {
	actualCleanText := cleanText(RelativelyLongTextRaw)
	c.Assert(actualCleanText, DeepEquals, expectedCleanText)
}
