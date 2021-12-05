// for循环的4种写法
/*
1. 形式1：死循环 (类似while循环)
	for { 代码 }

2. 形式2：(有点像python中的for循环)
	for 条件 { 代码 }
	每次循环都会判断条件

3. 形式3：(for循环的常见形式，没有括号)
	for 循环变量;条件;循环语句 {代码}

4. 形式4

*/

// 练习
/*
1.猜数字
	设定一个理想数字，比如66，一直提示让用户输入数字，如果比66大，则显示猜测的结果大了
	如果比66小，则显示猜测的结果小了，只有输入的等于66，显示猜测结果正确，，然后退出循环
2. 使用循环输出1-100所有整数
3. 使用循环输出10以内除了7以外的整数
4. 输出1-100内所有奇数
5. 输出1-100内所有偶数
6. 求1-100的所有整数的和
7. 输出10-1所有整数
*/
package main

import "fmt"

// 猜数字
//func main() {
//	flag := 66
//	var input int
//	for input != flag{
//		fmt.Print("请输入你要猜的数字:")
//		_,_ = fmt.Scan(&input)
//		if input > flag {
//			fmt.Println("大了")
//		} else if input < flag {
//			fmt.Println("小了")
//		} else {
//			fmt.Println("对了")
//		}
//	}
//}

// 输出1-100所有整数
//func main() {
//	for i := 1;i<=100;i++{
//		fmt.Println(i)
//	}
//}

//输出10以内除了7以外的整数
//func main() {
//	for i := 1;i <= 10 ;i++ {
//		if i != 7 {
//			fmt.Println(i)
//		}
//	}
//}

//输出1-100内所有奇数/偶数
//func main() {
//	for i := 1;i <= 100;i++{
//		if i % 2 == 1 {
//			fmt.Println(i)
//		}
//	}
//}

//求1-100的所有整数的和
//func main() {
//	var sum = 0
//	for i := 0;i <= 100;i++ {
//		sum = sum + i
//	}
//	fmt.Println(sum)
//}

//输出10-1所有整数
func main() {
	for i := 10;i > 0; i-- {
		fmt.Println(i)
	}
}
