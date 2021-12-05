// 格式化输出
// Sprintf和printf都可以
// 格式化输出其实就是学习转移字符

package main

import "fmt"

func main() {

	// 格式化字符串输出
	var name,address,action string

	fmt.Print("请输入姓名:")
	fmt.Scan(&name)

	fmt.Print("请输入位置:")
	fmt.Scan(&address)

	fmt.Print("请输入行为:")
	fmt.Scan(&action)

	result := fmt.Sprintf("我叫%s,在%s正在%s",name,address,action)
	fmt.Println(result)
}
