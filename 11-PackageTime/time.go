package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println("Current time:", now)

	var utc time.Time = time.Date(2009, time.August, 17,0,0,0,0, time.UTC)
	fmt.Println("UTC time:", utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 08:25:00"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed time:", valueTime)
	fmt.Println("Formatted time:", valueTime.Format(formatter))
	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())

}
