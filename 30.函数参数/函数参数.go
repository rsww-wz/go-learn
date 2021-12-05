/*
参数传递
	Go语言中参数传递的类型是：值传递(也就是把实参复制一遍在传给形参，相互不影响)

定义参数
	先定义名字，再定义类型,和C语言刚好相反
	如果有多个参数的类型是相同的，可以在这些参数的最后写一个类型即可

参数是引用类型时
	实际上还是会把实参复制一遍，但是里面的内容是地址
	所以函数内修改引用类型的数据时，还是会影响实际的数据(和C语言一样)

参数是指针时
	其他引用类型和指针类型的原理都是一样，因为他们实际上都是地址
	指针的好处是：如果不是引用类型的数据，也想直接操作数据，可以借助于指针实现
	类型前面加个星号即可

参数是函数
	把函数名和类型抄一遍
	可以理解为去掉func，和函数体的部分，有点类型C语言中的函数声明
	但是这样写，参数列表可能会写的很长，这是可以借助于type关键字
	type关键字和C语言中typedef的用法是一样的，作用是起一个别名
		Go：type 新名字 原来名字
	这是再把别名写入函数参数即可

变长参数
	语法：name ...type
	理解：可以传任意数量个这个类型的参数
	实现：name要可以接受不指定数量的参数，而且参数的类型要一致
		所以这个name的类型应该是切片类型，所以可以按照切片的方法处理这个变长参数
	注意事项
		一个函数只能出现一次变长参数
		如果有多个参数，变长参数应该放在最后
总结
	如果函数没有参数，也没有返回值，返回值的括号可以不写，但是参数的括号必须写
	参数可以定义在main函数的上面，也可以定义在main函数的下面，不需要声明函数
	返回多个值的时候，不需要用括号，用逗号分隔
	接收多个返回值和python一样，需要用多个变量接收
	参数是需要命名的
		使用参数的时候，不需要时间简短声明变量
		可变参数存放在切片中
	如果给返回值命
		只有一个返回值可以不用括号
		使用返回值的时候，不需要使用简短声明变量
		可以只写一个return语句，不用写返回变量

补充
	高阶函数：接受了一个函数作为参数的函数
	回调函数：作为另一个函数参数的函数，也就是说回调函数是一个参数
 */

package main

import "fmt"

// 参数简写,如果返回值有名称需要用括号
func add(a,b int) (sum int) {
	sum = a + b
	return sum
}

// 返回值只有类型，不需要括号
func sub(a int,b int) int {
	sum := a - b
	return sum
}

// 变长参数,需要指定类型,没有返回值可以省略括号
func myPrint(args ...string) {
	for _,item := range args {
		fmt.Print(item)
	}
}

// 参数类型是引用类型
func getBounderAge(age []int) []int{
	for index,value := range age {
		if value < 18 {
			age[index] = 0
		} else {
			age[index] = 1
		}
	}
	return age
}

func power(a,b int) int {
	mul := a
	for i := 1; i < b; i++ {
		mul *= a
	}
	return mul
}

// 函数做为参数
func f1(a,b int,function func(a,b int) int) int {
	// 执行function
	r1 := function(a,b)
	return r1
}

// 用type简化写法，user就是后面的类型
type user func(a,b int) int

// 重新改写
func f2(a,b int,function user) int {
	// 执行function
	r1 := function(a,b)
	return r1
}

func main() {
	// 调用函数
	var r1,r2 int

	r1 = add(89,23)
	fmt.Println(r1)

	r2 = sub(r1,r1-20)
	fmt.Println(r2)

	// 调用变长参数
	myPrint("hello","world","this","is","my","name")
	fmt.Println()

	var age = []int{12,15,24,5,26,56,18,20,17}
	fmt.Println(age)
	age = getBounderAge(age)
	fmt.Println(age)

	// 调用参数是函数的函数
	// 回调函数
	r1 = f1(34,56,add)
	r2 = f1(34,56,sub)
	r3 := f1(2,5,power)
	r4 := f2(3,3,power)
	fmt.Println(r1,r2,r3,r4)
}
