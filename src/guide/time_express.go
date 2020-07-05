//在处理时间时始终使用 "time" 包，因为它有助于以更安全、更准确的方式处理时间
package main

import (
	"fmt"
	"time"
)

/*
- 对外部系统使用 time.Time 和 time.Duration
尽可能在与外部系统的交互中使用 time.Duration 和 time.Time 例如 :
Command-line 标志: flag 通过 time.ParseDuration 支持 time.Duration
JSON: encoding/json 通过其 UnmarshalJSON method 方法支持将 time.Time 编码为 RFC 3339 字符串
SQL: database/sql 支持将 DATETIME 或 TIMESTAMP 列转换为 time.Time，如果底层驱动程序支持则返回
YAML: gopkg.in/yaml.v2 支持将 time.Time 作为 RFC 3339 字符串，并通过 time.ParseDuration 支持 time.Duration。
*/


//当不能在这些交互中使用 time.Duration 时，请使用 int 或 float64，并在字段名称中包含单位。 eg: {"intervalMillis": 2000}
type Config struct {
	IntervalMillis int `json:"intervalMillis"`
}

//当不能在这些交互中使用 time.Time 时，除非达成一致，否则使用 string 和 RFC 3339 中定义的格式时间戳。默认情况下，Time.UnmarshalText 使用此格式，并可通过 time.RFC3339 在 Time.Format 和 time.Parse 中使用。

func main() {

	//poll(10 * time.Second)

	t := time.Now()
	//获取24h后的时刻
	maybeNewDay := t.Add(24 * time.Hour)
	fmt.Println(maybeNewDay)

	//获取下一个日历日
	newDay := t.AddDate(0 /* years */, 0 /* months */, 1 /* days */)
	fmt.Println(newDay)

}

//处理时间点 time.time
func isActive(now,start,stop time.Time) bool {
	return (start.Before(now) || start.Equal(now)) && now.Before(stop)
}


//处理时间段 time.Duration
func poll(delay time.Duration) {
	for{
		//...
		time.Sleep(delay)
	}
}




