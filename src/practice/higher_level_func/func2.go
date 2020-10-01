package main

import (
	"fmt"
)

func main() {

	//用内置函数的方式经进行变量赋值
	flag := false
	res, err := func() (string, error) {
		if flag != true {
			return "", fmt.Errorf("XXXX")
		}
		return "success", nil
	}()
	fmt.Printf("res:%s err:%v", res, err)

}
