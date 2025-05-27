package main

import (
	"fmt"
)

func recursiveLinearSearch(arr []int, target int, index int) int {
	// Early return for empty slice (small optimization)
	if len(arr) == 0 || index >= len(arr) {
		return -1
	}

	// Reduce function calls by checking once and returning immediately
	if arr[index] == target {
		return index
	}
	return recursiveLinearSearch(arr, target, index+1)
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	target := 14

	result := recursiveLinearSearch(arr, target, 0)
	if result == -1 {
		fmt.Println("Element not found")
	} else {
		fmt.Printf("Element found at index: %d\n", result)
	}
}
