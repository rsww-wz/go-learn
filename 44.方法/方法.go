/*
Go语言中同时有函数和方法

一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针
所以给定类型的方法属于该类型分方法集

方法只是一个函数，它带有一个特殊的接收器类型，它是func关键字和方法名之间编写的
接收器可以是struct类型或者非struct类型，接收方可以在方法内部方法

语法
	方法：func (t Type) methodName (parameter list) (return list) {}
	函数：func funcName (parameter list) (return list) {}

使用语法
	接受者：可以是类型，也可以是这个类型的指针
	区别
		接受者是类型:每次执行的时候，这个对象会重新拷贝一份，再传递给函数
		接受者是指针:每次执行的时候，用的同一个数据，不会重新拷贝

	Go语言中，定义变量不使用是不能通过编译的
	如果接受者的变量用不到，可以用下划线代替

理解
	方法的接收器就好像面向对象的self,谁调用这个方法，谁就是这个接收器(self)
	接收者：name 类型

	方法的名字可以一样，接收器不一样就可以了，函数的名字不可以一样
	相同名字的方法有点像鸭子类型，为不同的类型设计相同的方法名，调用的时候执行的是对应类型的方法



对比函数
	意义
		方法：某个类型的行为功能，需要指定接受者调用
		函数：一段独立的功能代码，可以直接调用
	语法
		方法：方法名可以相同，只要接受者不同
		函数：命名不能冲突

 */

package main

import "fmt"

func main() {
	// 创建对象
	w1 := Worker{"王二狗",30,"男"}

	// 创建指针类型对象
	w2 := &Worker{"王铁柱",36,"男"}
	fmt.Printf("%T\n",w2)

	// 结构体对象和结构体指针类型对象都可以调用方法
	w1.work()
	w2.work()

	fmt.Println("类型是结构体，修改后的数据:",w1.age)
	fmt.Println("类型是结构体，修改后的数据:",w2.age)

	// 接受者是指针会修改数据
	w1.rest()
	w2.rest()
	fmt.Println("类型是结构体指针，修改后的数据:",w1.age)
	fmt.Println("类型是结构体指针，修改后的数据:",w2.age)
}

// 定义一个工人结构体
type Worker struct {
	name string
	age int
	gender string
}

// 定义方法
// 只能由Worker创建的对象调用这个方法，绑定类Worker这个结构体(类)
func (w Worker) work() {
	fmt.Println(w.name,"在工作")
	// 修改数据
	w.age++
}

func (w *Worker) rest() {
	fmt.Println(w.name,"在休息")
	// 修改数据
	w.age++
}