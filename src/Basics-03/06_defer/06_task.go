package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const sourceSrc = "src/Basics-03/06_defer"

func main() {
	defer logTime()()

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("currentDir err", err)
		return
	}

	bytesWritten, linesQty := 0, 0
	rFilePath := filepath.Join(currentDir, sourceSrc, "data", "in.txt")
	wFilePath := filepath.Join(currentDir, sourceSrc, "data", "out.txt")

	fR, err := os.Open(rFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closeFile(fR)

	fW, err := os.OpenFile(wFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closeFile(fW)

	s := bufio.NewScanner(fR)
	for s.Scan() {
		err = s.Err()
		if err != nil {
			fmt.Println(err)
			return
		}

		formattedLine := fmt.Sprintf("%d %s\n", linesQty+1, s.Text())

		l, err := fW.WriteString(formattedLine)
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
		fmt.Println(err)
	} /*else {
		fmt.Printf("Closed file %s!\n", f.Name())
	}*/
}
