/*
注意：字符串不可变(和python一样)

1. 获取字符串的长度
	len(name)
	返回字符串的字节长度
2. 是否以xxx开头
	strings.HasPrefix(name,xxx)
	返回布尔值
3. 是否以xxx结尾
	strings.HasSuffix(name,xxx)
	返回布尔值
4. 是否包含xxx
	strings.Contains(name,xxx)
	返回布尔值
	注意：中文返回的都是false
5. 变大写
	strings.ToUpper(name)
	返回的结果是大写的，原来的字符串不变，因为字符串是不可变的
6. 变小写
	strings.ToLower(name)
	返回的结果是大写的，原来的字符串不变，因为字符串是不可变的
7. 去两边
	strings.TrimRight(name,xxx)      去除右边
	strings.TrimLeft(name,xxx)       去除左边
	strings.Trim(name,xxx)           去除两边
	无论去除的是什么，都需要指定，即使是空格
	返回值也需要用变量接收
8. 替换
	strings.Replace(name,被替换，替换成，次数)
	如果次数是-1,表示全部替换
9. 分割
	strings.Split(name,"分割字符")
	返回的是一个切片类型(列表)
10. 拼接
	方法1：用加号
	方法2：Join(切片类型,"连接符")
	方法3：创建一个bytes.Buffer对象，用里面的WriteString和String方法
	方法4(推荐)：创建一个strings.Builder对象， 类中的builder.WriteString()
11. string -> int
	strconv.Atoi(string)
12. int -> string
	strconv.Itoa(int)
13. 字节集合
	[]byte(string)
14.rune集合
	[]rune(string)
15. 字节集合和rune集合 -> string --- (类型转换)
	string(name)
*/

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	name1 := "Good morning"
	name2 := "早上好"

	// 获取字符串长度
	fmt.Println(len(name1))
	fmt.Println(len(name2))

	// 判断字符串是否以xxx开头
	fmt.Println(strings.HasPrefix(name1,"G"))
	fmt.Println(strings.HasPrefix(name2,"早上"))

	// 判断字符串是否以xxx结尾
	fmt.Println(strings.HasSuffix(name1,"ing"))
	fmt.Println(strings.HasSuffix(name2,"早上"))

	// 字符串是否包含xxx
	strings.Contains(name1,"mor")
	strings.Contains(name2,"早")

	// 变大小写
	result1 := strings.ToUpper(name1)
	result2 := strings.ToLower(name1)
	fmt.Println(result1)
	fmt.Println(result2)

	// 去两边
	result3 := strings.TrimRight(name1,"ing")
	result4 := strings.TrimLeft(name1,"G")
	result5 := strings.Trim(name1," ")
	fmt.Println(result3,result4,result5)

	// 替换
	str := "looking,taking,making,playing"
	res1 := strings.Replace(str,"ing","ed",1)
	res2 := strings.Replace(str,"ing","ed",3)
	res3 := strings.Replace(str,"ing","ed",-1)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// 分割
	res4 := strings.Split(str,",")
	fmt.Println(res4)

	// 拼接

	// 方法1
	str1 := "数据结构"
	str2 := "与算法"
	r1 := str1 + str2

	// 方法2
	dataList := []string{"数据结构","与算法"}
	r2 := strings.Join(dataList,"")

	// 方法3
	// 创建一个bytes.Buffer对象
	var buffer bytes.Buffer
	buffer.WriteString("数据结构")
	buffer.WriteString("与算法")
	r3 := buffer.String()

	// 方法4
	// 创建一个Strings.Builder对象
	var builder strings.Builder
	builder.WriteString("数据结构")
	builder.WriteString("与算法")
	r4 := builder.String()

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
}
