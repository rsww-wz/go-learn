/*
和C语言也是大同小异，语法上的不同，思想上是一样的
结构体就是自定义类型，和其他类型的地位是一样的

结构体
	结构体是一个符合类型，用于表示一组数据
	结构体由一系列属性组成，每个属性都有自己的类型和值

定义语法
	type 结构体名称 struct {
	字段 类型
	...
}

 */

// 匿名结构体和匿名字段
/*
匿名结构体
	没有名字的结构体
	在创建匿名结构体时，同时创建对象
	可以把匿名结构体赋值给一个变量，就相当于有名字了，基本没什么用
	变量名 := struct{字段 类型}

匿名字段
	一个结构体的字段没有字段名，只有类型
	其实默认是创建了一个和类型同名的变量作为字段

	注意：一个结构体不能有多个相同类型的匿名字段，因为相当于他们的字段名和类型都是一样的，否则会报错


简写
	如果有字段的类型是相同的，可以写在一起，类似函数参数
	类型可以是其他结构体
 */

package main

func main() {
	// 定义结构体
	type Animal struct {
		Type string
		Fly bool
	}

	// 匿名字段
	type Human struct {
		gender,name string
		Animal  // 匿名字段
	}

	// 结构体嵌套
	type Student struct {
		age    int
		height float64
		grade  map[string]int
		person Human
	}
}
