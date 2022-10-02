package main

import "fmt"

func average(nums ...float32) float32 {
	var total float32 = 0.0
	for _, value := range nums {
		total += value
	}
	return total / float32(len(nums))

}

func main() {
	fmt.Println(average(1.2, 1.3, 2.4, 5.5))
}
