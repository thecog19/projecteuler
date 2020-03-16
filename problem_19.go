package main

import (
        "time"
        //"math"
        "fmt"
	//"strconv"
)

func main() {
        start := time.Now()
	fmt.Printf("Result: %f", count_days("Tue Jan 01 11:01:01 -0700 1901", "Sun Dec 31 11:01:01 -0700 2000"))
        elapsed := time.Since(start)
        fmt.Printf("\ntook %s", elapsed)
}

func count_days(start_date_s string, end_date_s string) int{
	curr_date, e := time.Parse(time.RubyDate , start_date_s)
	end_date, e2 := time.Parse(time.RubyDate, end_date_s)

	if(e2 != nil){
		fmt.Println(e2)
	}
	if(e != nil){
		fmt.Println(e)
	}
	count := 0
	for curr_date.Before(end_date) {
		if(curr_date.Weekday() == time.Weekday(0)){
			count += 1
		}
		curr_date = curr_date.AddDate(0, 1, 0)
	}
	return count
}
