package main

import "fmt"

// Optimized Interpolation Search
func interpolationOptimized(arr []int, x int) int {
	low, high := 0, len(arr)-1

	for low <= high && x >= arr[low] && x <= arr[high] {
		if arr[high] == arr[low] {
			if arr[low] == x {
				return low
			}
			break
		}

		pos := low + int(float64(high-low)*float64(x-arr[low])/float64(arr[high]-arr[low]))

		if pos < low || pos > high {
			break
		}

		if arr[pos] == x {
			return pos
		} else if arr[pos] < x {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	target := 14

	idx := interpolationOptimized(arr, target)

	if idx != -1 {
		fmt.Printf("Found %d at index %d using optimized version\n", target, idx)
	} else {
		fmt.Printf("%d not found using optimized version\n", target)
	}
}
