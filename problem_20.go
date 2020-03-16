package main

import (
        "time"
        "math/big"
        "fmt"
	"strconv"
)

func main() {
        start := time.Now()
	fmt.Printf("Result: %f", factorial_count(100))
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

func factorial_count(start int64) int {
	i := start-1
	var total *big.Int
	total = big.NewInt(start)
	for i > 0 {
		total = total.Mul(total,big.NewInt(i))
		i--
	}
	s := fmt.Sprintf("%f", total)
        sum := 0
        for _,e := range s {
                 val, _ := strconv.Atoi(string(e))
                 sum += val
         }
	return(sum)
}
