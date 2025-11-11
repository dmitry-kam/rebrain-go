package main

import "fmt"

func main() {
	defer fmt.Println("main: defer")

	//someFunc()
	fmt.Println("main: after someFunc call")

	err := someFunc2()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("after someFunc")
}

func someFunc() {
	defer fmt.Println("someFunc: defer")

	panic("something get wrong")
}

func someFunc2() (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			switch panicErr {
			case "regular error":
				err = fmt.Errorf("application error")
			default:
				// bad practice
				panic("critical")
			}
		}
	}()

	panic("regular error")
}
