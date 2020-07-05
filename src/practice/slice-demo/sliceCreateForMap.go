package main

import "fmt"

/*
创建map类型的切片
	1.分配切片
	2.分配切片中每个元素

*/
func main() {

	items := make([]map[int]string, 5)

	//通过索引使用切片的 map-demo 元素
	for i := range items {
		items[i] = make(map[int]string)
		items[i][1] = "a"
	}
	fmt.Printf("items: %v \n", items)

	//错误的方式： item只是map的一个拷贝, 实质切片中的map没有得到初始化。
	for _, item := range items {
		item = make(map[int]string)
		item[1] = "b"
	}
	fmt.Printf("items: %v \n", items)

}
