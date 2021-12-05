/*
什么是接口
	面向对象世界中的接口的一般定义是“接口定义对象的行为”
	表示让指定对象应该做方法，实现这种行为的方法是针对对象的

Go语言接口
	Go语言中接口是一组方法的签名
	当类型为接口中的所有方法提供定义时，它被称为实现接口。
	它与OOP非常相似。接口指定了类型应该具有的方法，类型决定类如何实现这些方法

	它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是现实了这个接口
	接口定义了一组方法，如果某个对象实现了某个接口的所有方法，则，此对象就是先实现了该接口

理解
	接口相当于面向对象的实现类，这个类只定义方法，不写方法的具体实现
	接口也是只写方法，不写具体实现方法

	如果定义了某个结构体，让他去实现接口中的方法，结构体就是接口的实现

语法
	type 接口名字 interface { 方法名称(parameter list)(return list)	}

总结
	1.当需要接口类型的对象时，可以使用任意实现类对象代替
	2.接口对象不能访问实现类中属性
*/

package main

import "fmt"

func main() {
	// 创建Mouse对象
	m1 := Mouse{"罗技"}
	fmt.Println(m1.name)

	f1 := FlaskDisk{"金士顿"}
	fmt.Println(f1.name)

	// 实现类对象
	testInterface(m1)
	testInterface(f1)

	// 接口对象
	var usb USB
	usb = m1
	usb.start()
	usb.end()
}

// USB 定义接口
type USB interface {
	start()
	end()
}

// Mouse 实现类
type Mouse struct {
	name string
}

type FlaskDisk struct {
	name string
}

// 结构体绑定方法
func (m Mouse) start() {
	fmt.Println(m.name, "鼠标开始工作")
}

func (m Mouse) end() {
	fmt.Println(m.name, "鼠标结束工作")
}

func (f FlaskDisk) start() {
	fmt.Println(f.name, "U盘开始工作了")
}

func (f FlaskDisk) end() {
	fmt.Println(f.name, "U盘结束工作了")
}

// 定义父类 ： 开放接口
func testInterface(usb USB) {
	usb.start()
	usb.end()
}
