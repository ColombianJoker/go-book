package main

import (
	"fmt"
)

// suma calculates the sum of integers from 1 to n.
func suma(n uint64) uint64 {
	var S uint64 = 0
	for i := uint64(1); i <= n; i++ {
		S += i
	}
	return S
}

// main function to execute the summation multiple times.
func main() {
	fmt.Println("Go:")

	times := uint64(30)
	block := uint64(1_000_000_000) // Go supports underscores for readability in numeric literals

	// Create a slice to store the results of each suma call
	t := make([]uint64, times)

	// Execute the suma function 'times' many times
	for i := uint64(0); i < times; i++ {
		t[i] = suma(block)
	}

	// Calculate the sum of all results
	var totalSum uint64 = 0
	for _, val := range t {
		totalSum += val
	}

	fmt.Printf("%d×suma(%d) = %d\n", times, block, totalSum)
}
