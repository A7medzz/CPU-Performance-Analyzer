package main

import (
	"fmt"
	"math"
	"time"
)

func jumpSearchNaive(arr []int, target int) int {
	n := len(arr)
	step := int(math.Sqrt(float64(n)))
	prev := 0

	for arr[int(math.Min(float64(step), float64(n)))-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	for i := prev; i < int(math.Min(float64(step), float64(n))); i++ {
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
	index := jumpSearchNaive(arr, target)
	duration := time.Since(start)

	fmt.Printf("Target found at index: %d\n", index)
	fmt.Printf("Execution time: %v\n", duration)
}
