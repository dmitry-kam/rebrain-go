package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	path := "f.txt"
	getFile(path)
	log()

	// Main func execution: 2
	// Deferred func call: 1
	i := 1
	defer prt(i)
	i++
	fmt.Printf("Main func execution: %d\n", i)

	defer func() {
		prt(i)
	}()

	i++
	fmt.Printf("Main func execution: %d\n", i)

	defer innerDefer()

	// read param from command line
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
}

func getFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		println("error opening file")
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Failed to close file: %v", err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

func log() {
	fmt.Println(0)
	defer fmt.Println(1) // executed third
	defer fmt.Println(2) // executed second
	defer fmt.Println(3) // executed first
	fmt.Println(5)
}

func prt(i int) {
	fmt.Printf("Deferred func call: %d\n", i)
}

func innerDefer() {
	fmt.Println("Defer1")
	defer fmt.Println(100)
	fmt.Println("Defer2")
}
