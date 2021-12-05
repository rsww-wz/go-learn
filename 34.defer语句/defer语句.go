/*
defer语句：用于一个函数执行完成之后自动触发的语句，一般用于结束操作之后释放资源

defer语句位置
	defer语句是写在函数体内的
	可以写在函数体内部的任意位置，就是不能写在return语句的后面

defer语句执行顺序
	函数执行的时候会跳过defer语句，defer后面的语句会照样执行
	等函数执行完成之后(执行完return语句)再执行defer语句

多个defer语句的执行顺序
	defer语句到倒序执行，也就是从下往上执行defer语句

defer语句内容
	defer语句后面执行的是一条语句，当然也可以是调用函数，或者直接调用匿名函数

 */

package main

import "fmt"

// 单个defer语句
func f1() {
	defer fmt.Println("hello world")
	age := []int{3,4,6,7,8}
	fmt.Println(age)
}

// 多个defer语句
func f2() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

// defer语句是函数
func f3() {
	defer f1()
	fmt.Println("f3函数结束")
}

// defer语句放在return语句后面
func f4(a int) {
	if a > 20 {
		fmt.Println(a,"大于20")
	} else {
		return
	}
	defer fmt.Println("会不会执行呢？")
}

func main() {
	f1()
	f2()
	f3()
	fmt.Println()

	// 跳转return语句，能执行defer语句
	f4(12)

	// 没有跳转return语句，不能执行defer语句
	f4(22)
}
