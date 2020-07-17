package main

import (
	"fmt"
)

func main() {
	//var a *int

	// 测试fmt.Sprintln对类型的普适性
	//slice1 :=  []int{
	//	1,3,4,
	//}
	//slice2 := []string{
	//	"你好","世界",
	//}
	//fmt.Print(fmt.Sprintln("ss",123,true,3.242,slice1,slice2,nil))

	//arrSlice := make([]int64)
	//arrSlice :=  []int{}
	//for _, value := range arrSlice {
	//	fmt.Println("--" , value)
	//}

	// test map capcity len
	//var testmap = make(map[int]int,10)
	//testmap[1] = 1
	//testmap[2] = 1
	//testmap[3] = 1
	//fmt.Printf("len:%d",len(testmap))

	// test point arr foreach
	var items []int = nil
	for _, data := range items {
		fmt.Printf("%v", data)
	}
	fmt.Printf("%v", "finish")
	fmt.Printf("%v", append(items, 1))

}
