package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	ptr1, ptr2 := 0, 0

	for ptr2 < len(str2) {

		if ptr1 < len(str1) && str1[ptr1] == str2[ptr2] {
			ptr1++
		}
		ptr2++
	}

	if ptr1 == len(str1) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

}
