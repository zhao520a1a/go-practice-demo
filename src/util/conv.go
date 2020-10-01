package util

import (
	"encoding/json"
	"errors"
	"fmt"
)

func I64ToBool(i int64) bool {
	if i == 0 {
		return false
	}
	return true
}

func BoolToI64(b bool) int64 {
	if b == true {
		return 1
	}

	return 0
}

func Int64ArrToJson(data []int64) string {
	if len(data) == 0 {
		return EmptyJson()
	}
	jsonStr := EmptyJson()
	if data, err := json.Marshal(data); err == nil {
		jsonStr = string(data)
	}
	return jsonStr
}

func JsonToInt64Arr(jsonStr string) (result []int64, err error) {
	if len(jsonStr) == 0 {
		return
	}
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(" get Unmarshal err:%s , json:%s", err, jsonStr))
	}
	return
}

func EmptyJson() string {
	return "[]"
}

func AreaNameToAreaId(areaName string) int {
	switch areaName {
	case "beijing":
		return 1
	case "wuhan":
		return 2
	default:
		return 0
	}
}
