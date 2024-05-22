package main

import (
	"fmt"
	"os"
)

func pigLatin(s string) (result string) {
// handle wordswith no vowels
	check := false
	for _, char := range s {
		if checkVowel(char) {
			check = true

		}
		if check {
			break
		}

	}
	if !check {
		result =("No vowels")
		return result
	}
	//handle words that start with vowels
	if checkVowel(rune(s[0])) {
		result = s + "ay"
		return result
	}
	//handle words that have vowels in the middle
	for i, char:=range s{
		if i!=0 && checkVowel(char){
			v:=s[i:]
			s=s[:i]
			result=v+s+"ay"
		}
	}

	return result
}

//function to check for vowels
func checkVowel(s rune) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u'
}

func main() {
	if len(os.Args)!=2{
		return
	}
	args := os.Args[1]
	fmt.Println(pigLatin(args))
}
