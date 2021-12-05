// 单行注释快捷键：Ctrl+？
// 自定义快捷键：Ctrl+W

package main

import "fmt"

func main() {
	// Print:默认不换行，可以通过转义字符换行
	fmt.Print("Good morning\n")

	// Println:默认换行
	fmt.Println("hello world")

	// 输出多个文本默认加空格显示
	fmt.Println("你好","我是","你的","小甜甜")

	// 格式化输出,不换行
	fmt.Printf("%s\n","这是格式化输出字符串")
	fmt.Printf("%d\n",89)

}
