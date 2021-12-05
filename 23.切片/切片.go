/*

首先声明：切片就是python中的列表，但是类型是指定的，或者其他语言的动态数组

切片存储结构
	内存
		每个切片元素都存储的三个东西：指针(arr)，长度(len)，容量(cap)
	扩容
		在向切片中追加的数据大于容量时，内部会自动扩容，每次扩容且扩容都是当前容量的2倍
		当容量超过1024时，每次扩容只增加1/4容量

声明切片
	var name []type
		如果中括号有数字，声明的是数组，没有数字声明的是切片

	var name = []type{value}
	name := []type{value}
		声明并初始化，中括号有数字就是数组，没有数字就是切片

	var name = make([]type,len,cap)
		len:切片长度(默认初始化多个元素)
		cap:切片总容量(总共申请多少个元素，剩下的元素不会默认初始化)
	var name = make([]type,size)
		如果只给了一个数字size,则表示长度和容量都是size

声明切片指针
	var name = new([]type)
		指向默认值
	var name *[]type
		指向nil

自动扩容
	获取长度：len(name)
	获取容量：cpa(name)

	添加内容
		return := append(name,value)
		返回值也是一个切片，返回的是修改后的切片，原来切片不会发生改变
		但是这两个切片的内存地址是一样的
		如果不需要原来的切片，可以用原来的切片去接收返回值
make
	只用于创建切片，字典，channel
 */
package main

import "fmt"

func main() {
	// 声明：方式1
	var name []int
	fmt.Println(name)

	// 声明并初始化：方式2
	var name1 = []int{1,2,3}
	fmt.Println(name1)

	// 声明并初始化：方式3
	name2 := []int{4,5,6}
	fmt.Println(name2)

	// 声明：方式4
	var name3 = make([]int,3,5)
	fmt.Println(name3)

	// 获取切片长度和容量
	fmt.Println("name3:",len(name3),cap(name3))

	// 添加元素
	name4 := append(name3,9)
	fmt.Println(name4)

	// 自动扩容,超过容量后，len和cap会自动翻倍
	name3 = append(name3,43,45,47)
	fmt.Println("name3:",len(name3),cap(name3))
	fmt.Println(name3)

}
