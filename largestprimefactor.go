package main

import (
	"fmt"
	"math"
)

func largestPrimeFactor(n int) int {
	var largest int

	// Remove all the even factors first
	for n%2 == 0 {
		largest = 2
		n /= 2
	}

	// Now, n must be an odd number
	for i := int(3); i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			largest = i
			n /= i
		}
	}

	// If n is still greater than 2, it's the largest prime factor
	if n > 2 {
		largest = n
	}

	return largest
}

func main() {
	var num int
	fmt.Println("enter the  number ")
	fmt.Scanln(&num)
	largest := largestPrimeFactor(num)
	fmt.Printf("The largest prime factor of %d is: %d\n", num, largest)
}
