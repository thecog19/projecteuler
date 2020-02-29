package main

import (
        "time"
        //"math"
        "fmt"
        //"strconv"
)

func main() {
        start := time.Now()
        fmt.Printf("Result: %f", diff(100))
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

func diff(x int) int{
	var sum_of_squares_res = sum_of_squares(x)
	var square_of_sum_res = square_of_sum(x)
	var diff = square_of_sum_res - sum_of_squares_res
	return diff
}

func sum_of_squares(x int) int{
	var i = 0
	var total = 0
	for i < x {
		total += (i + 1) * (i + 1)
		i += 1
	}
	return total
}

func square_of_sum(x int)int{
	var i = 0
        var total = 0
        for i < x {
                total += i + 1
                i += 1
        }
        return total * total
}
