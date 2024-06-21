package main

import (
	"fmt"
)

func main() {
	var b byte = 0b00100110

	fmt.Printf("%08b", ReverseBits(b))

}
func ReverseBits(oct byte) byte {
	var reversed byte = 0

	for i := 0; i < 8; i++ {
		// Extract the i-th bit from the original byte
		bit := (oct >> i) & 1
		// Place the bit in the reversed position
		reversed = reversed | (bit << (7 - i))
	}

	return reversed
}

// DURING CHECKPOINT CHANGE PACKAGE TO PISCINE AFTER TESTING YOU FUNCTION WITH THE MAIN 
// FUNCTION. DELETE THE MAIN FUNCTION BEFORE SUBMITTING YOUR FUNCTION. THIS APPLIES TO PRINTBITS AND SWAPBITS 
// AS WELL.