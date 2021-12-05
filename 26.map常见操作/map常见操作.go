/*
1. 获取长度
	len(name):获取的是实际存放数据的容量
	注意：map没有容量，容量是底层自动扩容的，不可以访问容量时多少，不然报错

2. 增
	name[新键] = 新值

3. 删
	delete(name,key)

4. 改
	name[key] = value

5. 查
	r1,r2 := name[key]
	r1:是返回值，如果没有返回value的默认值
	r2：是否查到此值，如果没有返回false

6. 循环
	用for range循环比较合适
	如果只想要value，可以用匿名变量接收key的返回值

7. 赋值
	把一个map赋值给一个变量，这两个变量的地址是一样的，修改元素会相互影响
	因为map是引用类型，操作的只是数据的地址，原理和切片一样
	即使是扩容了，map的地址也是始终一样的
 */

package main

import "fmt"

func main() {
	data := map[string]int{"小红": 20, "小明": 18}
	fmt.Println(len(data))

	// 增
	data["小刚"] = 23
	fmt.Println(data)

	// 改
	data["小红"] = 16
	fmt.Println(data)

	// 删除
	delete(data, "小明")
	fmt.Println(data)

	// 查
	v1, v2 := data["小刚"]
	v3, v4 := data["小明"]
	fmt.Println(v1, v2)
	fmt.Println(v3, v4)

	// 循环
	for key, value := range data {
		fmt.Println(key, value)
	}
}
