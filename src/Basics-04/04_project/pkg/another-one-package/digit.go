package another_one_package

import (
	"gopackages-layout/pkg/wordz"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}

func Digit() string {
	originalWords, originalPrefix := wordz.Words, wordz.Prefix
	wordz.Words = digits
	wordz.Prefix = ""
	defer func() {
		wordz.Words = originalWords
		wordz.Prefix = originalPrefix
	}()

	return wordz.Random()
}
