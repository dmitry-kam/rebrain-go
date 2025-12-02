package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type errNotFound struct {
	message string
}

func (e *errNotFound) Error() string {
	return e.message
}

var (
	errNotFoundMessage = "not found"
	errNotFound1       = errors.New("not found")
	values             = []string{"aaa", "bbb", "1"}
)

func searchAndConvert(expectedValue string) (int, error) {
	for _, v := range values {
		if v == expectedValue {
			convertedValue, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}

			return convertedValue, nil
		}
	}

	return 0, errNotFound1
}

func searchAndConvert1(expectedValue string) (int, error) {
	for _, v := range values {
		if v == expectedValue {
			convertedValue, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}

			return convertedValue, nil
		}
	}

	return 0, &errNotFound{message: errNotFoundMessage}
}

func searchAndConvert2(expectedValue string) (int, error) {
	return 0, fmt.Errorf("searchAndConvert error: %w", &errNotFound{message: errNotFoundMessage})
}

func main() {
	res, err := searchAndConvert("1")
	if err != nil {
		if err == errNotFound1 {
			log.Println("value not found")
			return
		}

		log.Fatal(err)
	}
	log.Println(res)

	///////////// struct errNotFound implements Error interface
	res, err = searchAndConvert1("1")
	if _, ok := err.(*errNotFound); ok {
		log.Println("value not found")
		return
	}
	log.Println(res)

	///////////// wrapped error
	/*res, err = searchAndConvert2("2")
	if _, ok := errors.Unwrap(err).(*errNotFound); ok {
		log.Println("value not found")
		return
	}
	log.Println(res)*/

	///////////// as
	res, err = searchAndConvert2("2")
	if err != nil {
		var enf *errNotFound
		if errors.As(err, &enf) {
			log.Println("value not found")
			return
		}

		log.Fatal(err)
	}
	log.Println(res)
}
