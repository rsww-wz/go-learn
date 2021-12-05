/*
异常和错误
	错误是指可能出现问题的地方出现了问题，意料之中
	异常是指不该出现问题的地方出现了问题，意料之外

Go语言错误处理机制
	Go语言没有其他语言的try，except语句
	而是引入了error这种数据类型作为处理异常的机制，所以error也是一种数据类型

	为什么不用try，except语句，因为把错误处理设置成控制语句，会造成程序结构混乱，降低程序可读性

Go语言错误处理机制缺点
	通过返回值的方式，来强迫调用者对错误进行处理，要么忽略，要么处理
	如果都处理的话，处理错误的代码会占很大一部分
	大多数第三包的函数都是需要返回错误信息的

错误处理方法
	如果一个函数或者方法返回一个错误，那么按照惯例，它必须是函数蒋方舟的最后一个值
	处理错误的惯用方法是将返回的错误与nil进行比较，nil值表示没有发生错误，而非nil值表示出现错误

创建错误数据
	errors包下有一个New函数,可以创建自己的错误类型：New("错误提示信息")
	fmt包下也有一个创建错误的函数：Errorf()，可以格式化错误信息

获取错误信息
	Go语言中，错误类型是一个接口，只有一个Error方法，返回值是字符串类型，但是这个函数的功能是返回错误信息的
	如果想要获取更多的错误信息需要:查看返回错误的结构体，然后断言错误类型，再做相应的处理
	通常包的函数都会有相应的错误结构体作为错误处理方法

建议
	永远不要忽略一个错误，忽视错误会导致麻烦
 */

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// errors包创建一个error数据
	err1 := errors.New("自己创建的错误信息")
	fmt.Println(err1)
	fmt.Printf("%T\n",err1)

	// fmt包创建一个error数据
	err2 := fmt.Errorf("错误提示信息：%d",100)
	fmt.Println(err2)
	fmt.Printf("%T\n",err2)

	err3 := checkAge(-15)
	err4 := checkAge(20)
	fmt.Println(err3,err4)

	// 获取错误信息
	getError()
	getMoreError()
}

// 返回错误类型的函数
func checkAge(age int) error {
	if age < 0{
		//return errors.New("年龄不合法")
		err := fmt.Errorf("你给定的年龄%d是不合法的",age)
		return err
	}
	fmt.Println("年龄是：",age)
	return nil
}

// 常见的错误处理方法
func getError() {
	file,err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
	}
}

// 获取更多错误信息
func getMoreError() {
	f,err := os.Open("test.text")
	if err != nil {
		fmt.Println(err)
		// 断言错误
		if ins, ok := err.(*os.PathError);ok{
			fmt.Println(ins.Op)
			fmt.Println(ins.Path)
			fmt.Println(ins.Err)
		}
		return
	}
	fmt.Println(f.Name())
}
