package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{1}, 1))
	fmt.Println(removeElement([]int{1, 1, 1, 1, 1}, 1))

}

func removeElement(nums []int, val int) int {
	L := 0
	for _, num := range nums {
		if num != val {
			nums[L] = num
			L++
		}
	}
	return L
}
