package main

import (
        "time"
        //"math"
        "fmt"
        //"strconv"
        "github.com/divan/num2words"
        "strings"
)

func main() {
        start := time.Now()
        fmt.Printf("Result: %f", count_chars(1000))
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

func count_chars(limit int) int{
       i := 1
       total := 0
       for i <= limit {
                total += len_string(num2words.ConvertAnd(i))
                i++
       }
       return total 
}

func len_string(str string) int{
        new_str := strings.Join(strings.Split(str, " "),"")
        new_str = strings.Join(strings.Split(new_str, "-"), "")
        return len(new_str)
}
