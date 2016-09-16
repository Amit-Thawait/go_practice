package main

import "fmt"

func main() {
	// Array
	nums := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("Type of nums : %T\n", nums)
	fmt.Println("Original Array :", nums)
	reverse(nums[:])
	fmt.Println("Reversed Array :", nums)

	// Slice
	a := []int{0, 1, 2, 3, 4, 5}
	reverse(a[:2])
	reverse(a[2:])
	reverse(a)
	fmt.Println("Rotated Array :", a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
