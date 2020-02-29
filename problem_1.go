package main

import (
	"math"
	"fmt"
)

func main() {
	var counter float64
	var accum float64
	for counter < 1000 {
		if(math.Mod(counter,3) == 0 || math.Mod(counter,5) == 0){
			accum += counter
		}
		counter += 1
	}
	fmt.Printf("%g", accum)
}
