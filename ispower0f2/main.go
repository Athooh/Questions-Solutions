package main

import (
	"fmt"
	"os"
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
