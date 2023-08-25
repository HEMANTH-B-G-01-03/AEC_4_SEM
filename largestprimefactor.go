package main

import (
	"fmt"
)

func largestPrimeFactor(n int) int {
	largest := 0

	// Find and remove all factors of 2
	for n%2 == 0 {
		largest = 2
		n /= 2
	}

	// Start with the next potential prime factor
	factor := 3
	for factor*factor <= n {
		if n%factor == 0 {
			largest = factor
			n /= factor
		} else {
			// Move to the next odd number
			factor += 2
		}
	}

	// If n is still greater than 1, it's the largest prime factor
	if n > 1 {
		largest = n
	}

	return largest
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	largest := largestPrimeFactor(num)
	fmt.Printf("The largest prime factor of %d is: %d\n", num, largest)
}
