package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const sourceSrc = "src/Basics-03/06_defer"

func main() {
	defer logTime()()
	inLines := readIn()
	writeOut(inLines)
}

func readIn() []string {
	currentDir, _ := os.Getwd()
	filePath := filepath.Join(currentDir, sourceSrc, "data", "in.txt")

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(f)

	lines := make([]string, 0, 20)
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lines
}

func writeOut(lines []string) {
	bytesWritten, linesQty := 0, 0
	currentDir, _ := os.Getwd()
	filePath := filepath.Join(currentDir, sourceSrc, "data", "out.txt")

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer closeFile(f)

	for k, line := range lines {
		formattedLine := fmt.Sprintf("%d %s\n", k+1, line)

		l, err := f.WriteString(formattedLine)
		if err != nil {
			fmt.Println(err)
			return
		}
		bytesWritten += l
		linesQty++
	}

	fmt.Printf("%d bytes written successfully in %d lines\n", bytesWritten, linesQty)
}

func logTime() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Program execution time: %v\n", time.Since(start))
	}
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatal(err)
	} /*else {
		fmt.Printf("Closed file %s!\n", f.Name())
	}*/
}
