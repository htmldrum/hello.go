package main

import(
	"fmt"
)

func main(){

	game := [][]string{ // C-style array of arrays
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	game[0][0] = "X"
	game[2][2] = "O"

	fmt.Printf("game: %v", game);

}
