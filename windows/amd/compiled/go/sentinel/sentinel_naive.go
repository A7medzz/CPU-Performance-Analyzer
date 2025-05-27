package main

import "fmt"

func sentinel(arr []int, target int) int {
    n := len(arr)
    last := arr[n-1]
    arr[n-1] = target
    i := 0
    for arr[i] != target {
        i++
    }
    arr[n-1] = last
    if i < n-1 || arr[n-1] == target {
        return i
    }
    return -1
}

func main() {
    arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	target := 14
    
    idx := sentinel(arr, target)
    fmt.Println("Index:", idx)
}
