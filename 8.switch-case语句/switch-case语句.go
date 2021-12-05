/*
使用方法和C语言的switch-case语句一样

语法
	switch n {
	case1:{}
	case2:{}
	case3:{}
	case4:{}
	default:{}
}
*/
package main

import "fmt"

func main() {
	var num int
	fmt.Println("请输入 一个整数:")
	_,_ = fmt.Scan(&num)

	switch num {
	case 1: {fmt.Println("This is 1")}
	case 2: {fmt.Println("This is 2")}
	case 3: {fmt.Println("This is 3")}
	case 4: {fmt.Println("This is 4")}
	default: {fmt.Println("超过4")}
	}
}
