package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]

	var words []string
	var word string

	for _, ch := range arg {
		if ch == 32 {
			words = append(words, word)
			word = ""
		} else {
			word += string(ch)
		}
	}
	if len(word) > 0 {
		words = append(words, word)
	}
	var newWords []string
	for _, c := range words {
		if c == "" {
			continue
		} else {
			newWords = append(newWords, c)
		}
	}

	if len(newWords) > 1 {
		for i := 1; i <= len(newWords); i++ {

			if i == len(newWords) {
				Printn(newWords[0])
			} else {
				Printn(newWords[i] + " ")
			}
		}
	} else {
		println(newWords[0])
	}

}
func Printn(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}

}
