//零值 sync.Mutex 和 sync.RWMutex 是有效的，指向 mutex 的指针基本是不必要的。
//因此，mutex不建议以指针的形式操作；
package main

import "sync"

func main() {
	mu1 := new(sync.Mutex) //不必要
	mu1.Lock()

	var mu2 sync.Mutex
	mu2.Lock()
}

