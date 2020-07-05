package main

import "fmt"

func main() {

	s := "hello"
	c := []byte(s)

	c[1] = 'd'
	s2 := string(c)
	fmt.Println(s2)
	fmt.Println(s)



}
