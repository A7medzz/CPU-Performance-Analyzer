package main

import (
	"fmt"
	"time"
)

func linearOptimized(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	target := 14

	start := time.Now()
	result := linearOptimized(arr, target)
	elapsed := time.Since(start)

	fmt.Printf("Optimized Linear Search Result: %d\n", result)
	fmt.Printf("Elapsed Time: %s\n", elapsed)
}
