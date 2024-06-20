package main

import (
	"fmt"
)

func main() {

	fmt.Printf("Base16: %s\n", ItoaBase(95, 16))
	fmt.Printf("Base2: %s\n", ItoaBase(95, 2))
	fmt.Printf("Base4: %s\n", ItoaBase(95, 4))
	fmt.Printf("Base8: %s\n", ItoaBase(95, 8))
	fmt.Printf("Base3: %s\n", ItoaBase(95, 3))

}
func ItoaBase(value, base int) string {
	var ans string

	if base == 16 {
		hexadecimal := "0123456789abcdef"

		for value > 0 {
			indx := value % 16
			ans = string(hexadecimal[indx]) + ans
			value /= 16
		}
		return ans
	}

	for value > 0 {
		ansInt := value % base
		fmt.Println(ansInt)
		ans = Itoa(ansInt) + ans
		value /= base
	}

	return ans
}

func Itoa(n int) string {
	var res string
	if n == 0 {
		res += "0"
	}
	for n > 0 {
		rem := n % 10
		res = string(rune(rem+48)) + res
		n /= 10
	}
	return res
}
