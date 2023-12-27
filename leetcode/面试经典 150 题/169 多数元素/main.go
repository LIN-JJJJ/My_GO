package main

import "fmt"

func main() {
	//fmt.Println(majorityElement([]int{1, 1, 2}))
	//fmt.Println(majorityElement([]int{3, 2, 2, 3}))
	//fmt.Println(majorityElement([]int{1}))
	//fmt.Println(majorityElement([]int{1, 1, 1, 1, 1}))
	//fmt.Println(majorityElement([]int{1, 1, 1, 2, 2, 3}))
	//fmt.Println(majorityElement([]int{8, 8, 7, 7, 7}))
	fmt.Println(majorityElement([]int{12, 52, 12, 70, 12, 61, 12, 12, 50, 72, 82, 12, 11, 25, 28, 43, 40, 12, 12, 53, 12, 12, 98, 12, 12, 92, 81, 2, 12, 15, 40, 12, 12, 9, 12, 31, 12, 12, 12, 12, 77, 12, 12, 50, 12, 12, 12, 93, 41, 92, 12, 12, 12, 12, 12, 12, 12, 12, 12, 37, 48, 14, 12, 70, 82, 12, 80, 12, 12, 12, 12, 56, 30, 12, 8, 12, 50, 12, 20, 12, 21, 97, 12, 42, 12, 10, 12, 38, 73, 15, 9, 11, 79, 12, 12, 28, 51, 12, 15, 12}))

}
func majorityElement(nums []int) int {
	m := make(map[int]int) // k v 次数
	max := len(nums)
	vmax := 0
	vk := 0
	for i := 0; i < max; i++ {
		m[nums[i]] += 1
	}
	if len(m) == 1 {
		return nums[0]
	}
	for k, v := range m {
		if v > vmax {
			vmax = v
			vk = k
		}
	}
	return vk
}
