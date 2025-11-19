package main

import (
	newcolor "04-02-task/color"
	. "04-02-task/homework_package"
	. "04-02-task/wordz"
	"fmt"
	"github.com/huandu/xstrings"
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world")
	// Re
	fmt.Printf("\033[1;31mHello world again\033[0m\n")
	fmt.Println(Hello)
	fmt.Println(Random())

	// Green
	fmt.Printf("\033[1;32m%s\033[0m\n", City())
	// Yellow
	fmt.Printf("\033[1;33m%s\033[0m\n", Digit())

	newCity := City()

	// Black on red
	fmt.Printf("\033[1;30;41mshuffled %s is %v\033[0m\n", newCity, xstrings.Shuffle(newCity))
}
