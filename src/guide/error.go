/*

Go 中有多种声明错误（Error) 的选项：
errors.New 对于简单静态字符串的错误
fmt.Errorf 用于格式化的错误字符串
实现 Error() 方法的自定义类型
用 "pkg/errors".Wrap 的 Wrapped errors


如果不需要额外信息的简单错误,则使用errors.New 足够了。
如果客户需要检测并处理此错误，则使用自定义类型并实现该 Error() 方法。
如果正在传播下游函数返回的错误，则使用错误包装Error Wrapping 手法
如果调用者不需要检测或处理的特定错误情况等其他情况，则使用 fmt.Errorf。
*/
package main

import (
	"errors"
	"fmt"
)

//不需要额外信息的简单错误
var ErrCouldNotOpen = errors.New("could not open1")

func open1(file string) error {
	return ErrCouldNotOpen
}

//可能需要客户端检测的错误，并且想向其中添加更多信息，使用自定义类型并实现该 Error() 方法。
type errNotFound struct {
	file string
}

var _ error = (*errNotFound)(nil)

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

func open2(file string) error {
	return errNotFound{file: file}
}

//最好公开匹配器功能以检查错误。
func IsNotFoundError(err error) bool {
	_, ok := err.(errNotFound)
	return ok
}

//Error Wrapping
/*
错误具有3种传播方式
1. 不需要添加上下文并想维护原始错误类型，则返回原始错误
2. 需要添加上下文，使用 errors.Wrap 以便错误消息提供更多上下文 ,errors.Cause 可用于提取原始错误。
3. 如果调用者不需要检测或处理的特定错误情况，则使用fmt.Errorf
*/
func open3(file string) error {
	//请避免使用“failed to”之类的短语以保持上下文简洁，这些短语会陈述明显的内容，但是，一旦将错误发送到另一个系统，就应该明确消息是错误消息（例如使用err标记，或在日志中以”Failed”为前缀）。
	err := open1(file)
	if err != nil {
		return fmt.Errorf(
			"new store: %s", err)
	}
	return nil
}

func main() {
	file := "testfile.txt"

	if err := open1(file); err != nil {
		if err == ErrCouldNotOpen {
			//handle
			fmt.Println(err)
		} else {
			panic("unkonwn error")

		}
	}

	if err := open2(file); err != nil {
		if IsNotFoundError(err) {
			//handle
			fmt.Println(err)
		} else {
			panic("unkonwn error")
		}
	}

	if err := open3(file); err != nil {
		fmt.Println(err)
	}

}
