/*
返回值有return关键字

如果return语句后面没有任何变量，它的作用就是结束程序，return后面的语句不再执行

可以返回多个不同类型的变量

返回值可以是任意类型

返回值是函数时的声明
	语法：func name() func(args type)(return type){}
	也可以用type简化写法
 */

package main

import "fmt"

// return语句终止函数
func loop(a int) {
	for i := 0; i <= 100; i++ {
		if i > a {
			return
		} else {
			fmt.Printf("%d ",i)
		}
	}
}

//return返回函数
func f1() (function func(int)) {
	return loop
}

// 用type简化
type user func(int)

// 重写
func f2() (function user) {
	return loop
}

func main() {
	// 调用函数
	loop(30)
	loop(4)
	fmt.Println()

	// 返回值是函数的函数
	v1 := f1()
	v1(2)
	fmt.Println()

	v2 := f2()
	v2(4)
}
