package main

import (
	"fmt"
	"time"
)

func main() {

	//date -> timestamp
	dateStr := "2020-06-29"
	begindate, err := time.ParseInLocation("2006-01-02",  dateStr, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("--", begindate.Unix())

	//timestamp -> date
	timestamp := int64(1593360000)
	tt := time.Unix(timestamp, 0)
	fmt.Println(tt.Format("2006-01-02 15:04:05"))
}
