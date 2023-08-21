package main

import (
	"fmt"
)

func main() {
	limit := 400 // Set your desired limit here
	sum := 0
	a := 1
	b := 2

	for b <= limit {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}

	fmt.Printf("The sum of even Fibonacci numbers up to %d is: %d\n", limit, sum)
}
