package main

import (
	"fmt"
	"time"
)

func linearNaive(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	target := 14

	start := time.Now()
	result := linearNaive(arr, target)
	elapsed := time.Since(start)

	fmt.Printf("Naive Linear Search Result: %d\n", result)
	fmt.Printf("Elapsed Time: %s\n", elapsed)
}
