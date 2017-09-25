package main

import(
	"fmt"
)

func binary_search(a []int, l int, r int, x int)bool{
	mid := l+(r-1)/2
	if a[mid] == x {
		return	true
	}
	if a[mid] > x {
		return binary_search(a, l, mid-1, x)
	}
	return binary_search(a, mid+1, r, x)
}

func main(){
	a := []int{2, 3, 4, 5, 6, 7, 8}
	n := len(a)
	if binary_search(a, 0, n-1, 7) {
		fmt.Println("Present")
	}else{
		fmt.Println("Not Present")
	}
}

