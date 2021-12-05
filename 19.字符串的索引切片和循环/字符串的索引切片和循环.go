/*
索引、切片和循环并不是字符串所特有的，而是其他数据类型也有的
索引
	字符串的所有是按字节提取数据的
	语法和python的索引一样，就是结果与不一样,结果都是以10进制表示的（utf8的字符码），需要通过类型转换
	索引也是从0开始的，超过最大索引会报错
切片
	和python的切片语法也是一样的,但是也是以字节作为单位，在中文切片上支持不是很好
	如果对中文切片的话，开头中中间不是3的倍数，切出来的是乱码，不是中文字符
	返回的结果就是字符串了

手动循环
	也是以字节为单位的，如果遍历的是中文字符串支持不太好
	学到集合之后，可以用把字节数据添加进集合，然后类型转换即可遍历中文

for range循环
	语法：for index,item := range 对象 {}
	index是序号，从零开始计算
	item是每次遍历得到的内容，不过是utf8的字符码，不是字符本身
	如果不想要index，可以用匿名变量接收

	for range循环默认是以字符为单位，不是以字节为单位，但是结果还是utf8，需要类型转换
	所以可以遍历中文字符串

	python中也有类似的循环

在切片和索引的时候，如果有中英混合的字符串，是不好按照字节来提取的
最好的方法就是按照字符来提取数据
所以可以统一转化为rune集合，操作的时候就可以以字符为单位进行操作
索引切片和遍历都可以以字符操作了，就类似于python中的操作了，就是最好还需要类型转换
*/

package main

import "fmt"

func main() {

	// 索引
	str1 := "hello world"
	str2 := "你好，世界"
	span1 := str1[0]
	span2 := str2[0]
	fmt.Println(string(span1))
	fmt.Println(string(span2))

	// 切片
	span3 := str1[0:5]
	span4 := str2[9:15]
	fmt.Println(span3)
	fmt.Println(span4)

	// 遍历英文字符串
	for i := 0; i < len(str1);i++ {
		fmt.Print(string(str1[i]))
	}

	// 遍历中文字符串
	fmt.Println()
	for i := 0; i < len(str2);i++ {
		fmt.Print(string(str2[i]))
	}

	// for range循环
	for index,item := range str1 {
		fmt.Println(index,string(item))
	}

	// 遍历中文字符串
	for _,item := range str2 {
		fmt.Print(string(item))
	}

	// 遍历中英混合字符串,以字符为单位，不管是什么语言的字符串都可以遍历
	fmt.Println()
	str3 := "hello 你好，世界 world"

	for _,item := range str3 {
		fmt.Print(string(item))
	}

	// 转化为rune集合
	fmt.Println()
	data := []rune(str3)
	fmt.Println(string(data[0:8]))
	fmt.Println(string(data[7]))
	for i := 0;i < len(data);i++ {
		fmt.Print(string(data[i]))
	}

	fmt.Println()
	for _,item := range str3 {
		fmt.Print(string(item))
	}
}
