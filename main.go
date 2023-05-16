package sample

import (
	"fmt"
	"rsc.io/quote"
)

func About() string {
	AUTHOR := "radlad"
	return fmt.Sprintf("This package was build by %v", AUTHOR)
}

func GetQuote() {
	fmt.Println(quote.Glass())
}

func Hello(name string) string {
	return fmt.Sprintf("Hi, %v. Welcome!", name)	
}