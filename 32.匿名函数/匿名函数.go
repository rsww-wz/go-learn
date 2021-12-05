/*
匿名函数定义
	和普通的函数定义一样，只不过没有名字

执行匿名函数
	方式1：定义的时候就用一个变量去接收这个匿名函数，这个变量结代指的是这个匿名函数
	方式2：直接在匿名函数后面加括号，有参数的传参数，就可以执行，有返回值的用变量接收（自执行函数）

返回匿名函数
	返回值可以写一个匿名函数
 */

package main

import "fmt"

// 返回匿名函数
func f1(a,b string) func(sum string) {
	fmt.Println(a,b)
	return func (sum string) {
		fmt.Println(sum)
	}
}

func main() {

	// 定义接收函数的变量，注意类型
	var run func(string,string)

	// 匿名函数
	run = func (a,b string) {
		fmt.Println(a,b)
	}

	// 执行匿名函数
	run("hello","world")

	// 执行返回匿名函数的函数
	v1 := f1("hello","world")
	v1("Good morning")

	// 自执行函数
	func (a,b int) {
		sum := a + b
		fmt.Println(sum)
	}(34,56)
}
