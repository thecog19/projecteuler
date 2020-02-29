package main

import (
        "time"
        //"math"
        "fmt"
	"strconv"
)

func main() {
        start := time.Now()
        fmt.Printf("Result: %f", largest_palindrome())
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

func largest_palindrome() int{
	var first = 100
	var second = 100
	var largest = 0
	for first < 1000 {
		for second < 1000{
			if(is_palindrome(strconv.Itoa(first * second)) && first*second > largest){
				largest = first*second
			}
			second += 1
		}
		second = 100
		first += 1
	}
	return largest
}


func is_palindrome(num string) bool{
	var length = len(num)
	var p1 = 0
	var p2 = length - 1
	for p1 < p2 {
		if(num[p1] != num[p2]){
			return false
		}
		p1 += 1
		p2 -= 1
	}
	return true
}
