package main

import (
	"fmt"
	"math"
	"time"
)

func jumpSearchOptimized(arr []int, target int) int {
	n := len(arr)
	step := int(math.Sqrt(float64(n)))
	end := step

	for end < n && arr[end-1] < target {
		end += step
	}
	start := end - step
	if end > n {
		end = n
	}
	for i := start; i < end; i++ {
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
	index := jumpSearchOptimized(arr, target)
	duration := time.Since(start)

	fmt.Printf("Target found at index: %d\n", index)
	fmt.Printf("Execution time: %v\n", duration)
}
