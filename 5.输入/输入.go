// 输入
/*
fmt.Scan:回车无效
fmt.Scanln：回车有效
fmt.Scanf：格式化输入

语法
	都需要变量的内存地址
	如果输入多个值，通过空格分开
	都会返回两个值
		count：用户输入了几个值
		err：输入错误则是错误信息
	这两个值如果不用变量接收会提出警告
	如果不想要，可以用匿名变量接收,不需要用简短声明，直接等号即可
	但是如果想要一个，另一个不想要，就需要分号了，可以省略冒号的原则是都是匿名变量

特别说明
	fmt.Scan:如果要求输入多个，而没有输入够，会一直等待，中间可以按多次回车
	fmt.Scanln:如果要求输入多个，而没有输入够，只要按下回车，就立即停止输入，执行下面的程序
	fmt.Scanf:可以根据设定模板，在输入的时候提取指定信息(不推荐使用，因为在某些情况下提取数据会有问题)

问题
	上面三个函数都无法输入带有空格的文本,只会获取到第一个空格之前的文本
	如果需要输入带有空格的文本，需要用其他方法，需要用系统标准输入，导入bufio和os库

	reader := buffio.NewReader(os.stdin)
	line,isPrefix,err := reader.ReadLine()
	data := string(line)
	fmt.Println(data)

	这个方法只是一个读输入的函数，并让程序输入的函数，所以还是需要用输入函数先输入文本
	但是在输入英文文本的时候还是会损失第一个空格之前的内容
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 不换行输入
	var name1 string
	fmt.Print("请输入你的姓名>>>")
	count1,err1 := fmt.Scan(&name1)
	fmt.Println(name1,count1,err1)

	// 换行输入
	var name2 string
	var age int
	fmt.Print("请输入你的姓名>>>")
	count2,err2 := fmt.Scanln(&name2,&age)
	fmt.Println(name2,count2,err2)
	fmt.Println(name2,age)

	// 格式化输入
	var name3 string
	fmt.Print("请输入你的姓名>>>")
	// 用匿名变量接收
	_, _ = fmt.Scanf("%s",&name3)
	fmt.Println(name3)

	// 无法输入带有空格的文本
	var message string
	fmt.Print("请输入文本信息:")
	_, _ = fmt.Scan(&message)
	fmt.Println(message)

	// 输入带有空格的文本
	fmt.Print("请输入文本信息:")

	_, _ = fmt.Scan(&message)
	// 创建一个reader对象
	reader := bufio.NewReader(os.Stdin)

	// reader对象调用ReadLine方法，返回三个值
	//line,isPrefix,err := reader.ReadLine()
	//fmt.Println(isPrefix,err)
	// line:从stdin中读取一行数据，（它是字节，可以转换成字符串类型）
	// reader默认一次能读取4096个字节(4096/3个汉字)
	//	一次性能读完，isPrefix = false
	//	一次性不能读完，isPrefix = true

	line,_,_ := reader.ReadLine()

	// 类型转换
	data := string(line)

	fmt.Println(data)
}
