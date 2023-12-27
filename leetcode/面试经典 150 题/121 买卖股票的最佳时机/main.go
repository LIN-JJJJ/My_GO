package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPrice := prices[0]
	maxPrice := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > maxPrice {
			maxPrice = prices[i] - minPrice
		}
	}
	return maxPrice
}
