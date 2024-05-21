package main

import (
	"os"
	"strconv"
)

//to add Atoi manual
func doop(s []string) {
	if len(s[0]) > 12 {
		return
	}
	num1, _ := strconv.Atoi(s[0])
	operator := s[1]
	num2, _ := strconv.Atoi(s[2])
	result := 0

	if num1 == 0 || num2 == 0 && operator == "/" {
		os.Stdout.Write([]byte("No division by 0\n\n"))
		return
	}
	if num1 == 0 || num2 == 0 && operator == "%" {
		os.Stdout.Write([]byte("No modulo by 0\n\n"))
		return
	}
	switch operator {
	case "+":
		result = num1 + num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))
	case "-":
		result = num1 - num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	case "*":
		result = num1 * num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	case "/":
		result = num1 / num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	case "%":
		result = num1 % num2

		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	default:
		return
	}
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	doop(os.Args[1:])
}

func Itoa(n int) string {
	//declare variables
	r := []byte{}
	temps := []byte{}
	right := 0

	//handle -ve numbers
	if n < 0 {
		r = append(r, '-')
		n *= -1
	}

	//convert each digit to a string by adding  "0"
	for {
		right = n % 10
		n = n / 10
		temps = append(temps, byte('0'+right))
		if n == 0 {
			break
		}

	}
	for i := len(temps) - 1; i >= 0; i-- {
		r = append(r, temps[i])
	}
	return string(r)
 }

//rays solution
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	if len(os.Args[1:]) != 3 {
// 		return
// 	}

// 	arg1 := os.Args[1]
// 	operator := os.Args[2]
// 	arg2 := os.Args[3]
// 	var num1 int
// 	var num2 int
// 	var ans int

// 	for i, v := range arg1 {
// 		if i == 0 && v == '-' {
// 			continue
// 		}
// 		num := int(v - 48)
// 		if num > 9 {
// 			return
// 		}
// 		num1 = (num1 * 10) + num

// 	}

// 	if arg1[0] == '-' {
// 		num1 = -1 * num1
// 	}

// 	for i, v := range arg2 {

// 		if i == 0 && v == '-' {
// 			continue
// 		}

// 		num := int(v - 48)

// 		if num > 9 {
// 			return
// 		}
// 		num2 = (num2 * 10) + num

// 	}
// 	if arg2[0] == '-' {
// 		num2 = -1 * num2
// 	}
// 	//fmt.Println(num1)
// 	//fmt.Println(num2)

// 	if (num1 >= 9223372036854775807 || num1 <= -9223372036854775807) || (num2 >= 9223372036854775807 || num2 <= -9223372036854775807) {
// 		return
// 	}

// 	if (num1 == 0 || num2 == 0) && operator == "/" {
// 		fmt.Println("No division by 0")
// 		return
// 	} else if (num1 == 0 || num2 == 0) && operator == "%" {
// 		fmt.Println("No modulo by 0")
// 		return
// 	}
// 	switch operator {
// 	case "+":
// 		ans = num1 + num2
// 	case "-":
// 		ans = num1 - num2
// 	case "/":
// 		ans = num1 / num2
// 	case "%":
// 		ans = num1 % num2
// 	case "*":
// 		ans = num1 * num2
// 	default:
// 		return
// 	}
// 	fmt.Println(ans)

// }

