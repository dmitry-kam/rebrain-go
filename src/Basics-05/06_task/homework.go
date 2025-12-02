package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type limitExceeded struct {
	message    string
	limit      int
	lastString string
}

func (l *limitExceeded) Error() string {
	return fmt.Sprintf("%s, limit: %d, last string: %s", l.message, l.limit, l.lastString)
}

func main() {
	limit := 25
	fR, err := getInFile()
	if err != nil {
		return
	}
	defer closeFile(fR)

	lineCount, err := countLines(fR, limit)
	if err != nil {
		if _, ok := err.(*limitExceeded); ok {
			fmt.Println("string count exceed limit, please read another file =) err: ", err.Error())
			return
		}
		// other errors
		log.Printf("\u001B[1;31mFailed to count lines: %v\u001B[0m\n", err)
		return
	}

	fmt.Printf("Total strings: %d\n", lineCount)
}

func countLines(f *os.File, limit int) (int, error) {
	counter := 0

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if len(line) > 0 {
			counter++
		}
		if counter == limit {
			return counter, &limitExceeded{message: "limit has been reached", limit: limit, lastString: line}
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				//fmt.Printf("%s EOF\n", f.Name())
				break
			}
			return 0, fmt.Errorf("reading error: %w", err)
		}
	}

	return counter, nil
}

func getInFile() (*os.File, error) {
	const sourceSrc = "src/Basics-03/08_errors"
	currentDir, err := os.Getwd()
	if _, ok := err.(*os.SyscallError); ok {
		log.Println("currentDir is undefined")
		return nil, err
	}

	rFilePath := filepath.Join(currentDir, sourceSrc, "data", "in.txt")

	fR, err := os.Open(rFilePath)
	if _, ok := err.(*fs.PathError); ok {
		log.Println("Can`t open ", rFilePath)
		return nil, err
	}
	return fR, nil
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Println(err)
		return
	}
	log.Printf("\033[1;32mClosed file %s!\033[0m\n", f.Name())
}
