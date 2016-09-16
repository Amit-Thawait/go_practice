package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("Type of nums : %T\n", nums)
	fmt.Println("Original Array :", nums)
	reverse(nums)
	fmt.Println("Reversed Array :", nums)
	reverse(nums)
	reverse(nums[:2])
	reverse(nums[2:])
	reverse(nums)
	fmt.Println("Rotated Array :", nums)
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
