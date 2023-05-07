package main

import "fmt"

func main() {

	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
	fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))
}

func numPairsDivisibleBy60(times []int) int {
	cnt := make([]uint16, 60)
	for _, t := range times {
		cnt[t%60]++
	}
	var res int
	res += int(1+(cnt[0]-1)) * int(cnt[0]-1) / 2   //正好是60秒的 等差公式
	res += int(1+(cnt[30]-1)) * int(cnt[30]-1) / 2 //正好是30秒的 等差公式

	for i := 1; i < 30; i++ {
		res += int(cnt[i]) * int(cnt[60-i])
	}
	return int(res)
}
