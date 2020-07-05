/*
结构体的定义
*/

package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName, lastName string
	int                 //匿名字段/内嵌字段
	contact
}

type contact struct {
	telephone string
	email     string
}

func main() {
	//pers1是结构体类型变量
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	pers1.int = 1
	pers1.telephone = "10086"
	pers1.email = "123@qq.com"

	//调用函数
	upPerson(&pers1)
	fmt.Println(pers1)

	//调用方法1
	fullName := pers1.fullName("-")
	fmt.Println(fullName)

	//调用方法2:指针调用方法
	var pers2 *Person
	pers2 = &pers1
	fmt.Println(pers2.firstName)
	//调用方法3：值调用方法
	pers1.SetFirstName("Eric")
	fmt.Println(pers1.firstName)

	//调用内嵌对象的方法
	pers1.changeTelephone("10000")
	fmt.Println(pers1)

}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

/*
方法
*/

//值传递
func (this Person) fullName(segment string) string {
	return this.firstName + segment + this.lastName
}

//将指针作为接收者,引用传递；Go会帮我们自动解引用，不过调用者是值还是指针，方法都支持运行
// setter - firstName
func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}

// getter - firstName
func (p *Person) FirstName() string {
	return p.firstName
}


func (this *contact) changeTelephone(newTeleNum string) {
	this.telephone = newTeleNum
}

func (this *Person) changeTelephone(newTeleNum string) {
	fmt.Println("覆盖匿名类型的同名方法")
	this.telephone = newTeleNum
}

//func (this *Person) String() string {
//	return "(" + this.firstName + "-" + this.lastName + ")"
//}
