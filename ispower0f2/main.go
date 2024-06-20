package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	newArg := 0
	for _, v := range arg {
		n := int(v - 48)

		newArg = (newArg * 10) + n

	}

	fmt.Println(isPowerof2(newArg))
}

func isPowerof2(newArg int) bool {
	return newArg%2 == 0
}

// Seth  Solution
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	numb := Atoi(os.Args[1])

	var ans string 
	
	if isPowerOf2(numb) {
		ans = "true"
	} else {
		ans = "false"
	}
	
	for _, v := range ans {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func isPowerOf2(numb int) bool {
	return numb > 0 && (numb&(numb - 1)) == 0
}
	
func Atoi(s string) int {
	var number int
	sign := 1

	for idx, c := range s {
		if idx == 0 && c == '-' {
            sign = -1
        } else if c == '+' && idx == 0 {
			sign = 1
		} else if c >= '0' || c <= '9' {
            number = number * 10 + int(c - '0')
        } else {
			return 0
		}
	}
	return sign * number
}
// ombima solution

// package main

// import (
// "os"

// "github.com/01-edu/z01"
// )

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]

	n := BasicAtoi(args)
	if isPowerOf2(n) {
		printStr("true")
	} else {
		printStr("false")
	}
}

func isPowerOf2(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

func BasicAtoi(s string) int {
	var number int
	for _, ch := range s {
		number = number*10 + int(ch-'0')
	}
	return number
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
