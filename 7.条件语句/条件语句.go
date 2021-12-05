// 条件语句
/*
语句1：if语句
	if 条件{}

语句2：if-else语句
	if 条件{}
	else {}

语句3：级联if-else
	if 条件 {}
	else if 条件 {}
	else if 条件 {}
	else {}

	按顺序匹配条件，有限匹配上面的条件，如果都不匹配则执行else语句

语句4：嵌套if-else语句
	if 条件 {if 条件{} else {}}
	else 条件 {if 条件{} else {}}

	少用嵌套，因为可读性差，可以用switch-case语句代替

这个条件其实是bool值，可以写表达式和变量，因为表达式会返回一个bool值，变量也能返回一个bool值
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Print("请输入一个整数:")
	_,_ = fmt.Scan(&num)

	// 语句1
	if num > 20 {
		fmt.Println(num)
	}

	// 语句2
	// 判断奇偶数
	if num % 2 == 0 {
		fmt.Printf("%d是一个偶数\n",num)
	}else {
		fmt.Printf("%d是一个奇数\n",num)
	}

	// 语句3
	// 判断数字
	var length int
	fmt.Print("请输入你的长度:")
	_,_ = fmt.Scan(&length)
	if length < 5 {
		fmt.Println("太短了")
	} else if length < 10 {
		fmt.Println("刚刚好")
	}else {
		fmt.Println("满意")
	}

	// 语句4
	// 10086
	fmt.Println("欢迎致电10086")

	var number int
	_, _ = fmt.Scan(&number)

	if number == 1 {
		fmt.Println("话费服务，1.交话费；2.业务办理；3.查询")
		var n int
		_, _ = fmt.Scan(&n)
		if n == 1 {
			fmt.Println("交话费")
		} else if n == 2{
			fmt.Println("业务办理")
		} else {
			fmt.Println("查询")
		}
	}else {
		fmt.Println("输入错误")
	}
}
