package main

import (
        "time"
        //"math"
        "fmt"
		//"strconv"
)

func main() {
        start := time.Now()
        fmt.Printf("Result: %f", find_triplet(1000))
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

func find_triplet(limit int) [3]int{
	var a = 1
	var b = 2
	var c = 3
	for a <= limit {
		if(a*a + b*b == c*c && a + b +c == limit){
			return [3]int{a, b, c}
		}
		c += 1 
		
		if(c == limit + 1){
			if(b == limit){
				b = limit
				c = limit
			}else{
				b += 1 
				c = b
			}
		}
		if(b == limit){
			a += 1
			b = a + 1
			c = b + 1
		}
		
	}
	return [3]int{1,1,1}
}
