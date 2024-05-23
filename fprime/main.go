package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		return
	}

	nb := Atoi(os.Args[1])

	if nb == 0 || nb == 1 {
		return
	}
	var factors []string

	for i := 2; i*i <= nb; i++ {
		for nb%i == 0 {
			factors = append(factors, Itoa(i))
			nb /= i
		}

	}
	if nb > 2 {
		factors = append(factors, Itoa(nb))
	}
	for i, sl := range factors {
		if i != len(factors)-1 {
			Printn(sl)
			Printn("*")
		} else {
			Printn(sl)
		}
	}
	Printn("\n")

}
func Atoi(s string) int {
	num := 0

	for _, v := range s {
		n := int(v - 48)

		if n < 0 || n > 9 {
			return 0
		}
		num = (num * 10) + n

	}
	return num

}
func Itoa(n int) string {
	s := ""
	if n == 0 {
		s = string(rune(48)) + s
	}
	for n > 0 {
		w := n % 10
		s = string(rune(w+48)) + s
		n = n / 10
	}

	return s

}
func Printn(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}

}
