package main

import "fmt"

func main() {
	fmt.Println(mergei([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
}
func mergei(nums1 []int, m int, nums2 []int, n int) []int {
	merge(nums1, m, nums2, n)
	return nums1
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2) // 补全
	//排序
	for i := 0; i < len(nums1)-1; i++ {
		for ii := 0; ii < len(nums1)-i-1; ii++ {
			if nums1[ii] > nums1[ii+1] {
				nums1[ii], nums1[ii+1] = nums1[ii+1], nums1[ii]
			}
		}
	}
}
