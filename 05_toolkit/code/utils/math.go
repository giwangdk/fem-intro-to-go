package utils

import "fmt"

func printNum(num int) {
	fmt.Println(num)
}

func Add(nums ...int) int {
	total := 0
	for _, v := range nums {
		printNum(v)
		total += v
	}

	return total
}
