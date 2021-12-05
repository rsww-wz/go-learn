// Go语句数据类型概览
/*
整型：用于表示整数
浮点型：用于表示小数
布尔型：用于表示bool值
字符串：用于表示文本信息
数组：用于表示多个数据
指针：用于表示内存地址的类型
切片：用于表示多个数据
字典：用于表示键值对结合
结构体：用于自定义一些数据集合
接口：用于约苏和泛指数据类型

补充
查看变量类型
	reflect.TypeOf(变量)
	%T：也可以用格式化输出
*/

// 整型
/*

整型主要划分为两大部分
	有符号整型：int8,int16,int32,int64
	无符号整型:uint8,uint16,uint32,uint64
int：在64位系统上表示int64；在32位系统上表示int32
uint：在64位系统上表示int64；在32位系统上表示int32

整型之间的转换(类型转换)
	Go语言不会像C语言一样，会发生隐式类型转换
	在Go语言中，不同的类型就是不能相互运算，所以如果两种数据类型想要做运算，只能通过类型转化成同一种类型才能做运算
	类型转化的语法：类型名称(变量) --- 和python是一样的
	但是：高位转向地位的话，数据范围会形成一个圈，然后从最小值开始，直到把所有的数都落到这个范围中

整型和字符串直接的转换
	字符串类型 := strconv.Itoa(整型)
	但是参数只能是int类型，如果是其他类型，可以通过整型之间的类型转换

字符串转化为整型
	整型 := strconv.Atoi(字符串)
	返回两个值，还一个是err，如果字符串不是数字，就不能发生转换，则会返回错误信息
	注意：两个值都必须接受，如果不想要就用匿名变量接收

进制转换
	十进制：整型的方式存在
	其他进制：字符串的形式存在
	整型转换成其他进制
		函数：strconv.FormatInt(int64(变量),进制数)

	其他进制转化为整型
		函数：strconv.ParseInt(变量，数据当做什么进制处理，转换过程中的范围约束)
		转换过程中的范围约束：并不是转换后整型的长度，如果转换过程在指定范围内，则转换成功，否则报错
		返回值：转化后的变量和error，返回值永远是int64

常见的数学运算：需要用到math包

*/

package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {
	var v1 int8 = 8
	var v2 int16 = 18
	v3 := int16(v1) + v2
	fmt.Println(v3)

	// 查看变量类型
	fmt.Println(reflect.TypeOf(v1),reflect.TypeOf(v2),reflect.TypeOf(v3))
	fmt.Printf("%T,%T,%T\n",v1,v2,v3)

	// 整型转向字符串
	var v4 = 23
	result := strconv.Itoa(v4)
	fmt.Printf("%s,%T\n",result,result)

	// 字符串转向整型
	var age = "22"
	result1, _ := strconv.Atoi(age)
	fmt.Printf("%d,%T\n",result1,result1)

	// 整型转换成其他进制
	var v5 = 64
	r1 := strconv.FormatInt(int64(v5),2)
	r2 := strconv.FormatInt(int64(v5),8)
	r3 := strconv.FormatInt(int64(v5),16)
	fmt.Printf("%s,%T\n",r1,r1)
	fmt.Printf("%s,%T\n",r2,r2)
	fmt.Printf("%s,%T\n",r3,r3)

	//其他进制转化为整型
	var data = "0100"
	c1,_ := strconv.ParseInt(data,2,16)
	c2,_ := strconv.ParseInt(data,8,16)
	c3,_ := strconv.ParseInt(data,16,16)
	fmt.Printf("%d,%T\n",c1,c1)
	fmt.Printf("%d,%T\n",c2,c2)
	fmt.Printf("%d,%T\n",c3,c3)

	// 常见的数学运算
	fmt.Println(math.Abs(-23))                     // 去绝对值
	fmt.Println(math.Floor(3.14))                  // 向下取整
	fmt.Println(math.Ceil(3.14))                   // 向上取整
	fmt.Println(math.Round(3.14))                  // 就近取整
	fmt.Println(math.Round(3.1415 * 100) / 100)    // 保留小数点后两位
	fmt.Println(math.Mod(11,3))                    // 取余数
	fmt.Println(math.Pow(2,4))                     // 做乘方运算
	fmt.Println(math.Pow10(3))                     // 底数为10的乘方运算
	fmt.Println(math.Max(34,36))                   // 取较大值
	fmt.Println(math.Min(34,36))                   // 取较小值

}
