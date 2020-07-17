package main

import (
	"fmt"
	"time"
)

func main() {

	//date -> timestamp
	dateStr := "2020-07-15"
	begindate, err := time.ParseInLocation("2006-01-02", dateStr, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(" %d \n--\n", begindate.Unix())

	//dateStr1 := "2020-07-15 08:00:00"
	//begindate1, err := time.ParseInLocation("2006-01-02 00:00:00",  dateStr1, time.Local)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("\n--\n", begindate1.Unix())

	//timestamp -> date
	//timestamp := int64(1588118400)  // 2020-04-29 08:00:00
	timestamp := int64(1594828800)
	//timestamp := int64(1594828862)
	tt := time.Unix(timestamp, 0)
	fmt.Println(tt.Format("2006-01-02 15:04:05"))

}
