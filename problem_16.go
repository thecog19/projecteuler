package main

import (
        "time"
        "math"
        "fmt"
	"strings"
	"strconv"
)

func main() {
        start := time.Now()
	big_num := math.Pow(2,1000)
	s := fmt.Sprintf("%f", big_num)
	s = strings.Split(s, ".")[0]
	sum := 0
	for _,e := range s {
		 val, _ := strconv.Atoi(string(e))
       		 sum += val
	 }
        fmt.Printf("Result: %f",sum)
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

