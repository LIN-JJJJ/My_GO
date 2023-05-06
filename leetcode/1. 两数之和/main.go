package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))

}
func twoSumV1(nums []int, target int) []int {
	var res []int
	maxLen := len(nums)
	for i := 0; i < maxLen-1; i++ {
		for i1 := i + 1; i1 < maxLen; i1++ {
			if target == nums[i]+nums[i1] {
				return append(res, i, i1)
			}
		}
	}
	return res
}
func twoSum(nums []int, target int) []int {
	var res []int
	maxLen := len(nums)
	numsMap := make(map[int]int, maxLen/2)
	for i := 0; i < maxLen; i++ {
		yi, ok := numsMap[target-nums[i]]
		if ok {
			return append(res, yi, i)
		}
		numsMap[nums[i]] = i
	}
	return res
}
