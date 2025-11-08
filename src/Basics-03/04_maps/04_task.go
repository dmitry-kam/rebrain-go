package main

import (
	"fmt"
	"strings"
)

func main() {
	var bookStats = map[string]map[string][]string{
		"Starikov Dmitrii": {
			"Ernst Jünger": {
				"Storm of Steel",
				"The Glass Bees",
				"Eumeswil",
			},
			"Friedrich Nietzsche": {
				"Human, All Too Human",
				"Twilight of the Idols",
				"Thus Spoke Zarathustra",
			},
			"Jack London": {
				"Martin Eden",
				"The Iron Heel",
				"The Sea-Wolf",
			},
		},
		"Ivanov Petr": {
			"Leo Tolstoy": {
				"War and Peace",
				"Anna Karenina",
			},
			"Newspapers": {
				"Izvestia",
				"Kommersant",
			},
			"Magazines": {
				"National Geographic",
			},
		},
		"Sidorova Maria": {
			"Agatha Christie": {
				"Murder on the Orient Express",
				"And Then There Were None",
			},
			"Fyodor Dostoevsky": {
				"Crime and Punishment",
				"The Brothers Karamazov",
			},
			"Journals": {
				"Nature",
				"Science",
			},
		},
		"Petrova Xenia": {},
	}

	readersWithBooks, booksIssued := 0, 0

	fmt.Println(strings.Repeat("#", 20))
	for reader, publications := range bookStats {
		printLnBold("Reader: " + reader)
		if len(publications) > 0 {
			readersWithBooks++
			booksIssued += len(publications)

			for pubType, items := range publications {
				fmt.Printf("  %s (%d): %v\n", pubType, len(items), strings.Join(items, "; "))
			}
		} else {
			fmt.Printf("  No books found\n")
		}
	}

	fmt.Println()
	fmt.Println(strings.Repeat("#", 20))
	fmt.Printf("Readers with books: %d\n", readersWithBooks)
	fmt.Printf("Books issued: %d\n", booksIssued)
}

func printLnBold(text string) {
	fmt.Printf("%s%s%s\n", "\033[1m", text, "\033[0m")
}
