// 结构体嵌套
/*
一个结构体可能包含一个，而这个字段反过来就是一个结构体，这种结构就称为嵌套结构

结构体嵌套使用结构体
	不会影响嵌套结构体的数据，但是会占用多一份内存
结构体嵌套使用结构体指针
	使用的是同一份数据和内存
 */

// 提升字段
/*
在结构体中属于匿名结构体的字段称为提升字段，因为他们可以被访问，就好像他们属于用于匿名字段结构体的结构一样

定义层面：匿名字段
使用层面：不需要上层结构体，直接访问属性
 */

	// 总结
	/*
	模拟继承性 ：is - a
		type A struct { 字段 类型 }
		type B struct { A // 匿名字段 }
	模拟聚合关系  has - a
		type C struct { 字段 类型}
		type D struct { c C // 聚合关系 }
	 */

package main

import "fmt"

func main() {
	b1 := Book{}
	b1.bookName = "西游记"
	b1.price = 45.7

	s1 := Student{}
	s1.name = "王二狗"
	s1.age = 18
	s1.book = b1

	fmt.Println(b1)
	fmt.Println(s1)

	fmt.Printf("%s,%d,%s,%.2f\n",s1.name,s1.age,s1.book.bookName,s1.book.price)

	// 修改数据，中间用的结构体复制，所以不会影响原来的数据
	b1.bookName = "红楼梦"
	fmt.Println(b1.bookName)
	fmt.Println(s1.book.bookName)

	b2 := Book{"三国演义",78.5}
	s2 := Student1{"李小花",26,&b2}
	fmt.Println(b2)
	fmt.Println(s2)

	// 用结构体指针传递数据
	b2.bookName = "水浒传"
	fmt.Println(b2.bookName)
	fmt.Println(s2.book.bookName)

	// 提升字段
	b3 := Book{"小王子",20.6}
	s3 := Student2{"王铁柱",34,b3}

	// 直接访问数据
	fmt.Println(b3.bookName)
	fmt.Println(s3.bookName)

}

type Book struct {
	bookName string
	price float64
}

type Student struct {
	name string
	age  int
	book Book  // 结构体嵌套
}

type Student1 struct {
	name string
	age  int
	book *Book  // 嵌套结构体指针
}

type Student2 struct {
	name string
	age  int
	Book    // 提升字段
}
