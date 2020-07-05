package main

import "fmt"

/*
切片是一个 长度可变的数组。
*/
func main() {
	var arr1 [6]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	fmt.Println("数组：", arr1)

	slice1 := arr1[2:4]
	fmt.Println("初始化：", slice1)

	slice1 = slice1[0:3]
	//slice1 = slice1[-1:3] -- 重新分片只能向后移动，不允许向前启动，否则会导致编译错误。
	fmt.Println("重新分片：", slice1)

	//切片可以反复扩展直到占据整个相关数组。
	slice1 = slice1[:cap(slice1)]
	fmt.Println("切片扩展到上限：", slice1)

	fmt.Println("传递切片给函数:", sum(arr1[:]))

	//用make创建切片
	slice2 := make([]int, 10, 50)
	for j := 0; j < len(slice2); j++ {
		slice2[j] = j
	}
	fmt.Println("make创建切片:", slice2)


	//切片常量
	slice3 := []string{"春","夏","秋","冬"}
	slice4 := [...]string{"春","夏","秋","冬"}
	fmt.Println(slice3)
	fmt.Println(slice4)


	//错误做法- 用new来创建切片 此时：*p == nil，切片的len和cap都为0
	var p *[]int = new([]int)
	fmt.Println("new 创建切片" ,*p)



}

//传递切片给函数
func sum(arr []int) int {
	s := 0
	for _, value := range arr {
		s += value
	}
	return s
}
