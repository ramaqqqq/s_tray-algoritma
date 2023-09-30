package main

import (
	"fmt"
	"sort"
)

// soal algoritma 1
func maxProductOfThree(nums []int) int {
	if len(nums) < 3 {
		panic("Setidaknya tiga angka")
	}
	sort.Ints(nums)
	max1 := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	max2 := nums[0] * nums[1] * nums[len(nums)-1]
	if max1 > max2 {
		return max1
	}
	return max2
}

// soal algoritma 2
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		panic("Setidaknya dua harga")
	}
	maxProfit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return maxProfit
}

func main() {
	//soal algoritma 1
	n1 := []int{1, 2, 3}
	fmt.Println(maxProductOfThree(n1))
	n2 := []int{-10, -10, 1, 3, 2}
	fmt.Println(maxProductOfThree(n2))

	//soal algoritma 2
	p1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(p1))
	p2 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(p2))
}
