package parse

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e!= nil {
			fmt.Println("Panicing",e)
		}
	}()

	badCall()

	fmt.Println("After bad call")
}

func main() {
	test()
}


