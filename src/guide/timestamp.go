package main

import (
	"fmt"
	"time"
)

func main() {

	//date -> timestamp
	dateStr := "2020-08-11"
	begindate, err := time.ParseInLocation("2006-01-02", dateStr, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(" %d \n--\n", begindate.Unix())

	//timestamp -> date
	timestamp := int64(1595284492) //2020-07-21 06:34:52
	//timestamp := int64(1583229422)
	//timestamp := int64(1597473988)
	tt := time.Unix(timestamp, 0)
	fmt.Println(tt.Format("2006-01-02 15:04:05"))

	//newDay := time.Now().AddDate(-1 /* years */, 0 /* months */, 0 /* days */)
	//fmt.Println(newDay)LoadLocation

}
