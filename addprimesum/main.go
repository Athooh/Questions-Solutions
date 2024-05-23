package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 || os.Args[1][0] == '-' {
		fmt.Println(0)
		return
	}

	nb := Atoi(os.Args[1])
	if nb == 0 || nb == 1 {
		fmt.Println(0)
		return
	}
	sum := 0

	for i := nb; i > 1; i-- {
		if Isprime(i) {
			sum += i
		}
	}
	fmt.Println(sum)

}

func Isprime(n int) bool {

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

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
