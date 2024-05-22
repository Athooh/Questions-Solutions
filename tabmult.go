// tabmult
// Instructions
// Write a program that displays a number's multiplication table.

// The parameter will always be a strictly positive number that fits in an int. Said number multiplied by 9 will also fit in an int.
// Usage
// $ go run . 9
// 1 x 9 = 9
// 2 x 9 = 18
// 3 x 9 = 27
// 4 x 9 = 36
// 5 x 9 = 45
// 6 x 9 = 54
// 7 x 9 = 63
// 8 x 9 = 72
// 9 x 9 = 81
// $ go run . 19
// 1 x 19 = 19
// 2 x 19 = 38
// 3 x 19 = 57
// 4 x 19 = 76
// 5 x 19 = 95
// 6 x 19 = 114
// 7 x 19 = 133
// 8 x 19 = 152
// 9 x 19 = 171
// $ go run .

// $

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}

// 	args := os.Args[1]
// 	numb, err := strconv.Atoi(args)
// 	if err != nil || numb <= 0 {
// 		fmt.Println("Error! Please provide a valid number.")
// 		return
// 	}

// 	for i := 1; i <= 9; i++ {
// 		fmt.Printf("%d x %d = %d\n", i, numb, i*numb)
// 	}
// }

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	stInt := os.Args[1]
	var number int
	var outPut []string
	for i, v := range stInt {
		if i == 0 && v == '-' {
			return
		}
		n := int(v - 48)
		number = ((number * 10) + n)
	}

	for i := 1; i <= 9; i++ {
		newStr := string(i + 48)
		res := (i * number)
		resStr := ""

		for res > 0 {
			digit := res % 10
			resStr = string(digit+'0') + resStr
			res /= 10
		}

		outPut = append(outPut, newStr+" "+"x"+" "+stInt+" "+"="+" "+resStr)

	}
	for i := 0; i < len(outPut); i++ {
		for _, v := range outPut[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}

}
