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

func (l limitExceeded) Error() string {
	return fmt.Sprintf("`%s`, limit: %d, last string: `%s`", l.message, l.limit, l.lastString)
}

func main() {
	limit := 25
	fR, err := getInFile()
	var fileNotFoundError *os.SyscallError
	var filePathError *fs.PathError

	if err != nil {
		if errors.As(err, &fileNotFoundError) {
			log.Printf("currentDir is undefined: %s\n", err.Error())
		} else if errors.As(err, &filePathError) {
			log.Printf("can`t open file: %s\n", err.Error())
		} else {
			log.Println(err.Error())
		}
		return
	}
	defer closeFile(fR)

	lineCount, err := countLines(fR, limit)
	if err != nil {
		if errors.As(err, &limitExceeded{}) {
			fmt.Printf("\u001B[1;31mstring count exceed limit, please set another limit:\u001B[0m %s\n", err.Error())
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
			return counter, limitExceeded{message: "limit has been reached", limit: limit, lastString: line}
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				//fmt.Printf("%s EOF\n", f.Name())
				break
			}
			return 0, fmt.Errorf("(wrapped reading error: %w)", err)
		}
	}

	return counter, nil
}

func getInFile() (*os.File, error) {
	const sourceSrc = "src/Basics-03/08_errors"
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	rFilePath := filepath.Join(currentDir, sourceSrc, "data", "in.txt")
	fR, err := os.Open(rFilePath)
	if err != nil {
		return nil, fmt.Errorf("(wrapped getInFile err: %w)", err)
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
