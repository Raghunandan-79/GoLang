package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var nums2 = make([]int, 2, 5)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 3, 4, 5)
	fmt.Println(nums2)
	fmt.Println(cap(nums2))
	nums2[0] = 1
	fmt.Println(nums2)

	var nums3 = []int{}
	nums3 = append(nums3, 1, 2)
	var nums4 = make([]int, len(nums3))
	copy(nums4, nums3)
	fmt.Println(nums4)

	// slice operator
	var nums5 = []int{1, 2, 3}
	fmt.Println(nums5[0:2])
	fmt.Println(nums5[:2])
	fmt.Println(nums5[0:])
	fmt.Println(nums5[0:2])

	// slice package
	var nums6 = []int{1, 2}
	var nums7 = []int{1, 2}

	fmt.Println(slices.Equal(nums6, nums7))

	// 2d slice
	var two_d_slice = [][]int{{1, 2}, {3, 4}}
	fmt.Println(two_d_slice)
	
}