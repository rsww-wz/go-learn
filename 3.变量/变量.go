// GO程序的语法和格式
/*
package main:是每个文件都必须具有的语句，定义文件所在的包
import：表示导入包，可以python类似

Go语言基本格式
	大括号不可以在一行的开头
	语句后面不需要分号或者冒号
	声明的变量没有使用也是报错的
	逻辑语句要放在函数内部，不能放在放在函数外面（函数外面的每个语句必须以关键字开头）
		例如：不能在函数外面放打印的语句，赋值语句等
		只能在函数外部在变量和函数的声明


Go语言有代码格式化，在编译的时候，会自动格式化代码

定义变量关键字：var
定义函数关键字：func

go的注释
	单行注释：//
	多行注释：和C语言一样
	注释不可以嵌套
*/

// 变量声明
/*
Go语言是静态语句，每个变量都必须用自己的类型，和C语言一样
而且Go语言的变量不可以强制类型转换，定义的时候是什么变量，就只能是什么变量了

同一个作用域不能重复声明同名一个变量

关键字：var
语法：var name T(类型)

标准声明
	var spam string
	var spam1 int

批量声明
	var(
		a string
		b bool
		c int
	)

变量的初始化
	声明变量的时候，系统会默认给变量赋值
		string类型：空字符串
		int类型：0
		float类型：0
		bool类型：false

声明变量的同时给变量初始化
	标准声明
		var name T = expression

	类型推导声明
		因为给变量赋值了，就可以根据值的类型推导出变量的类型
		所以可以不用指定变量类型
		如果指定了反而会有警告,表示可以优化该语句

	短变量声明
		语法：:=
		声明并初始化变量，编译器会自动推导出其类型

	匿名变量声明
		匿名变量用下划线表示
		匿名变量不占用空间，不会分配内存
		多用于循环变量

	声明变量可以不给变量赋值就可以使用

 */

// 变量的名称和作用域
/*
变量名支持utf-8，开头字母和其他语言一样，不是数字，符号等

变量的作用域也和其他语言一样，局部变量都是不能被上一级作用域访问，可以被同级和下级作用域访问

同一个作用域不能声明两个相同的变量,也无法声明两个名称相同，类型不同的变量
不同的作用域就可以声明两个相同的变量

全局变量的声明只能用标准声明和推导式声明，不能用简短声明，否则会报错
*/

// 变量的内存地址
/*
	每个变量的地址都不一样，和C语言一样，只要是声明了变量，变量的内存地址都是不一样的,值改变了内存地址也不会变
	如果把一个变量赋值给另一个变量也只是把值复制过去，内存地址是不一样的，这点和python不一样
	python是把一个值的地址赋值给变量，存的是值的地址，Go是为变量创建了一个内存地址存放值，因为类型是固定的，所以变量的大小可以确定下来

如果int，string，bool这三种类型的变量遇到复制，就会把值拷贝一份
*/

// 示例代码
/*
package main

import "fmt"

func main() {

	//标准声明
	var name string
	var age int
	var weight float32
	var ok bool
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(weight)
	fmt.Println(ok)
	fmt.Println('\n')

	//批量声明
	var(
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a,b,c,d)
	fmt.Println('\n')

	//声明的时候初始化
	//var name1 string = "hello"
	//var age1 int = 18
	//var name2,age2 = "world",28
	//fmt.Println(name1,age1)
	//fmt.Println(name2,age2)
	//fmt.Println('\n')

	//类型推导声明
	var name3 = "hello world"
	var age3 = 34
	fmt.Println(name3,age3)
	fmt.Println('\n')

	//短变量声明
	name4 := "rs"
	fmt.Println(name4)
	fmt.Println('\n')
}
*/

package main

import "fmt"

// 声明全局变量：标准声明和推导声明
var gender1 = "female"

func main() {
	// 变量的标准声明:var 变量 类型 = 值
	var name string  = "小明"
	fmt.Println(name)

	// 变量的推导式声明：因为根据值可以推导出变量的类型，所以可以省略类型声明
	var name1 = "小红"
	fmt.Println(name1)

	// 变量的简短声明:声明并初始化变量,不需要关键字var
	name2 := "小刚"
	fmt.Println(name2)

	// 匿名变量声明匿名变量用下划线表示，匿名变量不占用空间，不会分配内存,多用于循环变量
	var _i = 99
	fmt.Println(_i)

	// 方式1：批量声明
	var v1,v2,v3 int
	fmt.Println(v1,v2,v3)

	// 方式2：批量声明:只是声明，并没有初始化，Go语言中没有初始化变量会赋默认值
	// int 0; string ""; bool False
	var (
		a int
		b string
		c bool
	)
	fmt.Println(a,b,c)

	// 批量声明并初始化1
	var b1,b2,b3 = 233,true,"hello world"
	fmt.Println(b1,b2,b3)

	// 批量声明并初始化2
	var (
		name3 = "rs"
		age = 18
		gender = true
	)
	fmt.Println(name3,age,gender)

	// 批量声明并初始化3
	c1,c2,c3 := 88,false,"Good morning"
	fmt.Println(c1,c2,c3)

	// 变量的作用域
	var age1 = 22
	if true {
		var age1 = 30
		fmt.Println(age1)
	}
	fmt.Println(age1)

	// 访问全局变量
	fmt.Println(gender1)

	// 内存地址
	name4 := "小王"
	name5 := name4
	fmt.Println(name4,&name4)
	fmt.Println(name5,&name5)

	// 修改了变量的值，内存地址不会改变，要也不会影响别的变量
	name4 = "小刘"
	fmt.Println(name4,&name4)
	fmt.Println(name5,&name5)
}
