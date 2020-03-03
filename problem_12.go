package main

import (
        "time"
	"math"
        "fmt"
)

func main() {
	start := time.Now()
	fmt.Printf("Result: %f", find_thrshold_number(500))
	elapsed := time.Since(start)
	fmt.Printf("\ntook %s", elapsed)
}

func sieve(x int) []int{
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
	for  i < x{
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
	return prime_array
}

func find_all_divisors(x int) []int{
	var limit int = int(math.Round(math.Sqrt(float64(x))+1))
	prime_factors := sieve(limit)
	prime_factors = append(prime_factors, 1)
	var divisors []int
	for _, factor := range prime_factors {
		i := factor
		for i < limit {
			if(math.Mod(float64(x),float64(i)) == 0 && !contains(divisors, i)){
				divisors = append(divisors, i)
				divisors = append(divisors, x/i)
			}
			i += factor
		}
	}
	return divisors
}

func generate_next_triangle_number(curr_number int, last_increment int) map[string]int{
	increment := last_increment + 1
	m := make(map[string]int)
	m["curr_number"] = curr_number + increment
	m["last_increment"] = increment
	return m
}

func find_thrshold_number(x int) int{
	var triangle_number map[string]int
	triangle_number = map[string]int{"curr_number": 0, "last_increment":0}
	divisors := 0
	for divisors  <= x  {
		triangle_number = generate_next_triangle_number(triangle_number["curr_number"], triangle_number["last_increment"])
		div_arr := find_all_divisors(triangle_number["curr_number"])
		divisors = len(div_arr)
	}
	return triangle_number["curr_number"]
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
