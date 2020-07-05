/*
用函数可以作为（其它函数的）参数的事实来使用高阶函数

 */
package main

import (
	"fmt"
	"strconv"
)

type Any interface {

}

type Car struct {
	Mode string //车型
	Facturers string //厂商
	BuildYear int //生产年份
	//...
}

//car集合
type Cars []*Car


func main() {
	ford := &Car{"Fiesta","福特", 2008}
	bmw  := &Car{"XL 450", "宝马", 2011}
	merc := &Car{"D600", "梅赛德斯", 2009}
	bmw2 := &Car{"X 800", "宝马", 2008}
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})

	//查找
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Facturers == "宝马") && (car.BuildYear > 2010)
	})
	fmt.Println("全部数据: ", allCars)
	fmt.Println("查找结果: ", allNewBMWs)

	//获取其他类型的集合
	facturerArr := allCars.Map(func(car *Car) Any {
		return car.Facturers
	})
	//将[]Any转换为[]string
	facturers := make([]string,0)
	for _, v := range facturerArr {
		if s,ok := v.(string); ok {
			facturers = append(facturers,s)
		}
	}
	fmt.Println("全部厂家: ",  facturerArr)

	//按车辆厂家对Cars集合做归类,返回Map集合
	//facturers := []string{"福特", "梅赛德斯", "路虎", "宝马", "奥迪"}
	sortedAppender, sortedCars := MakeSortedAppender(facturers)
	allCars.Process(sortedAppender)
	BMWCount := len(sortedCars["宝马"])
	fmt.Println("最终结果: ", sortedCars)
	fmt.Println("宝马厂家的车辆数量：", BMWCount)
}



func (c Car) String() string {
	return strconv.Itoa(c.BuildYear) + "-" + c.Facturers + "-"+c.Mode
}


//功能：通用函数，执行入参函数；  -- 函数入参
func (cs Cars) Process(f func(car *Car))  {
	for _, c := range cs {
		f(c)
	}
}

//功能：将满足条件的找出来，添加到集合中。 -- 将有返回值的函数入参
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car,0)
	cs.Process(func(c *Car) {
		if(f(c)) {
			cars = append(cars,c)
		}
	})
	return cars
}

//功能： 通过cars集合获取其他类型结合 -- 将外层函数作为变量传入高级函数中
func (cs Cars) Map(f func(car *Car) Any) []Any  {
	result := make([]Any,0)
	ix := 0
	cs.Process(func(c *Car) {
		result = append(result,f(c))
		ix ++
	})
	return result
}



//产生特定的添加函数(根据不同的厂商添加汽车到不同的集合)和 map集合 (等价于Java中Map<String,List<Car>>集合)
func MakeSortedAppender(facturers []string) (func(car *Car), map[string]Cars) {
	//初始化carMap及其key值，用于后续使用
	carMap := make(map[string]Cars)
	carMap["其他"] = make([] *Car,0)
	for _, v := range facturers {
		carMap[v] = make([]*Car,0)
	}

	//添加函数
	appender := func(c *Car) {
		if _,ok:=carMap[c.Facturers] ; ok {
			carMap[c.Facturers] = append(carMap[c.Facturers],c)
		} else {
			carMap["其他"] = append(	carMap["其他"],c)
		}
	}
	return appender,carMap
}


