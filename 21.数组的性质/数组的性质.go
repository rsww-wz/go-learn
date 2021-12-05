/*
1. 可变
	数组的元素可以被更改，但是数组的长度和类型都不可变
2. 拷贝
	变量赋值时，数组会重新拷贝一份复制过去
3. 长度
	返回数组的元素个数
4. 索引
	可以做左值和右值(访问和修改)
5. 切片
	和字符串一样
6. 循环
	可以手动循环，也可以用for range循环

 */
package main

import (
	"fmt"
)

func main() {
	// 拷贝:复制数组修改后不会相互影响
	name1 := [2]string{"Anna","Alex"}
	name2 := name1
	name1[0] = "xh"
	fmt.Println("name1:",name1)
	fmt.Println("name2:",name2)

	// 获取数组的长度
	var numbers = [...]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(len(numbers))

	// 数组切片
	fmt.Println(numbers[3:6])

	// 手动循环
	for i := 0;i < len(numbers);i++ {
		fmt.Printf("%d,",numbers[i])
	}

	// for range循环
	fmt.Println()
	for _,value := range numbers {
		fmt.Printf("%d,",value)
	}

}
