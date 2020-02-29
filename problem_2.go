package main

import (
        "math"
        "fmt"
)

func main() {
	fmt.Printf("Result: %f", fibb(4000000))
}

func fibb(x float64) float64{
	var curr float64
	var p1 float64
	var p2 float64
	p1 = 1
	p2 = 1
	var accum float64
	for p2 < x{
		curr = p1 + p2
		if(math.Mod(curr,2) == 0){
			accum += curr
		}
		p1 = p2
		p2 = curr
	}
	return accum
}
