package main

import "fmt"

func main() {
	//fmt.Println(isWinner([]int{4, 10, 7, 9}, []int{6, 5, 2, 3}))
	//fmt.Println(isWinner([]int{5, 6, 1, 10}, []int{5, 1, 10, 5}))
	//fmt.Println(isWinner([]int{0, 4, 7, 2, 0}, []int{2, 3, 3, 0, 1}))
	fmt.Println(isWinner([]int{3, 6, 10, 8}, []int{9, 9, 9, 9}))

}

//n == player1.length == player2.length
//1 <= n <= 1000
//0 <= player1[i], player2[i] <= 10

func isWinner(player1 []int, player2 []int) int {
	sum1, sum2 := 0, 0
	n := len(player1)
	for i := 0; i < n; i++ {
		if (i > 0 && player1[i-1] == 10) || (i > 1 && player1[i-2] == 10) {
			sum1 += player1[i] * 2
		} else {
			sum1 += player1[i]
		}
		if (i > 0 && player2[i-1] == 10) || (i > 1 && player2[i-2] == 10) {
			sum2 += player2[i] * 2
		} else {
			sum2 += player2[i]
		}
	}
	if sum1 > sum2 {
		return 1
	} else if sum2 > sum1 {
		return 2
	}
	return 0
}
