package main

import(
	"fmt"
)


func permute(inp []string, l int, r int){
	if l == r {
		fmt.Println("%s\n", inp)
	}else {
		for i := 1; i <=r; i++ {
			inp[l], inp[i] = inp[i], inp[l]
			permute(inp, l+1, r)
			inp[1], inp[i] = inp[i], inp[1] 
		}
	}
}

func main(){
	inp := make([] string, 3) 
	inp[0], inp[1], inp[2] = "a","b","c"
	n := len(inp)
	//fmt.Println("string is:")
	permute(inp, 0, n-1)
}




