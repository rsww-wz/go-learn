/*
断言的含义
	空接口没有定义任何方法，因为Go中所有类型都实现了空接口
	当一个方法的形参是空接口时，那么函数中，需要对形参进行断言，从而得到它的真实类型

	也是上就是判断出它是什么类型的参数，也就是类型检查

语法格式
	方法1
	安全类型断言
		<目标类型的值>,<布尔参数> := <表达式>,(目标类型)
		断言成功：返回true，否则false
	非安全类型断言
		<目标类型的值> := <表达式>,(目标类型)
		如果不是会panic，程序直接终断
	方法2
	switch-case语句
		switch instance := 接口对象.(type){}

	推荐使用switch语言断言
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	// 创建一个三角形对象
	var t1 = Triangle{3,4,5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())

	// 创建一个圆形对象
	var c1 = Circle{3}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())

	// 创建一个接口对象
	var s1 Shape
	// 把三角形队形赋值给接口对象
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	// 断言函数
	getType(t1)
	getType(c1)

	getType2(t1)
	getType2(c1)
}

// 定义接口
type Shape interface {
	peri() float64  // 形状的周长
	area() float64  // 形状的面积
}

// if-else语句断言
func getType(s Shape) {
	if ins,ok := s.(Triangle) ; ok {
		fmt.Println("是三角形,三边是:",ins.a,ins.b,ins.c)
	} else if ins,ok := s.(Circle) ; ok {
		fmt.Println("是圆形，圆形半径是:",ins.radius)
	}else {
		fmt.Println("未知类型")
	}
}

// switch-case语句断言
func getType2(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形",ins.a,ins.b,ins.c)
	case Circle:
		fmt.Println("圆形",ins.radius)
	default:
		fmt.Println("位置类型")
	}
}

// 定义实现类
type Triangle struct {
	a,b,c float64
}

// 绑定方法
func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func(t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p-t.a) * (p-t.b) * (p-t.c))
	return s
}

type Circle struct {
	radius float64
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius,2) * math.Pi
}