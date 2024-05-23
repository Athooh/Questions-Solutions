package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	var outPut1 []string
	var outPut2 []string

	num := Atoi(arg)
	if num == 0 {
		for _, c := range "ERROR: cannot convert to roman digit" {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
		return
	}
	if num >= 4000 {
		for _, c := range "ERROR: cannot convert to roman digit" {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
		return
	}
	for num > 0 {
		if num >= 1000 {
			outPut1 = append(outPut1, "M")
			outPut2 = append(outPut2, "M")
			num -= 1000
		} else if num >= 900 {
			outPut1 = append(outPut1, "CM")
			outPut2 = append(outPut2, "(M-C)")
			num -= 900
		} else if num >= 500 {
			outPut1 = append(outPut1, "D")
			outPut2 = append(outPut2, "D")
			num -= 500
		} else if num >= 400 {
			outPut1 = append(outPut1, "CD")
			outPut2 = append(outPut2, "(D-C)")
			num -= 400
		} else if num >= 100 {
			outPut1 = append(outPut1, "C")
			outPut2 = append(outPut2, "C")
			num -= 100
		} else if num >= 90 {
			outPut1 = append(outPut1, "XC")
			outPut2 = append(outPut2, "(C-X)")
			num -= 90
		} else if num >= 50 {
			outPut1 = append(outPut1, "L")
			outPut2 = append(outPut2, "L")
			num -= 50
		} else if num >= 40 {
			outPut1 = append(outPut1, "XL")
			outPut2 = append(outPut2, "(L-X)")
			num -= 40
		} else if num >= 10 {
			outPut1 = append(outPut1, "X")
			outPut2 = append(outPut2, "X")
			num -= 10
		} else if num >= 9 {
			outPut1 = append(outPut1, "IX")
			outPut2 = append(outPut2, "(X-I)")
			num -= 9
		} else if num >= 5 {
			outPut1 = append(outPut1, "V")
			outPut2 = append(outPut2, "V")
			num -= 5
		} else if num >= 4 {
			outPut1 = append(outPut1, "IV")
			outPut2 = append(outPut2, "(V-I)")
			num -= 4
		} else if num >= 1 {
			outPut1 = append(outPut1, "I")
			outPut2 = append(outPut2, "I")
			num -= 1
		}
	}
	// fmt.Println(outPut1)
	// fmt.Println(outPut2)

	for i, slice := range outPut2 {
		for _, c := range slice {
			z01.PrintRune(c)
		}
		if i != len(outPut2)-1 {
			z01.PrintRune('+')
		}

	}
	z01.PrintRune('\n')
	for _, slice := range outPut1 {
		for _, c := range slice {
			z01.PrintRune(c)
		}

	}
	z01.PrintRune('\n')
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
