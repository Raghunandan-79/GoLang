package main

import "fmt"

func main() {
	var nums [4]int

	fmt.Println(len(nums))

	fmt.Println(nums[0])

	nums1  := [3]int{1, 2, 3}
	fmt.Println(nums1)

	// 2d arrays
	nums3 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums3)
}
