package main

import "fmt"

func main() {
	fmt.Println(hardestWorker(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}))
	fmt.Println(hardestWorker(26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}}))
	fmt.Println(hardestWorker(2, [][]int{{0, 10}, {1, 20}}))
}
func hardestWorker(n int, logs [][]int) int {
	startDay := 0 //记录起始时间
	id := 500     //员工ID
	MaxDat := 0   //最大耗时
	for _, log := range logs {
		expend := log[1] - startDay //计算耗时时间
		startDay = log[1]
		if expend > MaxDat {
			MaxDat = expend
			id = log[0]
		} else if expend == MaxDat {
			if log[0] < id {
				id = log[0]
			}
		}
	}
	return id
}
