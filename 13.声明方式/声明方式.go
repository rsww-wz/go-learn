//
/*
变量声明
	var v1 int
	v2 := 100

指针声明
	var v3 *int
	v4 := new(int)

new
	new用于创建内存并进行内部数据的初始化
	并返回一个指针类型

	官方描述：内建函数new用来分配内存，它的第一个参数是一个类型，不是一个值，它的返回值是一个指向新分配类型零值的指针
	作用：是初始化一个指向类型的指针(*T)

nil
	是Go语言中的空值
	但是并不是直接把空值赋值给变量，而是把nil的地址赋值给变量

建议
	尽量用new方式初始化并返回一个指针类型的方式
 */

package main

import "fmt"

func main() {
	v1 := new(int)
	v2 := new(int)
	fmt.Println(v1,&v1)
	fmt.Println(v2,&v2)
}
