package main

import "fmt"

func rumoveDuplicate(n []int) []int {
	numbMap := make(map[int]bool)
	numbSlice := []int{}

	for _, v := range n {
		if !numbMap[v] {
			numbMap[v] = true
			numbSlice = append(numbSlice, v)
		}
	}
	return numbSlice
}

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 7, 4, 20, 100, 1234, 3, 9, 15, 8, 3, 7, 9, 10, 11, 12, 13, 18, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	rm := rumoveDuplicate(n)
	fmt.Println(rm)
}
