/*
作用
	1.定义类型
	2.给类型起别名

语法
	定义新类型
		type 类型名 Type
		不能类型不能相互赋值，即使是用同一个类型创造出来的类型，在语法上，他们是两种数据类型
	定义函数类型
		type 类型名 函数原型

	起别名
		type 别名 = 类型
		两种名字是通用的，不同的名字而已
		例如：byte是int8的别名；rune是int32的别名
			type byte int8
			type rune int32

		注意：非本地类型不能定义方法
			能够随意地为各种类型起名字，是否意味看可以在己包里为这些类型任意添加方法？ 否
			也就是引用别人的包时，不能给不是自己的数据类型起别名
			解决方法是用它的数据类型直接创建一个新的数据类型

 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 定义新类型
	type myInt int
	type myString string

	var v1 myInt
	var v2 myString
	var v3 int
	var v4 string

	v1 = 10
	v2 = "hello"
	v3 = 12
	v4 = "world"

	fmt.Println(v1,v2,v3,v4)

	res := f1()
	fmt.Println(res(34,56))
}

// 定义函数类型
type myFun func(int,int) string

func f1() myFun {
	fun := func(a,b int)string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}