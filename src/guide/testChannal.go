package main

import "fmt"

func main() {

	done := make(chan string)
	c := make(chan string)

	go func() {
		fmt.Println("########")
		s := <-c
		fmt.Println("--" + s)
		fmt.Println("%%%%%%")

		s1 := <-c
		fmt.Println("--" + s1)
		close(done)
		//done <- " Bye!"
	}()

	c <- "hello world"
	fmt.Println("@@@@@@@")
	c <- "hello world123"
	<-done
	<-done
	<-done
}
