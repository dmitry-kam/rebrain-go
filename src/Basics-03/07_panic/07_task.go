package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const sourceSrc = "src/Basics-03/07_panic"

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("currentDir err", err)
		return
	}

	rFilePath := filepath.Join(currentDir, sourceSrc, "data", "in.txt")
	wFilePath := filepath.Join(currentDir, sourceSrc, "data", "data_out.txt")

	fR, err := os.Open(rFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closeFile(fR)

	fW, err := os.OpenFile(wFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closeFile(fW)

	defer func() {
		if panicErr := recover(); panicErr != nil {
			errMessage := fmt.Sprintf("%v", panicErr)

			switch true {
			case strings.Contains(errMessage, "parse error"):
				fmt.Printf("\033[1;31m%s\033[0m\n", errMessage)
			}

			printAlreadyWritten(wFilePath)
		}
	}()

	writeFormattedData(fR, fW)
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\033[1;32mClosed file %s!\033[0m\n", f.Name())
	}
}

func writeFormattedData(fR *os.File, fW *os.File) {
	k := 0
	s := bufio.NewScanner(fR)
	for s.Scan() {
		fString := getFormattedString(k, s.Text())
		_, err := fW.WriteString(fString)
		if err != nil {
			fmt.Println(err)
			return
		}
		k++
	}
	err := s.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func printAlreadyWritten(fileAddress string) {
	f, err := os.Open(fileAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer closeFile(f)
	s := bufio.NewScanner(f)

	for s.Scan() {
		err := s.Err()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(s.Text())
	}
}

func getFormattedString(order int, line string) string {
	fields := strings.Split(line, "|")
	if len(fields) == 3 && validateFields(fields...) {
		return fmt.Sprintf("%d\nName: %s\nAddress: %s\nCity: %s\n\n\n\n", order+1, fields[0], fields[1], fields[2])
	} else {
		panic(fmt.Sprintf("parse error: empty field on string %d\n", order+1))
	}
}

func validateFields(fields ...string) bool {
	for _, field := range fields {
		if 0 == len(field) {
			return false
		}
	}
	return true
}
