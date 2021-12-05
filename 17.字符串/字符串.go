/*
字符串的本质
	Go语言中的字符串编码是utf-8
	所以本质是utf-8编码的序列

len
	len(变量)：获取一个变量的长度，单位是字节
	因为英文中一个字母就是一个字节，所以可以算出英文的实际长度
	如果要算出中文的长度，需要除以3，因为utf8中，一个汉字的长度是三个字节

字节集合
	字符串转换为字节集合
		byteSet := []byte(name)
	字节集合转换为字符串
		byteList := []byte{229,176,143,231,142,139}
		targetString := string(byteList)

集合rune
	字符串转换为rune码点
		unicode的码点，实际上就是uint32的别名
		将字符串转换为Unicode的码点，返回的结果是按10进制的形式表现的
			tempSet := []rune(name)
	rune码点转换为字符串
		runeList := []rune{23567,29579}
		targetName := string(runeList)

长度处理
	runeLength := utf8.RuneCountInString(name)
	返回的是实际的字符长度
	无论是中文还是英文都可以算出实际长度
 */


package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var name = "小王"

	// utf-8中，一个汉字是三个字节
	// 通过索引可以获取字符的一个字节，一个字节是8位
	fmt.Println(name[0],strconv.FormatInt(int64(name[0]),2))
	fmt.Println(name[1],strconv.FormatInt(int64(name[1]),2))
	fmt.Println(name[2],strconv.FormatInt(int64(name[2]),2))

	// 获取长度，单位是字节
	fmt.Println(len(name))

	// 字符串转换为字节集合
	// []表示一个集合，后面跟着的是集合的类型，小括号是存放的内容
	// byte的类型是一个字节，也就是uint8
	// 返回的结果是字符以一个字节为单位的编码，但是是以10进制的形式表示
	byteSet := []byte(name)
	fmt.Println(byteSet)

	// 字节集合转换为字符串
	byteList := []byte{229,176,143,231,142,139}
	targetString := string(byteList)
	fmt.Println(targetString)

	// rune集合
	tempSet := []rune(name)
	fmt.Println(tempSet)

	// rune转换为字符串
	runeList := []rune{23567,29579}
	targetName := string(runeList)
	fmt.Println(targetName)

	// 长度处理
	var name1 = "hello world"
	var name2 = "我的名字是：hello"
	runeLength := utf8.RuneCountInString(name)
	runeLength1 := utf8.RuneCountInString(name1)
	runeLength2 := utf8.RuneCountInString(name2)
	fmt.Println(runeLength)
	fmt.Println(runeLength1)
	fmt.Println(runeLength2)
	fmt.Println(len(name2))

}
