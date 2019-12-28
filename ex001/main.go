package main

import (
	"fmt"
	"strconv"
)

func main() {

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range numbers {
		if i%2 == 0 {
			fmt.Println(strconv.Itoa(numbers[i]) + " is even")
		} else {
			fmt.Println(strconv.Itoa(numbers[i]) + " is odd")
		}
	}
}
