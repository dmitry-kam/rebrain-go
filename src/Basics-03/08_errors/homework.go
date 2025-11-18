package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	sourceSrc   = "src/Basics-03/08_errors"
	logFileName = "app.log"
)

func main() {
	logFile := setLogging()
	if logFile != nil {
		defer logFile.Close()
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("\u001B[1;31mFailed to get current directory: %v\u001B[0m\n", err)
		return
	}

	rFilePath := filepath.Join(currentDir, sourceSrc, "data", "in.txt")

	fR, err := os.Open(rFilePath)
	if err != nil {
		log.Printf("\u001B[1;31mFailed to open file %s: %v\u001B[0m\n", rFilePath, err)
		return
	}
	defer closeFile(fR)

	lineCount, err := countLines(fR)
	if err != nil {
		log.Printf("\u001B[1;31mFailed to count lines: %v\u001B[0m\n", err)
		return
	}

	fmt.Printf("Total strings: %d\n", lineCount)
}

func countLines(f *os.File) (int, error) {
	k := 0

	// r := bufio.NewScanner(f)  r.Scan() - no EOF error
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if len(line) > 0 {
			k++
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				//fmt.Printf("%s EOF\n", f.Name())
				break
			}
			return 0, fmt.Errorf("reading error: %w", err)
		}
	}

	return k, nil
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Println(err)
		return
	}
	log.Printf("\033[1;32mClosed file %s!\033[0m\n", f.Name())
}

func setLogging() *os.File {
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open log file: %v. Using console only.", err)
		return nil
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return logFile
}
