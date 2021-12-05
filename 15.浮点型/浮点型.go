/*
Go语言提供了两种浮点型
	float32：用32位(4个字节)来储存浮点型
	float64：用64位(8个字节)来储存浮点型

	简短变量声明浮点型：默认是float64
	不同类型之间的浮点数也是不能运算的

浮点型的精度处理
	由于浮点型的特殊性质，不能精确的表示某些小数，任何语言都一样，这是计算机的实现原理上的问题
	所以在处理浮点数的精度问题上，需要一些函数

	decimal
	Go语言没有内置decimal包，需要自己下载
	安装第三方包
		终端上：go get 包地址
		源码会放在scr文件夹里面的GitHub文件夹里面
	调用第三方包
		import "github.com/shopspring/decimal"
 */

package main

import "fmt"

func main() {
	var v1 float32
	v1 = 3.14

	v2 := 9.99

	v3 := float64(v1) + v2
	fmt.Println(v3)
}
