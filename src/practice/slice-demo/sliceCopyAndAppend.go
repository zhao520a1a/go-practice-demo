package main

import "fmt"

/*
切片的复制和追加
 */
func main() {
	slFrom  := []int{1,2,3}
	slTo := make([]int,10)

	//复制
	n := copy(slTo,slFrom)
	fmt.Println("复制的元素数量：",n)
	fmt.Println(cap(slTo) ,"-",slTo)

	//追加的元素必须和原切片的元素同类型。如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。返回的切片可能已经指向一个不同的相关数组了。
	slTo = append(slTo,4,5,6,7)
	fmt.Println(cap(slTo), "-",slTo)

	slTo = AppendInt(slTo,8,9)
	fmt.Println(cap(slTo),"-", slTo)

}


// 手动实现append方法
func AppendInt(slice []int, elems ...int) []int{
	m := len(slice)
	n := m + len(elems)
	if(n > cap(slice)) {
		newSlice := make([] int, (n+1)*2) // 创建新的分片
		copy(newSlice,slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n],elems)
	return slice
}

