package main

import(
	"fmt"
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func bubblesort(a[] int, n int){
	for i := 0; i < n-1; i++ {
		for  j := 0; j < n-i-1; j++ {
			if a[j] > a [j+1] {
				swap(&a[j], &a[j+1])
			}
		}
	} 
}

func printarray(a[] int, size int){
	for i := 0; i < size; i++ {
		fmt.Println("%d", a[i])
	}
}

func main(){
	a := make([] int, 5)
	a[0], a[1], a[2], a[3], a[4] = 3, 4, 5, 2, 6
	n := len(a)
	bubblesort(a, n)
	fmt.Println("sorted array is:")
	printarray(a, n)

}



