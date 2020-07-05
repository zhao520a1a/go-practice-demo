/*
协程的同步： 关闭通道-测试阻塞通道
 */
package main

import "fmt"

func main() {
	ch := make(chan int)
	go sendData(ch)
	getData(ch)
}

func getData(ch chan int) {
	//使用for range读channel
	//for i := range ch {
	//	fmt.Println(i)
	//}

	//使用_,ok判断channel是否关闭
	for{
		if i,ok := <- ch; ok {
			fmt.Println(i)
		} else {
			fmt.Println("receive close chan msg")
			break
		}
	}
}

func sendData(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
}