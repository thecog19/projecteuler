package main

import (
	"io/ioutil"
        "time"
        //"math"
        "fmt"
	"sort"
	"strings"
)

func main() {
        start := time.Now()
	dat, err := ioutil.ReadFile("./names.txt")
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Printf("Result: %f", generate_scores(strings.Replace(string(dat),"\"", "", -1)))
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

func generate_scores(s string) int{
	arr :=  strings.Split(s, ",")
	sort.Strings(arr)
	total_sum := 0
	score := 0
	for i,n := range(arr){
		for _, l := range(n){
			score += letter_value(string(l))
		}
		score += i*score
		total_sum += score
		//fmt.Printf("\n%f, score: %d", n, score)
		score = 0
	}
	return total_sum
}

func letter_value(letter string) int {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(alphabet, letter) + 1
}


