package main

import "fmt"

func main() {
	n := 10
	evenSum := 0

	arr := []int{}

	for i := range n {
		arr = append(arr, i)
	}

	for _, num := range arr {
		if num%2 == 0 {
			evenSum += num
		}
	}

	fmt.Printf("The sum of even nums is: %d \n", evenSum)
}
