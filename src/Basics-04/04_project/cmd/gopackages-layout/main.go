package main

import (
	"fmt"
	"github.com/huandu/xstrings"
	anotherpackage "gopackages-layout/pkg/another-one-package"
	newcolor "gopackages-layout/pkg/color"
	. "gopackages-layout/pkg/wordz"
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world")
	// Re
	fmt.Printf("\033[1;31mHello world again\033[0m\n")
	fmt.Println(Hello)
	fmt.Println(Random())

	// Green
	fmt.Printf("\033[1;32m%s\033[0m\n", anotherpackage.City())
	// Yellow
	fmt.Printf("\033[1;33m%s\033[0m\n", anotherpackage.Digit())

	newCity := anotherpackage.City()

	// Black on red
	fmt.Printf("\033[1;30;41mshuffled %s is %v\033[0m\n", newCity, xstrings.Shuffle(newCity))
}
