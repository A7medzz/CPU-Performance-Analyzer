package main

import (
	"fmt"
	"time"
)

func binarySearchNaive(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2 // no overflow protection
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	target := 14

	start := time.Now()
	index := binarySearchNaive(arr, target)
	duration := time.Since(start)

	fmt.Printf("Target found at index: %d\n", index)
	fmt.Printf("Execution time: %v\n", duration)
}
