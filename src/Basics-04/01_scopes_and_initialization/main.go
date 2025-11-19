package main

import (
	"fmt" // simple
	fatihcolor "github.com/fatih/color"
	_ "os"
	mycolor "rebrain-go-09/color"
	. "rebrain-go-09/color1"
	. "rebrain-go-09/color2"
	_ "rebrain-go-09/colorInit"
	"rebrain-go-09/wordz"
)

// Import types and their binary size impact:
// "package/path"   	   // Normal import: minimal size
// alias "package/path"    // Custom name, minimal size (tree-shaking applied)
// _ "package/path"        // Blank import (underscore): Init only, large size (whole package included)
// . "package/path"        // Dot import:  all functions are available without package name (No prefix),
//							  minimal size (namespace only change), conflicts are available (name collisions)

func main() {
	fmt.Println("Hello world")
	fatihcolor.Red("Hello world again")

	fatihcolor.Green(wordz.Hello)
	// Cannot use the unexported variable 'private' in the current package
	// ./main.go:12:20: undefined: wordz.private
	//color.Green(wordz.private)
	for i := 0; i < 5; i++ {
		fatihcolor.Yellow(wordz.Random())
	}

	mycolor.Greet()
	Greet1()
	Greet2()
}
