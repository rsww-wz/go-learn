/*
指针是一种数据类型，用于表示数据的内存地址
指针的类型和切片，map一样，都是引用类型
但是指针的作用是帮助值类型的数据类型实现引用类型的效果

声明不同数据类型的指针(整型,浮点型,字符串,数组)
	都是在类型前面加星号即可
	默认值都是指向nil

声明并初始化指针
	var name *p = &name
	name *p := &name

解引用
	语法：*地址

二级指针的声明
	var name **p = &name
	name **p := &name

 */

package main

import (
	"fmt"
	"reflect"
)

func point() {
	// 声明不同数据类型的指针,默认值都是指向nil
	var p1 *int
	var p2 *float64
	var p3 *string
	var p4 *[3]int
	fmt.Println(p1,p2,p3,p4)

	var v1 = 34
	var v2 = 3.14
	var v3 = "hello world"
	var v4 = [3]int{4,6,8}

	// 指针初始化
	p1 = &v1
	p2 = &v2
	p3 = &v3
	p4 = &v4
	fmt.Println(p1,p2,p3,p4)

	// 查看类型
	fmt.Println(reflect.TypeOf(p1))
	fmt.Println(reflect.TypeOf(p2))
	fmt.Println(reflect.TypeOf(p3))
	fmt.Println(reflect.TypeOf(p4))

	// 声明并初始化指针
	p5 := &v1
	var p6 = &v3
	fmt.Println(p5,p6)

	// 解引用
	fmt.Println(*p1,*p2,*p3,*p4)
}

// 二级指针
func pointPlus() {
	// 声明不同数据类型的指针,默认值都是指向nil
	var p1 **int
	var p2 **float64
	var p3 **string
	var p4 **[3]int
	fmt.Println(p1,p2,p3,p4)

	var v1 = 34
	var v2 = 3.14
	var v3 = "hello world"
	var v4 = [3]int{4,6,8}

	// 查看类型
	fmt.Println(reflect.TypeOf(p1))
	fmt.Println(reflect.TypeOf(p2))
	fmt.Println(reflect.TypeOf(p3))
	fmt.Println(reflect.TypeOf(p4))

	// 一级指针初始化
	b1 := &v1
	b2 := &v2
	b3 := &v3
	b4 := &v4

	// 二级指针初始化
	p1 = &b1
	p2 = &b2
	p3 = &b3
	p4 = &b4

	// 二级指针解引用
	fmt.Println(**p1)
	fmt.Println(**p2)
	fmt.Println(**p3)
	fmt.Println(**p4)
}

func main() {
	point()
	pointPlus()
}
