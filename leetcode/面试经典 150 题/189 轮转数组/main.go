package main

import "fmt"

func main() {
	fmt.Println(rotate1([]int{1, 2, 3, 4, 5, 6, 7}, 3))
}
func rotate1(nums []int, k int) []int {
	rotate(nums, k)
	return nums
}

func rotate(nums []int, k int) {
	// (下标+k)%len(nums) = 现下标
	// 偏移量 = k % len(nums)
	k = k % len(nums)
	l := len(nums)
	nl := append(nums[l-k:], nums[:l-k]...)
	copy(nums, nl)
}
