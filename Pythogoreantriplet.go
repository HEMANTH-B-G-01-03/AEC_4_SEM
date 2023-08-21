package main

import "fmt"

func findSpecialPythagoreanTriplet(sum int) (int, int, int) {
	for a := 1; a <= sum/3; a++ {
		for b := a + 1; b <= sum/2; b++ {
			c := sum - a - b
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return -1, -1, -1 // No special Pythagorean triplet found
}

func main() {
	targetSum := 1000 // Replace with your desired sum
	a, b, c := findSpecialPythagoreanTriplet(targetSum)

	if a != -1 {
		product := a * b * c
		fmt.Printf("Special Pythagorean triplet with a + b + c = %d found: a = %d, b = %d, c = %d\n", targetSum, a, b, c)
		fmt.Printf("The product abc is: %d\n", product)
	} else {
		fmt.Printf("No special Pythagorean triplet with a + b + c = %d found.\n", targetSum)
	}
}
