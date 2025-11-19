package homework_package

import (
	"04-02-task/wordz"
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
