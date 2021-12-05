/*
Go语言的指针和C语言的指针最大的不同就是：Go语言的指针不能进行运算，不能发生指针偏移
由于指针不能偏移，指针的效果大打折扣，基本上只能在传递值类型参数的时候变成引用类型（内核就是通过指针修改数据）

当用指针访问数组的时候，可以省略星号，直接用指针名配合下标即可访问元素，如果用星号的话，需要用括号括起来
指针数组：是一个数组，存储的数据类型是指针，类似切片

判断是数组指针还是指针数组
	数组指针：中括号前面有星号
	指针数组：类型前面有星号
 */

package main

import "fmt"

func loop(list [5]int) {
	// 指针遍历数组
	pointer := &list
	for i := 0;i < len(list);i++{
		//fmt.Println((*pointer)[i])
		// 简化写法
		fmt.Println(pointer[i])
	}
}

func exchange(a,b *int){
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	age := [5]int{45,56,21,16,62}
	loop(age)

	var v1,v2 = 3,6
	fmt.Println(v1,v2)
	exchange(&v1,&v2)
	fmt.Println(v1,v2)
}
