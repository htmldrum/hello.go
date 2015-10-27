package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32}

func main(){
	for n, v := range pow {
		fmt.Printf("2**%d = %d\n", n, v)
	}
}

/*
	2**0 = 1
	2**1 = 2
	2**2 = 4
	2**3 = 8
	2**4 = 16
	2**5 = 32
*/
