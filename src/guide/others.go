package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	// 测试fmt.Sprintln对类型的普适性
	//slice1 :=  []int{
	//	1,3,4,
	//}
	//slice2 := []string{
	//	"你好","世界",
	//}
	//fmt.Print(fmt.Sprintln("ss",123,true,3.242,slice1,slice2,nil))


	//测试Json转码
	//Studentuidsjson := ""
	//if data, err := json.Marshal([]int64{}); err == nil {
	//	Studentuidsjson = string(data)
	//}
	//fmt.Println(Studentuidsjson)

	//res, err := TransJsonToInt64Arr("")
	//fmt.Printf("%v -- %s",res, err)
	//fmt.Printf("%v - %s",result,err)

	//result,err := TransInt64ArrToJson(nil,"dd")
	//fmt.Printf("%v - %s",result,err)

	//fmt.Printf("-- %v --", TransInt64ArrToJson(nil))
	//fmt.Printf("-- %s --", TransEmptyJson())


	//arrSlice := make([]int64)
	arrSlice :=  []int{}
	for _, value := range arrSlice {
		fmt.Println("--" , value)
	}

}

func TransInt64ArrToJson(data []int64) string {
	if len(data) == 0 {
		return TransEmptyJson()
	}
	jsonStr := TransEmptyJson()
	if data, err := json.Marshal(data); err == nil {
		jsonStr = string(data)
	}
	return jsonStr
}
func TransJsonToInt64Arr(jsonStr string) (result []int64, err error) {
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(" get Unmarshal err:%s", err))
	}
	return
}



func TransEmptyJson() string {
	emptyJson := ""
	if data, err := json.Marshal([]int64{}); err == nil {
		emptyJson = string(data)
	}
	return emptyJson
}