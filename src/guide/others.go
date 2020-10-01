package main

import (
	"fmt"
)

func main() {
	var backendApiList1 []string
	var backendApiList []string
	backendApiList = append(backendApiList, backendApiList1...)

	fmt.Println("--", len(backendApiList))

}
