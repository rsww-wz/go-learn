//闭包理论
/*
Go语言中，无论是main函数，还是自己定义的函数，都不能直接实现高阶函数
但是可以通过匿名函数实现高阶函数，把匿名函数赋值给一个变量，可以实现高阶函数的功能

什么是闭包
	在函数外部定义，在函数内部使用的自由变量
	脱离了形成闭包的上下文，闭包也能照常使用这些自由变量(内部函数结束，依然可以访问这些变量)

闭包的概念
	承载闭包变量的函数称为外层函数
	调用闭包变量的函数称为闭包函数

	闭包基本思想都是借助于高阶函数实现

闭包的理解
	闭包是对全局变量的优化
	按照闭包的定义
		函数外定义，函数内使用，全局变量满足
		变量离开函数，依然可以访问到变量，全局变量也满足
	但是，无论是什么语句，都不推荐使用过多的全局变量，因为它会污染全局变量的环境，造成难以寻找的bug
	所以，闭包为了满足上面的情况，说采用的一种语法

	在python中，闭包很好理解，因为python可以直接定义高阶函数，第一层函数定义的变量就称为闭包
	它可以被第二层或者更内层的函数调用，并且离开调内层函数也能访问到变量

	闭包和面向对象一样，都是一种思想，不依赖于语法而存在

闭包的体现
	只要不离开最外层函数，内层函数操作闭包就好像操作全局变量一样
	但是，在Go语言中，实际上底层操作的并不是同一个变量，底层做了处理，让你在实现上好像操作的是同一个变量

闭包总结
	闭包 = 函数 + 外层变量的引用
	函数作为返回值
		在函数内再定义一个函数，返回函数内的定义的函数
		和python的高阶函数是一样的
	在两个函数间定义的变量，即第一层函数内定义的变量
	闭包：用一个变量接受返回值，此时这个变量成了一个函数，是内层函数，它就叫做闭包
	闭包结构：内层函数 + 外层函数的局部变量(参数或者自己定义的变量)
	功能：闭包可以接受最外层函数的参数
	判断是不是闭包：看返回函数内部有没有用到不是自己函数内部定义的变量

 */

//闭包实现
/*
Go语言中闭包的实现
	Go语言中不支持函数内部在定义一个函数
	但是支持函数内部定义匿名函数，所以可以借助于匿名函数实现闭包

实现闭包的步骤
	1.外层函数的返回值需要返回一个函数类型，这个函数类型要和闭包的函数类型一样
	2.外层函数的局部变量包括：自己定义的变量 ， 参数列表传过来的变量
	3.匿名函数需要访问外层函数的局部变量
	4.匿名函数需要用一个变量接收
	4.外层函数需要返回接收匿名函数的变量

闭包变量的声明周期
	表象：随着外层函数的结束，它也应该结束，但是匿名函数要访问他，所以会一直保持着声明周期
	实际
		函数的入口地址是一个二级指针，也就是函数名指向一个地址，这个地址才指向实际的函数数据
		这样的设计就是为了实现闭包，中间这一级指针，存放的就是闭包变量
		这个闭包变量是从外部函数的局部变量复制过来的，终止的结果就是好像操作的是外部函数的局部变量
 */

// 变量作用域
/*
局部变量
	函数内定义的变量
	全局作用域不能访问局部变量

全局变量
	在函数外面定义的变量
	在函数内可以访问全局变量
	如果全局变量和局部变量重名，会先找局部变量，没有就找全局变量
	也就是重名的话，局部变量会覆盖全局变量
	全局变量不能用简短声明，因为main函数外都必须是以关键字开头

语句块内定义
	for循环和switch语句内定义的变量
	只在语句块内存在

 */

package main

import "fmt"

// 闭包
func f1(a,b int) func()int {
	sum := a + b
	f2 := func () int{
		sum++
		return sum
	}
	return f2
}

func main() {
	// 调用闭包
	f3 := f1(3,6)
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())

}