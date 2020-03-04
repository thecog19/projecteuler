package main

import (
        "time"
        "fmt"
	"math"
)
var memoized map[int]int
func main() {
	start := time.Now()
	memoized = make(map[int]int)
	fmt.Printf("Result: %f", longest_sequence(1000000))
	elapsed := time.Since(start)
	fmt.Printf("\ntook %s\n", elapsed)
}

func longest_sequence(limit int) int {
	i := 1
	max := 0
	max_creator := 0
	for i < limit {
		length := generate_sequence(i)
		if(length > max){
			max = length
			max_creator = i
		}
	//	fmt.Printf("\n length:  %d, i: %d" ,length, i)
		i++
	}
	return max_creator
}

func generate_sequence(x int) int{
	//input a number, get a sequence length
	curr := x
	count := 1
	next := 0
	//test_count := 0
	for curr != 1 {
		if(math.Mod(float64(curr),2) == 0){
			next = curr/2
		}else{
			next = curr*3 + 1
		}
		mem_val, exists := memoized[next]
		if(exists){
			count = mem_val+count
			break
		}else{
			count += 1
			curr = next
		}

	}
	//if(test_count != count || test_count != 0){
	//	fmt.Printf("\nint: %d test: %d, real %d",x, test_count, count)
	//}
	memoized[x] = count
	return count
}
