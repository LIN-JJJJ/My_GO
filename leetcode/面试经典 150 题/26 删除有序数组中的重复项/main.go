package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{3, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 1}))

}

func removeDuplicates(nums []int) int {
	L := 0
	r := make(map[int]struct{})
	for _, num := range nums {
		_, ok := r[num]
		if !ok {
			nums[L] = num
			r[num] = struct{}{}
			L++
		}
	}
	return L
}
