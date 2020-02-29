package main

import (
        "time"
	//"math"
        "fmt"
)

func main() {
	start := time.Now()
        //sieve(90000000)
	fmt.Printf("Result: %f", sieve(900000000))
	elapsed := time.Since(start)
	fmt.Printf("\ntook %s", elapsed)
}

func sieve(x int) int{
	//in this array each index maps to an int
	//false means isn't a prime factor
	//true means it is
	// assume all of them are 
	var prime_array []int
	//var limit int = int(math.Round(math.Sqrt(float64(x))+1))
	array_of_ints := make([]bool, x)
	for i := range  array_of_ints {
		array_of_ints[i] = true
	}
	var i = 2
	for  len(prime_array) < 10002{
		fmt.Printf("\n %f", len(prime_array))
		if(array_of_ints[i]){
			prime_array = append(prime_array, i)
			if(i < x){
				var j = i
				for j < x{
					array_of_ints[j] = false
					j += i
				}
			}
		}
		i += 1
	}
	return prime_array[10000]
}

