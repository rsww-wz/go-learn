// 常量
/*
单个声明
	声明常量和变量的方式一样，只是关键字是const，不是var
	常量声明可以使用标准声明，推导式声明和简短声明
	但是常量是不可以修改的，所以一般用推导式声明常量即可

批量声明
	和变量的批量声明一样，三种方法都可以用
	但是有一点不同的是，常量是不可以修改的，所以不能只是声明，声明的时候必须赋值

和变量的不同点
	不可修改
	声明时必须赋值
*/

// iota
/*
iota,可有可无，当做计数器即可
作用：在声明连续的变量的时候可以使用，iota默认是从0开始的
iota实际上就是一个常量
如果要跳过某个数，可以用下划线代替
如果中断了，需要显式恢复,恢复后的常量继续iota的值
*/

package main

import "fmt"

func main() {
	// 单个声明
	const Name = "小明"
	fmt.Println(Name)

	// 批量声明
	const v1, v2 = 78, 22
	fmt.Println(v1 + v2)
	fmt.Println(v1 * v2)
	fmt.Println(v1 / v2)
	fmt.Println(v1 % v2)

	// iota
	const (
		a1 = iota
		a2
		a3
		a4
		a5
	)
	fmt.Println(a1, a2, a3, a4, a5)

	const (
		b1 = iota + 2
		b2
		b3
		b4
		b5
	)
	fmt.Println(b1, b2, b3, b4, b5)

	// 空标识符
	const (
		c1 = iota + 1
		_
		c2
		c3
		_
		c4
		c5
	)
	fmt.Println(c1, c2, c3, c4, c5)

	// 中断
	const (
		d1 = iota * 10
		d2
		d3 = 200
		d4
		d5 = iota * 10
	)
	fmt.Println(d1, d2, d3, d4, d5)
}
