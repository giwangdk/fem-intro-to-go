package main

import "fmt"

func main() {
	sentences := "12333444"
	for _, value := range sentences {
		if int(value)%3 == 0 {
			fmt.Println(string(value))
		}
	}
}
