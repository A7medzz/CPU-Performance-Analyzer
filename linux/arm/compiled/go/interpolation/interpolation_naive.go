package main

import "fmt"

// Naive Interpolation Search
func interpolationNaive(arr []int, x int) int {
	low, high := 0, len(arr)-1

	for low <= high && x >= arr[low] && x <= arr[high] {
		if arr[high] == arr[low] {
			if arr[low] == x {
				return low
			}
			break
		}

		pos := low + ((high - low) * (x - arr[low])) / (arr[high] - arr[low])

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

	idx := interpolationNaive(arr, target)

	if idx != -1 {
		fmt.Printf("Found %d at index %d using naive version\n", target, idx)
	} else {
		fmt.Printf("%d not found using naive version\n", target)
	}
}
