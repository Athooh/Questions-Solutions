package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Valid options are: -abcdefghijklmnopqrstuvwxyz")
		return
	}
	args := os.Args
	var optionsInt uint32 = 0

	for _, arg := range args[1:] {
		if len(arg) > 0 && arg[0] == '-' {
			if len(arg) > 1 && contains(arg) {
				fmt.Println("Valid options are: -abcdefghijklmnopqrstuvwxyz")
				return
			}
			for _, ch := range arg[1:] {
				if ch >= 'a' && ch <= 'z' {
					bitposition := ch - 'a'
					optionsInt |= (1 << bitposition)
				} else {
					fmt.Println("Invalid Option")
					return
				}

			}
		} else {
			fmt.Println("Invalid Option")
			return
		}
	}

	bitString := fmt.Sprintf("%032b", optionsInt)

	for i := 0; i < 32; i += 8 {
		if (i + 8) < 32 {
			fmt.Print(bitString[i:i+8] + " ")
		} else {
			fmt.Print(bitString[i : i+8])
		}
	}
	fmt.Println()
}
func contains(s string) bool {
	return s[:2] == "-h"
}
