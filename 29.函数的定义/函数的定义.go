/*
定义函数：声明函数
调用函数：执行函数

关键字：func
函数名
	函数名命名需要使用驼峰命名法，且不能出现下划线
	golang中根据首字母的大小写来确定可以访问的权限。无论是方法名、常量、变量名还是结构体的名称
	如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用
	函数名首字母大写：可以被其他的包访问
	函数名首字母小写：只能在本包中使用
参数列表
	参数类型
	参数名 类型
返回值列表
	返回值类型
	返回值名 类型
注意
	参数列表无论有没有返回值都需要括号
	返回值列表如果没有参数不需要括号，如果有一个参数，也不不需要括号，只有多个参数才需要括号
	参数列表和返回值列表的定义都要两种形式
		如果有名字和类型，函数体和返回值可以直接使用名字
		如果只有类型，函数体和返回值必须定义名字
 */

package main

import "fmt"

// 定义最简单的函数
func f1(){
	fmt.Println("hello world")
}

// 定义有一个返回值
func f2() bool {
	fmt.Println("Good morning")
	return true
}

// 定义有多个参数
func f3(v1 int,v2 bool,v3 string){
	fmt.Println(v1,v2,v3)
}

// 定义有多个返回值
func f4(v1 int,v2 bool,v3 string) (int,bool,string){
	fmt.Println(v1,v2,v3)
	return v1,v2,v3
}

func main() {
	// 调用无参数无返回值函数
	f1()

	// 调用有返回值函数
	a1 := f2()
	fmt.Println(a1)

	// 调用有参数函数
	f3(3,false,"rs")

	// 调用有参数，有返回值的函数
	b1,b2,b3 := f4(4,true,"ok")
	fmt.Println(b1,b2,b3)
}