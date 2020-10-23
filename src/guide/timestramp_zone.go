package main

import (
	"fmt"
	"os"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func parseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		//zone, offset := time.Now().In(location).Zone()
		lt, err := time.ParseInLocation(TIME_LAYOUT, timeStr, l)
		if err != nil {
			println(err.Error())
		}
		fmt.Println(locationName, lt)
		return lt, nil
	}
}

func transformTimestamp(timestamp int64, locationName string) (target int64, err error) {
	//设置时区
	location, err := time.LoadLocation(locationName)
	if err != nil {
		println(err.Error())
		return
	}

	//parse time
	tt := time.Unix(timestamp, 0).In(location)

	zone, offset := tt.Zone()
	tStr := tt.Format(TIME_LAYOUT)
	fmt.Printf("\n-- zone:%s, offset:%d tStr:%s \n ", zone, offset, tStr)

	target = tt.Unix()
	return
}

func main() {
	//source := time.Now().Unix()
	//str := "2020-03-11 00:00:00"   //美国夏令时（3月11日至11月7日）
	str := "2020-01-01 00:00:00" //美国夏令时（3月11日至11月7日）
	fmt.Println("str: ", str)
	fmt.Println("-- Parse 把时间字符串转换为Time，默认时区是UTC时区 ")
	t, _ := time.Parse(TIME_LAYOUT, str)
	source := t.Unix()

	//设置环境变量，
	os.Setenv("ZONEINFO", "/Users/Golden/Documents/GoProjects/practice-demo/src/source/zoneinfo.zip")

	locationNames := []string{
		"Asia/Shanghai",   //	+08:00	+08:00
		"America/Creston", //		−07:00	−07:00

		//夏令时城市
		"America/Anchorage", //	−09:00	−08:00
		"America/Dawson",    //	−08:00	−07:00
	}

	for _, name := range locationNames {
		target, err := transformTimestamp(source, name)
		fmt.Printf("locationName:%s, source: %d, res: %d, err:%v", name, source, target, err)
	}
}
