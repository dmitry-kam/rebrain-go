package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

type wordLogger func(w ...string) (cnt int)

func main() {
	customer := map[string]string{
		"name":     "John",
		"lastName": "Smith"}

	if lastName, isExist := customer["lastName"]; isExist {
		fmt.Printf("Hello Mr. %s!\n", lastName)
	}

	////////////////
	if _, isExist := customer["fieldName"]; !isExist {
		fmt.Println("Empty field!")
	}

	if lastName, lastNameExist := customer["lastName"]; !lastNameExist {
		fmt.Println("Empty lastName!")
	} else if name, nameExist := customer["name"]; !nameExist {
		fmt.Println("Empty name")
	} else {
		fmt.Printf("Hello, %s %s!\n", name, lastName)
		fmt.Printf("nameExist, lastNameExist are %v, %v\n", nameExist, lastNameExist)

		switch true {
		case nameExist && false:
			fmt.Println("switch true 1")
		case nameExist && true:
			fmt.Println("switch true 2")
		case nameExist && lastNameExist:
			fmt.Println("switch true 2")
		}
	}

	userRole := "admin"
	//userRole = "not_admin"
	switch userRole {
	case "Admin":
		fmt.Println("Access granted") // the same
	case "admin":
		fmt.Println("Access granted")
	//case "not_admin": // nothing done
	case "user":
		fmt.Println("Access granted for user")
	default:
		fmt.Println("Access denied")
	}

	userRole = "guest"
	switch userRole {
	case "admin", "user":
		fmt.Println("Access granted")
	case "guest":
		fallthrough // next condition without check
	default:
		fmt.Println("Access denied!")
	}

	level := "error"
	switch level {
	case "error":
		fmt.Println("Sending email alert")
		fallthrough
	case "warning":
		fmt.Println("Writing to log file") // executed for "error", "warning"
		fallthrough
	case "info":
		fmt.Println("Displaying message") // executed for any level
		//break
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	var isBreak = false
	for {
		if isBreak {
			fmt.Println("Break!")
			break
		} else {
			r := rand.IntN(1024)
			fmt.Println(r)
			if r > 777 {
				isBreak = true
			}
		}
	}

	///////////////
	isBreak = true
Loop1:
	for {
		fmt.Println("Loop 1")
	Loop2:
		for {
			fmt.Println("Loop 2")
			if isBreak {
				fmt.Println("break Loop 2")
				break Loop2
			}
		}
		if isBreak {
			fmt.Println("break Loop 1")
			break Loop1
		}
	}
	fmt.Println("After Loop")

	////////////
	iterate := true
	for iterate {
		fmt.Println("Iteration")
		iterate = false
	}

	m := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	a, _ := foo(1, "2")
	fmt.Println(a)

	s := sum(1, 2)
	fmt.Println(s)

	w, ll := words("aaa", "bbb", "ccc")
	fmt.Println(w, ll)

	//////////////////////////////////////

	printWord := func(w ...string) {
		for _, word := range w {
			fmt.Println(word)
		}
	}

	func(w ...string) {
		printWord(w...)
	}("str1", "str2")

	printWord1 := func(w ...string) (cnt int) {
		for _, word := range w {
			fmt.Println(word)
		}

		return len(w)
	}

	func(printer wordLogger, w ...string) {
		fmt.Printf("Count words: %d\n", printer(w...))
	}(printWord1, "str3", "str4")

	printWord1("str5", "str6")

	printer := generatePrinter()

	printer("word1", "word2", "word3", "word4")

	////////////////////////////////
	// closure

	sequence := sequenceGen(4)
	fmt.Println(sequence()) // 4
	fmt.Println(sequence()) // 8
	fmt.Println(sequence()) // 12
	fmt.Println(sequence()) // 16
	fmt.Println(sequence()) // 20
	fmt.Println(sequence()) // 24
	fmt.Println(sequence()) // 28
	fmt.Println(sequence()) // 32

	////////////////////////////////

	numbers := []int{1, 11, -5, 8, 2, 0, 12}
	sort.Ints(numbers)
	fmt.Println("Sorted:", numbers)

	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= 7
	})
	fmt.Println("The first number >= 7 is at index:", index)
	fmt.Println("The first number >= 7 is:", numbers[index])
}

func foo(p1 int, p2 string) (int, string) {
	return p1, p2
}

func sum(p1 int, p2 int) (sum int) {
	sum = p1 + p2
	return
}

func words(w ...string) ([]string, int) {
	return w, len(w)
}

func sequenceGen(n int) func() int {
	i := 0
	return func() int {
		i += n
		return i
	}
}

func generatePrinter() wordLogger {
	return func(w ...string) (cnt int) {
		for _, word := range w {
			fmt.Println("---print:" + word)
		}

		return len(w)
	}
}
