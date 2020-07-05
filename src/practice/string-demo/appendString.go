package main

import (
	"bytes"
	"fmt"
)

/*
通过buffer串联字符串
*/
func main() {
	s1 := "你好，"
	s2 := "世界！"

	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	fmt.Println(buffer.String())

}
