/*
初始化也叫创建结构体对象
方式1：先后顺序
	如果有嵌套也是按照定义的顺序初始化
	匿名字段也是按住定义顺序

方式2：关键字方式(key，value的形式)
	有嵌套需要用结构体名指定类型
	如果是匿名字段，就是使用结构体名作为key(匿名字段默认使用结构体名作为字段)

方式3：先声明再赋值(使用访问成员赋值)
	结构体嵌套需要用句点访问到每层的结构体，直到访问到实际变量
	可以直接访问匿名字段成员


访问字段
	成员运算符:句点
 */

package main

import "fmt"

func main() {
	// 定义结构体
	type Animal struct {
		Type string
		Fly bool
	}

	// 匿名字段
	type Student struct {
		name string
		age int
		height float64
		grade map[string]int
		Animal
	}

	// 方式1：先后顺序
	std1 := Student{
		"小明",
		22,
		180,
		map[string]int{
			"语文":90,
			"数学":125,
			"英语":130,
		},
		Animal{"human",false},
	}

	// 方式2：关键字
	std2 := Student{
		name:"小刚",
		age:25,
		height: 170,
		grade: map[string]int{
			"语文":70,
			"数学":115,
			"英语":90,
		},
		Animal:Animal{Type: "human",Fly:false},
	}

	// 方式3：先声明再赋值
	var std3 Student
	std3.name = "小红"
	std3.age = 20
	std3.height = 160
	std3.grade = map[string]int{"语文":120,"数学":110,"英语":140}
	std3.Type = "human"
	std3.Fly = false

	fmt.Println(std1)
	fmt.Println(std2)
	fmt.Println(std3)
}
