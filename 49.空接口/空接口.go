/*
空接口
	不包含任何方法的接口
	所有类型都实现了空接口，在创建对象的时候，任何类型都可以实例化成空接口
	因此空接口可以存储任意类型的数值

空接口理解
	因为空接口没有方法，所以任何类型都实现了空接口
	也就是任何对象都有了两种形态：一种是空接口类型，一种是自己的类型
	所以空接口类型创建的变量，可以接受任何类型的数据
*/

package main

import "fmt"

// MyPrint 接收任意类型作为参数的函数
func MyPrint(a empty) {
	fmt.Println(a)
}

func main() {
	// 实例化成空接口对象
	var a1 empty = Cat{"black"}
	fmt.Println(a1)

	// 实例化成自己的对象
	var cat = Cat{"white"}
	fmt.Println(cat)

	// 其他类型实例化成自己的类型
	var a2 empty = Person{"王二狗",30}
	var a3 empty = "hello"
	var a4 empty = 100
	fmt.Println(a2,a3,a4)

	// 调用空接口函数
	MyPrint(a1)
	MyPrint(a2)
	MyPrint(a3)
	MyPrint(a4)

	// 定义存储任意类型的容器
	type list []interface{}

	lst := list{"rs",20,"male"}
	fmt.Println(lst)
}

// 定义空接口
type empty interface {}

type Person struct {
	name string
	age int
}

type Cat struct {
	color string
}
