package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	num1, num2 := Atoi(os.Args[1]), Atoi(os.Args[2])

	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	fmt.Println(num1)

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
