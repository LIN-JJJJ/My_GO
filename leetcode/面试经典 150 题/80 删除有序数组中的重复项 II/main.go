package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{3, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 1}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))

}
func removeDuplicates(nums []int) int {
	L := 0
	r := make(map[int]int)
	for _, num := range nums {
		f, _ := r[num]
		if f < 2 {
			nums[L] = num
			r[num] = f + 1
			L++
		}
	}
	return L
}
