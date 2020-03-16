package main

import (
        "time"
	"math"
        "fmt"
)

func main() {
	start := time.Now()
	fmt.Printf("Result: %f", sum_of_amicable(10000))
	elapsed := time.Since(start)
	fmt.Printf("\ntook %s", elapsed)
}

func sum_of_divisors(x int) int {
	divisors := find_all_divisors(x)
	sum := 0
	for _, d := range(divisors){
		if(d != x){
			sum += d
		}
	}
	return sum
}

func amicable(x int) bool {
	sum_x := sum_of_divisors(x)
	if(x == sum_x){
		return false
	}else{
		return x == sum_of_divisors(sum_x)
	}
}

func sum_of_amicable(limit int) int {
	i := 1
	sum := 0
	for i < limit {
		if(amicable(i)){
			sum += i
		}
		i++
	}
	return sum
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


func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
