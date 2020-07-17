package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	//测试Json转码

	res, err := TransJsonToInt64Arr("")
	fmt.Printf("%v -- %s", res, err)
	fmt.Printf("-- %v --", TransInt64ArrToJson(nil))
	fmt.Printf("-- %s --", TransEmptyJson())
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
	//emptyJson := ""
	//if data, err := json.Marshal([]int64{}); err == nil {
	//	emptyJson = string(data)
	//}
	//return emptyJson
	return "[]"
}
