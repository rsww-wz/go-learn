// 字典简介
/*
在学习任何的编程语言时，一般都会有一种数据类型称为：字典(dict)或者映射(map),以键值对为元素的数据集合
这种类型最大的特点就是查找速度非常快，因为他的底层存储是基于哈希表存储的

这种结构之所以快，是因为根据key可以直接找到数据存放的位置；而其他的数据类型是需要从前到后去逐一比对，相对来说比较耗时
以上只是基本的存储模型，而各个编程语言中的字典都会在此基础上进行相应的修改和优化

map的特点
	键不能重复
	键必须可哈希（int,bool,float,string,array）
	无序
 */

// 字典声明和初始化
/*
方式1
	声明：var name map[key type]value type
		不能通过不存在的key添加value
		这种操作一般是给其他map赋值，用的整体操作
	默认初始化：name := map[key type]value type{}
	自定义初始化：name := map[key type]value type{key:value...}

方式2
	name := make(map[key type]value type)
	name := make(map[key type]value type,size)

方式3
	name := new(map[key type]value type)
		返回的是一个指针,实际上map指针
		不能通过不存在的key添加value，一般也是用作整体操作
		name = &name1
 */
package main

import "fmt"

func main() {
	// 声明
	var name map[string]int
	fmt.Println(name)

	// 默认初始化
	name1 := map[string]int{}
	fmt.Println(name1)

	//自定义初始化
	name2 := map[string]int{"rs":21}
	fmt.Println(name2)

	// 通过make创建
	name3 := make(map[string]int)
	fmt.Println(name3)

	// 指定长度
	name4 := make(map[string]int,5)
	fmt.Println(name4)

	// 通过new，返回指针
	name5 := new(map[string]int)
	fmt.Println(name5)
}
