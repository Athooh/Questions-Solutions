// chunk
// Instructions
// Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.

// If the size is 0 it should print a newline ('\n').
// Expected function
// func Chunk(slice []int, size int) {

// }
// Usage
// Here is a possible program to test your function :

// package main

// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }
// And its output :

// $ go run .
// []

// [[0 1 2] [3 4 5] [6 7]]
// [[0 1 2 3 4] [5 6 7]]
// [[0 1 2 3] [4 5 6 7]]
// $

package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	m := len(slice) / size
	rem := len(slice) % size
	if rem >= 1 {
		m++
	}

	maped := make([][]int, m)

	n := 0
	count := 0

	for i := 0; i < len(slice); i++ {
		if count == size {
			n++
			count = 0
			i--
		} else {
			maped[n] = append(maped[n], slice[i])
			count++
		}
	}
	println(size, ": ", maped)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

// seth solution

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
		return
	}

	var newSlice [][]int

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		newSlice = append(newSlice, slice[i:end])
	}
	fmt.Println(newSlice)
}
//Rays solution
// func Chunk(slice []int, size int) {
// 	count := 0
// 	var newSlice []int
// 	var outputSlice [][]int
// 	if size == 0 {
// 		fmt.Println()
// 		return
// 	}
// 	if len(slice) == 0 {
// 		fmt.Println(slice)
// 		return

// 	}

// 	for _, n := range slice {
// 		newSlice = append(newSlice, n)
// 		count++
// 		if count == size {
// 			outputSlice = append(outputSlice, newSlice)
// 			newSlice = []int{}
// 			count = 0
// 		}

// 	}
// 	if len(newSlice) > 0 {
// 		outputSlice = append(outputSlice, newSlice)
// 	}

// 	fmt.Println(outputSlice)

// }

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
