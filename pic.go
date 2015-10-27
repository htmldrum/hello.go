package main

import(
	"golang.org/x/tour/pic"
//	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	bytes := make([][]uint8, dy)

	for n := range bytes {
		bytes[n] = make([]uint8, dx)
		for o := range bytes[n] {
			bytes[n][o] = uint8((n * o)/2)
		}
	}

	//fmt.Printf("dx: %d\n", dx)
	//fmt.Printf("dy: %d\n", dy)
	//fmt.Printf("bytes: %s len=%d cap=%d %v\n", bytes, len(bytes), cap(bytes), bytes)

	return bytes
}

func main(){
	pic.Show(Pic)
}
