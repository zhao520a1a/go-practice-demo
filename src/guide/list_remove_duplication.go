package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
去除slice和list的重复元素，非常有用的功能
*/

func main() {
	b := []string{"a", "b", "c", "c", "e", "f", "a", "g", "b", "b", "c"}
	sort.Strings(b)
	fmt.Println(Duplicate(b))

	c := []int{1, 1, 2, 4, 6, 7, 8, 4, 3, 2, 5, 6, 6, 8}
	sort.Ints(c)
	fmt.Println(Duplicate(c))
}
func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}
