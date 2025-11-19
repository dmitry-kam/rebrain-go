package homework_package

import "04-02-task/wordz"

var cities = []string{
	"Jakarta",
	"Paris",
	"Tokyo",
	"Berlin",
	"Petropavlovsk",
	"Vladivostok",
	"Novosibirsk",
	"Antalya",
	"Athens",
	"Kazan",
}

func City() string {
	originalWords, originalPrefix := wordz.Words, wordz.Prefix
	wordz.Words = cities
	wordz.Prefix = ""
	defer func() {
		wordz.Words = originalWords
		wordz.Prefix = originalPrefix
	}()

	return wordz.Random()
}
