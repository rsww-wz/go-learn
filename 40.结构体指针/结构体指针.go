// 深浅拷贝
/*
Go语言中所有赋值都是把数据变量的数据拷贝一份
但是有的变量的类型是值类型，有的是引用类型，拷贝的内容就不一样

深拷贝
	值类型，拷贝实际数据

浅拷贝
	引用类型，拷贝的是实际数据的地址

结构体的拷贝是深拷贝，也就是结构体拷贝的实际的数据
但是如果结构体内部的字段是引用类型，拷贝的也是引用类型

 */

// 结构体指针
/*
定义结构体指针
	name *结构体类型(结构体名称)

通过指针访问结构体数据
	正常访问：(*p).name
	简化写法：p.name
 */

// new
/*
通过new函数创建结构体指针

new函数本质上说根据其他语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间
并且返回其地址，即一个*T类型的值。用Go语言的术语说，它返回了一个指针，指向新分配的类型T的零值。
非常重要的一点：new返回指针

可以通过new函数创建Go语言中各种数据类型的指针
指向的并不是nil，指向新分配类型的零值
 */

package main

import "fmt"

func main() {
	type Student struct {
		name string
		age int
		address string
	}

	// 深拷贝
	std1 := Student{"王二狗",49,"广州"}
	std2 := std1
	fmt.Printf("%p,%T\n",&std1,std1)
	fmt.Printf("%p,%T\n",&std2,std2)

	std1.name = "李小花"
	fmt.Println(std1)
	fmt.Println(std2)

	// 结构体指针
	var p *Student
	fmt.Printf("%T\n",p)

	// 结构体赋值
	p = &std1

	// 正常访问结构体数据
	fmt.Println((*p).name)

	// 简化写法
	fmt.Println(p.name)

	// 修改属性
	p.age = 20

	// 三种写法等价
	fmt.Println(p.age)
	fmt.Println((*p).age)
	fmt.Println(std1.age)

	// 通过new创建结构体指针,创建的是这种类型的指针
	p1 := new(Student)
	fmt.Printf("%v,%T\n",p1,p1)

	// 赋值
	p1 = &std2
	p1.address = "北京"
	fmt.Println(*p1)
}
