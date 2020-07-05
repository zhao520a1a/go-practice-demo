package main

import (
	"fmt"
	"interface/classifier"
)

//定义接口
type Shaper interface {
	Area() float32
}


type Square struct {
	side float32
}

//实现接口
func (sq *Square)Area() float32  {
	return sq.side * sq.side
}

type Rectangle struct {
	length,width float32
}

//实现接口
func (r Rectangle) Area() float32{
	return r.length * r.width
}


func main() {
	sq := &Square{5}
	rect := Rectangle{5,3}

	var areaIntf Shaper
	areaIntf = rect
	fmt.Println(areaIntf.Area())
	areaIntf = sq  //多态：接口变量包含一个指向 Square 变量的引用，通过它可以调用 Square 上的方法 Area()
	fmt.Println(areaIntf.Area())

	//接口断言：检测和转换接口变量的类型
	if t,ok := areaIntf.(Rectangle); ok {
		fmt.Printf("the classifier of areaIntf is %T \n",t)
	}
	fmt.Println(areaIntf.(*Square).Area())

	// 类型判断  可以用 classifier-switch 进行运行时类型分析，但是在 classifier-switch 不允许有 fallthrough
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type：%T - Value: %v \n",t,t)
	case Rectangle:
		fmt.Printf("Type：%T - Value: %v \n",t,t)
	case nil:
		fmt.Printf("nil value : nothing to check \n")
	default:
		fmt.Printf("Unexpected classifier %T \n",t)
	}


	shapes := []Shaper{sq,rect}
	for i := range shapes {
		fmt.Println(shapes[i] , " - " , shapes[i].Area())
	}


    //类型分类函数
	classifier.Classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)

}
