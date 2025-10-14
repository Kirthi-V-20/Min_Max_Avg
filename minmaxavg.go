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

func main() {
	fmt.Println(largest([]int{5, 7, 3}))
}
