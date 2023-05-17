package util

import (
	"fmt"

	"rsc.io/quote"
)

func About() string {
	AUTHOR := "radlad"
	return fmt.Sprintf("This package was build by %v", AUTHOR)
}

func GetQuote() string {
	return fmt.Sprintf("%v", quote.Glass())
}

func Hello(name string) string {
	return fmt.Sprintf("Hi, %v. Welcome!", name)	
}