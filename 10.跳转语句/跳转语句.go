// continue语句
/*
	for循环中遇到continue，会停止当前循环，进入下一次循环
	如果有多层循环的话，continue只能停止一层循环
	continue语句还可以使用标签，可以跳出多重循环
*/

// break语句
/*
	for循环中遇到break语句，跳出当前循环
	如果有多重循环，也是只能跳出一层循环
	break语句也可以配合标签使用
*/

// 标签
/*
	break语句和continue语句都可以使用标签
	标签是对for循环使用的，就是for循环的前面设置一个标签，这个标签就代表的这个for循环
	当break和continue设置标签后，就可以直接停止或者跳出有标签的这层循环
	主要作用就是远距离停止或者跳出多重循环
*/

// goto语句
/*
	作用：跳转到指定位置，无条件跳转语句
	语法
		goto 标签
		指定位置打上相同的标签
	慎用这个语句，基本上用不到，因为他会破坏程序的结构，也会降低程序的可读性
*/

package main

import "fmt"

//func main() {
//	for i := 0;i < 100;i++{
//		if i < 66 {
//			continue
//		} else {
//			fmt.Println(i)
//		}
//	}
//}

// 使用标签
//func main() {
//	re:for i := 0; i < 5; i++{
//		for j := 0; i < 10 ;j++ {
//			fmt.Println(i,j)
//			if j == 5 {
//				continue re
//			}
//		}
//	}
//}

// 猜数字
func main() {
	flag := 66
	var input int
	for {
		fmt.Print("请输入你要猜的数字:")
		_,_ = fmt.Scan(&input)
		if input > flag {
			fmt.Println("大了")
		} else if input < flag {
			fmt.Println("小了")
		} else {
			fmt.Println("对了")
			break
		}
	}
}