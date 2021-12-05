/*
第一步:创建一个超大整型对象
	方式1：var v1 big.Int
	方式2：v3 := new(big.Int)

	方式3：var v2 *big.Int
		一般用不到，要直接给它赋值的时候可以使用

第二步:写入数据
	直接写入整型数据：obj.SetInt64(value)
	通过字符串写入整型数据：obj.SetString(value,10)
		第二个参数表示以什么进制写入超大整型

推荐
	用指针的方式，即new来进行创建和初始化
	写入数据用字符串，因为直接写入整型数据是超过int64的范围的，会报错

第三部：运算
	返回的结果也是需要用超大整型存放
	加：Add
	减：Sub
	乘：Mul
	除：Div
	除余：DivMod（返回商和余数）

	运算的结果可以转换成字符串类型：result.String()
*/


package main

import (
	"fmt"
	"math/big"
	"reflect"
)

func main() {

	// 创建超大整型对象：方式1
	// 指向的是默认的内存地址，已经初始化了
	var v1 big.Int
	fmt.Println(v1)

	// 创建超大整型对象：方式2
	// 指向的是nil的内存地址，没有初始化
	v2 := new(big.Int)
	fmt.Println(v2)

	// 创建超大整型对象：方式3
	// 指向的是nil的内存地址，没有初始化
	var v3 *big.Int
	fmt.Println(v3,reflect.TypeOf(v3))

	// 直接写入整型
	v1.SetInt64(890342)
	fmt.Println(v1)

	// 通过字符串写入
	v2.SetString("123456789",10)
	fmt.Println(v2,reflect.TypeOf(v2))

	// 简写
	//v4 := new(big.Int)
	//v4.SetInt64(34209)
	v4 := big.NewInt(34209)

	// 运算
	result := new(big.Int)

	// 加
	result.Add(v2,v4)
	fmt.Println(result)

	// 减
	result.Sub(v2,v4)
	fmt.Println(result)

	// 乘
	result.Mul(v2,v4)
	fmt.Println(result)

	// 除
	minder := new(big.Int)
	result.DivMod(v2,v4,minder)
	fmt.Println(result,minder)

	// 转换成字符串类型
	fmt.Printf("%d,%T\n",result,result.String())

}
