package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindPrevPrime(1))
}

func FindPrevPrime(nb int) int {
	if nb <= 1 {
		return 0
	}
	for {
		if IsPrime(nb) {
			return nb
		}
		nb--
	}

}
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}
