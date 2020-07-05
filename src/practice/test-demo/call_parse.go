package main

import (
	parse2 "defer-panic-recover"
	"fmt"
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		//fmt.Printf("Parsing %q:\n  ", ex)
		nums, err := parse2.Parse(ex)  // 向包的调用者返回错误值（而不是 panic）。
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}

}