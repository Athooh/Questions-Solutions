// findprevprime
// Instructions

// Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

// If there are no primes inferior to the int passed as parameter the function should return 0.
// Expected function

// func FindPrevPrime(nb int) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import "fmt"

// func main() {
// 	fmt.Println(FindPrevPrime(5))
// 	fmt.Println(FindPrevPrime(4))
// }

// And its output :

// $ go run .
// 5
// 3
// $

package main

import "fmt"

func FindPrevPrime(nb int) int {
	var prevPrime int
	for i:= 2; i <= nb; i++ {
		if IsPrime(i) {
			prevPrime = i
		}
	}
	return prevPrime
}

func IsPrime(nb int) bool {
	if nb < 0 || nb == 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(FindPrevPrime(5)) // Output: 5
	fmt.Println(FindPrevPrime(4)) // Output: 3
	fmt.Println(FindPrevPrime(-2))
	fmt.Println(FindPrevPrime(3))
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(0))
	fmt.Println(FindPrevPrime(11))
	fmt.Println(FindPrevPrime(29))
}

