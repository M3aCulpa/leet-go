package main

import "fmt"

// first attempt
func sortColors(nums []int) []int {

	// selection sort
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

// second attempt
// Could you come up with a one-pass algorithm using only constant extra space?
// func sortColorsBetter(nums []int)  []int {
// }

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(sortColors(nums))
}
