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

func testTime() {

	fmt.Println("-- now()得到的当前时间的时区跟电脑的当前时区一样")
	fmt.Println("0. now: ", time.Now())

	str := "2018-09-10 00:00:00"
	fmt.Println("str: ", str)
	fmt.Println("-- Parse 把时间字符串转换为Time，默认时区是UTC时区 ")
	t, _ := time.Parse(TIME_LAYOUT, str)
	fmt.Println("2. Parse time: ", t)

	fmt.Println("-- Format 把Time转换为时间字符串")
	tStr := t.Format(TIME_LAYOUT)
	fmt.Println("3. Format time str: ", tStr)

	fmt.Println("-- Zone方法可以获得变量的时区和时区与UTC的偏移秒数，应该支持夏令时和冬令时。")
	name, offset := t.Zone()
	name2, offset2 := t.Local().Zone()
	fmt.Printf("4. Zone name: %v, Zone offset: %v\n", name, offset)
	fmt.Printf("5. Local Zone name: %v, Local Zone offset: %v\n", name2, offset2)
	tLocal := t.Local()
	tUTC := t.UTC()
	fmt.Printf("6. t: %v, Local: %v, UTC: %v\n", t, tLocal, tUTC)
	fmt.Printf("7. t: %v, Local: %v, UTC: %v\n", t.Format(TIME_LAYOUT), tLocal.Format(TIME_LAYOUT), tUTC.Format(TIME_LAYOUT))

	fmt.Println("-- Unix不管Time变量存储的是什么时区，其Unix()方法返回的都是距离UTC时间：1970年1月1日0点0分0秒的秒数。")
	fmt.Printf("8. Local.Unix: %v, UTC.Unix: %v\n", tLocal.Unix(), tUTC.Unix())
	fmt.Println("-- Unix()返回的秒数可以是负数，如果时间小于1970-01-01 00:00:00的话。")
	str2 := "1969-12-31 23:59:59"
	t2, _ := time.Parse(TIME_LAYOUT, str2)
	fmt.Printf("9. str2：%v，time: %v, Unix: %v\n", str2, t2, t2.Unix())

	fmt.Printf("10. %v, %v\n", tLocal.Format(time.ANSIC), tUTC.Format(time.ANSIC))
	fmt.Printf("11. %v, %v\n", tLocal.Format(time.RFC822), tUTC.Format(time.RFC822))
	fmt.Printf("12. %v, %v\n", tLocal.Format(time.RFC822Z), tUTC.Format(time.RFC822Z))

	//指定时区
	fmt.Println("-- 指定时区，time.LoadLocation可以根据时区名创建时区Location，所有的时区名字可以在$GOROOT/lib/time/zoneinfo.zip文件中找到，解压zoneinfo.zip可以得到一堆目录和文件，我们只需要目录和文件的名字，时区名是目录名+文件名，比如\"Asia/Shanghai\"。中国时区名只有\"Asia/Shanghai\"和\"Asia/Chongqing\"，而没有\"Asia/Beijing\"。time.ParseInLocation可以根据时间字符串和指定时区转换Time。")
	parseWithLocation("America/Cordoba", str)
	parseWithLocation("Asia/Shanghai", str)
	parseWithLocation("Asia/Beijing", str)

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

	//lt, err := time.ParseInLocation(TIME_LAYOUT, timeStr, location)
	//if err != nil {
	//	println(err.Error())
	//}
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
