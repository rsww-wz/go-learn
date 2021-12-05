/*
概述
	Go语言不是面向对象语言（不直接支持面向对象的语法），但是可以实现面向对象的属性
	面向对象是一种编程思想，不应该局限于语法是否能够实现

面向对象继承性的性质
	如果两个类存在继承关系，其中一个是子类，另一个是父类，那么：

	1.子类可以直接访问父类的属性和方法
	2.子类可以新增自己的属性和方法
	3.子类可以重写父类的方法，就是将父类已有的方法重新实现

Go语言继承性的实现
	1.子类访问父类的属性和方法：
		结构体嵌套：子类对象可以访问父类对象属性
		方法绑定：子类对象可以调用父类方法
	2.子类可以新增自己的属性和方法：给子类绑定方法即可
	3.子类重写父类方法：鸭子类型(同名方法，不同接受者实现)

 */

package main

import "fmt"

func main() {
	// 创建Person对象
	p1 := Person{"王二狗",30}
	fmt.Println(p1.name,p1.age) // 父类对象访问父类属性
	p1.eat() // 父类对象访问父类方法

	// 创建Student对象
	std1 := Student{Person{"Ruby", 18},"一中"}
	fmt.Println(std1.name,std1.age) // 子类对象访问父类属性（提升字段）
	std1.eat()                      // 子类对象访问父类方法

	// 子类访问自己的新增方法
	std1.study()

	// 子类重写父类方法
	std1.eat()
}

// 定义一个“父类”
type Person struct {
	name string
	age int
}

// 定义一个“子类”
type Student struct {
	Person  // 匿名字段，模拟继承性
	school string
}

// 定义父类方法
func (p Person) eat() {
	fmt.Println("父类的方法")
}

// 新增子类方法
func (s Student) study() {
	fmt.Println("子类新增的方法")
}

// 重写父类方法
func (s Student) eat() {
	fmt.Println("子类重写的方法")
}