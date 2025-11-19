package main

import (
	newcolor "09-task/color"
	h "09-task/homework_package"
	. "09-task/wordz"
	"fmt"
	"github.com/fatih/color"
	"github.com/huandu/xstrings"
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world")
	color.Red("Hello world again")
	fmt.Println(Hello)    //Вызов переменной из пакета wordz
	fmt.Println(Random()) //Вызов функции из пакета wordz

	color.Green(h.City())
	color.Yellow(h.Digit())

	newCity := h.City()

	color.New(color.FgBlack).Add(color.BgRed).Println(fmt.Sprintf("shuffled %s is %v", newCity, xstrings.Shuffle(newCity)))
}
