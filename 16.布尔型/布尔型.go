/*
	表示真假，一般是和条件等配合使用，用于满足某个条件时，执行某个操作。

	字符串转换为布尔类型
	result,err := strconv.ParseBool(字符串)
		返回true：1,t,T,true,TRUE,True
		返回false：0,f,F,false,FALSE,False

	布尔类型转字符串
		result := strconv.FormatBool(bool)
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串转布尔类型
	v1,v2,v3,v4,v5,v6,v7,v8,v9,v10 := "1","0","t","T","F","f","True","true","false","False"
	r1,_ := strconv.ParseBool(v1)
	r2,_ := strconv.ParseBool(v2)
	r3,_ := strconv.ParseBool(v3)
	r4,_ := strconv.ParseBool(v4)
	r5,_ := strconv.ParseBool(v5)
	r6,_ := strconv.ParseBool(v6)
	r7,_ := strconv.ParseBool(v7)
	r8,_ := strconv.ParseBool(v8)
	r9,_ := strconv.ParseBool(v9)
	r10,_ := strconv.ParseBool(v10)
	fmt.Println(r1,r2,r3,r4,r5,r6,r7,r8,r9,r10)
	fmt.Printf("%T\n",r2)

	// 布尔类型转字符类
	result1 := strconv.FormatBool(true)
	result2 := strconv.FormatBool(false)
	fmt.Printf("%s,%T\n",result1,result2)
}
