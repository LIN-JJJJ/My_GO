package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 9, 1, 1}))
}
func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}
	return res
}
