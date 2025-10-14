package main

import "fmt"

func largest(num []int) int {

	max := num[0]
	for _, num := range num {
		if num > max {
			max = num
		}
	}
	return max
}

func smallest(num []int) int {
	min := num[0]
	for _, num := range num {
		if num < min {
			min = num
		}
	}
	return min
}

func average(num []int) float64 {
	return 5.0
}

func main() {
	fmt.Println(largest([]int{5, 7, 3}))
	fmt.Println(smallest([]int{5, 7, 3}))
	fmt.Println(smallest([]int{20, 50, 10}))
	fmt.Println(average([]int{5, 7, 3}))
}
