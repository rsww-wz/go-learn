/*
数组：定长，且元素类型一致的数据集合
	和C语言中的数据是类似的
	两个关键点：元素个数、类型必须确定

声明数组
	方法1:默认初始化
		var name [size]type
		已经默认初始化了，里面的值都是0，可以直接使用
	方法2：直接初始化（注意有等号）
		var name = [size]type{value}
	方式3：指定位置赋值
		var name = [size]type{0:value,1:value}
	方式4：省略个数
		var name = [...]type{value}
	方式5：省略关键字
		name := [size]type{value}

数组内存管理(和C语言是一模一样的)
	数组的内存是连续的
	数组的内存地址实际上就是数组第一个元素的内存地址
	取数组名的地址和取数组第一个元素的地址是相同的
	数组存字符串的时候存的是字符串的地址和长度，(原理上和C语言是一样的)

拓展
	声明指针数组（不会开辟内存空间，默认值为nil）
		var name *[size]type
	声明指针数组：返回的是指针类型的数组
		name := new([size]type)
 */
package main

import "fmt"

func main() {
	// 方式1：默认初始化
	// 声明时已在内存开辟空间，内存初始化的值都是0
	var numbers [3]int
	fmt.Println(numbers)

	// 方式2：声明时直接初始化
	var names = [2]string{"Anna","Alex"}
	fmt.Println(names)

	// 方式3：指定位置赋值
	var age = [2]int{0:18,1:20}
	fmt.Println(age)

	// 方式4：省略个数
	// 编译器自动推导个数
	var ages = [...]int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(ages)

	// 方式5：省略var
	gender := [3]string{"male","female","other"}
	fmt.Println(gender)

	// 数组内存地址
	fmt.Printf("数组的名字:%p\n",&ages)
	fmt.Printf("数组的第一个元素:%p\n",&ages[0])

	// 字符串数组地址
	// 无论字符串的长度是多少，他们的地址之间总是差16，说明他们实际上存的不是字符串，而是字符串的信息(地址，长度)
	fmt.Printf("第一个元素地址:%p\n",&gender[0])
	fmt.Printf("第二个元素地址:%p\n",&gender[1])
	fmt.Printf("第三个元素地址:%p\n",&gender[2])
}
