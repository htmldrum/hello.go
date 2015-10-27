// http://play.golang.org/p/7gQJt_UohM
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var repeat_count int = 10
	var z float64 = 1
	
	for i:=0; i < repeat_count; i++ {
		z = z - float64(math.Pow(z,2) - x)/2*z
	}
	
	return z
}

func ConvergantSqrt(x float64) float64 {
	var (
		z float64 = 1
		last_z float64
		count int
	)
	
	for(!AlmostTheSame(z,last_z)){
		last_z = z
		z = z - float64(math.Pow(z,2) - x)/2*z
		count += 1
	}
	
	fmt.Printf("Loops: %v\n", count);
	
	return z
}

func AlmostTheSame(val float64, estimate float64) bool {
	const CONFIDENCE_INTERVAL float64 = 0.005
	return ((val < estimate * (1 + CONFIDENCE_INTERVAL)) && (val > estimate * (1 - CONFIDENCE_INTERVAL)));
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(ConvergantSqrt(2))
}
