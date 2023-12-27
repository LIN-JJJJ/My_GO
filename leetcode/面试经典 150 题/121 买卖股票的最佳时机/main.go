package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	p := 0    // 累计利润
	maxp := 0 // 最大利润
	for i := 1; i < len(prices); i++ {
		pp := prices[i] - prices[i-1] // 今日对比昨日的利润
		p += pp
		if p <= 0 {
			p = 0
			continue
		} else {
			if maxp < p {
				maxp = p
			}
		}
	}
	return maxp
}
