package main

import (
	"fmt"
)

func isPalindrome(n int) bool {
	original := n
	reverse := 0

	for n > 0 {
		remainder := n % 10
		reverse = reverse*10 + remainder
		n /= 10
	}

	return original == reverse
}

func largestPalindromeProduct() int {
	largest := 0

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			if isPalindrome(product) && product > largest {
				largest = product
			}
		}
	}

	return largest
}

func main() {
	result := largestPalindromeProduct()
	fmt.Printf("The largest palindrome product of two 3-digit numbers is: %d\n", result)
}
