package main

import (
	"fmt"
)

func search(arr[] int, n int, x int) int {
	for i := 0; i < n; i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}

func main(){
	arr := make([]int, 5)
	arr[0], arr[1], arr[2], arr[3], arr[4] = 1, 10, 20, 34, 45
	x := 34
	n := len(arr)
	fmt.Printf("%d is present at index %d\n", x, search(arr, n, x))
}

