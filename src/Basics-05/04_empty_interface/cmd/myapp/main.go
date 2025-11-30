package main

import (
	"errors"
	"fmt"
	"io"
	"unsafe"
)

func f(w interface{}) error {
	writer, ok := w.(io.Writer)
	if !ok {
		return errors.New("Incorrect type")
	}
	_, err := writer.Write([]byte("Hello, world!"))
	return err
}

func processAnything(value interface{}) {
	fmt.Printf("Received: %v, type: %T\n", value, value)

	switch v := value.(type) {
	case int:
		fmt.Printf("Int: %d\n", v*2)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Bool: %t\n", !v)
	default:
		fmt.Printf("Unknown: %T\n", v)
	}
	fmt.Println("---")
}

func main() {
	s := struct{}{}
	var s1 struct{}
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(s1))

	processAnything(42)
	processAnything("Hello")
	processAnything(true)
	processAnything([]int{1, 2, 3})
	processAnything(struct{ Name string }{Name: "John"})

	emptyStructExample()
}

func emptyStructExample() {
	// any value = 1 byte
	mapWithBool := make(map[string]bool)
	mapWithBool["user1"] = true
	mapWithBool["user2"] = true

	// any value = 0 byte
	mapWithEmptyStruct := make(map[string]struct{})
	mapWithEmptyStruct["user1"] = struct{}{}
	mapWithEmptyStruct["user2"] = struct{}{}

	fmt.Printf("Size of bool: %d byte\n", unsafe.Sizeof(true))
	fmt.Printf("Size of empty struct: %d byte\n", unsafe.Sizeof(struct{}{}))

	if _, exists := mapWithEmptyStruct["user1"]; exists {
		fmt.Println("user1 has found!")
	}
}
