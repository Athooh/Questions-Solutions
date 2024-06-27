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

// Seth Solution
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	er := "ERROR: cannot convert to roman digit"

	number, err := Atoi(os.Args[1])

	if !err {
		for _, v := range er {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		return
	}

	if number == 0 || number >= 4000 {
		for _, v := range er {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		return
	}

	var romanMethod []string
	var romanNumbers []string

	for number > 0 {
		if number >= 1000 {
			romanNumbers, romanMethod = append(romanNumbers, "M"), append(romanMethod, "M")
			number -= 1000
		} else if number >= 900 {
			romanNumbers, romanMethod = append(romanNumbers, "CM"), append(romanMethod, "(M-C)")
			number -= 900
		} else if number >= 500 {
			romanNumbers, romanMethod = append(romanNumbers, "D"), append(romanMethod, "D")
			number -= 500
		} else if number >= 400 {
			romanNumbers, romanMethod = append(romanNumbers, "CD"), append(romanMethod, "(D-C)")
			number -= 400
		} else if number >= 100 {
			romanNumbers, romanMethod = append(romanNumbers, "C"), append(romanMethod, "C")
			number -= 100
		} else if number >= 90 {
			romanNumbers, romanMethod = append(romanNumbers, "XC"), append(romanMethod, "(C-X)")
			number -= 90
		} else if number >= 50 {
			romanNumbers, romanMethod = append(romanNumbers, "L"), append(romanMethod, "L")
			number -= 50
		} else if number >= 10 {
			romanNumbers, romanMethod = append(romanNumbers, "X"), append(romanMethod, "X")
			number -= 10
		} else if number >= 9 {
			romanNumbers, romanMethod = append(romanNumbers, "IX"), append(romanMethod, "(X-I)")
			number -= 9
		} else if number >= 5 {
			romanNumbers, romanMethod = append(romanNumbers, "V"), append(romanMethod, "V")
			number -= 5
		} else if number >= 4 {
			romanNumbers, romanMethod = append(romanNumbers, "IV"), append(romanMethod, "(V-I)")
		} else if number >= 1 {
			romanNumbers, romanMethod = append(romanNumbers, "I"), append(romanMethod, "I")
			number -= 1
		}
	}
	var method string
	for i, v := range romanMethod {
		if i < len(romanMethod)-1 {
			method += v + "+"
		} else {
			method += v
		}
	}
	for _, v := range method {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')

	for _, v := range romanNumbers {
		for _, n := range v {
			z01.PrintRune(rune(n))
		}
	}
	z01.PrintRune('\n')
}

func Atoi(s string) (int, bool) {
	var number int
	sign := 1

	for idx, n := range s {
		if idx == 0 && n == '-' {
			sign = -1

		} else if idx == 0 && n == '+' {
			sign = 1

		} else if n >= '0' && n <= '9' {
			number = number*10 + int(n-'0')
		} else {
			return 0, false
		}
	}
	return number * sign, true
}

//Stella's soln
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		PrintError()
		return
	}

	arg := os.Args[1]
	num := Atoi(arg)
	if num <= 0 || num >= 4000 {
		PrintError()
		return
	}

	roman, calculation := ToRoman(num)
	PrintStr(calculation)
	PrintStr(roman)
}

func ToRoman(num int) (string, string) {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman, calculation := "", ""

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman += sym[i]
			if calculation != "" {
				calculation += "+"
			}
			calculation += sym[i]
		}
	}
	return roman, calculation
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func PrintError() {
	msg := "ERROR: cannot convert to roman digit"
	PrintStr(msg)
}

func Atoi(s string) int {
	q := 0
	sign := 1

	for idx, char := range s {
		if char == '-' && idx == 0 {
			sign = -1
		} else if char == '+' && idx == 0 {
			sign = 1
		} else if char >= '0' && char <= '9' {
			q = q*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return q * sign
}
