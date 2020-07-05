/*
使用选择依赖注入方式避免使用可变的全局变量
 */
package main

import (
	"time"
)

// sign.go
var _timeNow = time.Now


type signer struct {
	now func() time.Time
}

//使用选择依赖注入方式避免改变全局变量,这样既适用于函数指针又适用于其他值类型
func newSigner() *signer {
	return &signer{
		now: time.Now,
	}
}

func (s *signer) Sign() string {
	//now := _timeNow()   避免可变全局变量
	now := s.now()
	return now.String()
}

func main() {

	s := newSigner()
	//更新变量
	s.now = func() time.Time {
		return time.Now()
	}

}