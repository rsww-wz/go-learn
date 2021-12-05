// 运算符
/*
算术运算符
	+,-,*,/,%

	++,--:不是运算符，也不是表达式，而是单独的一条语句，不能和其他语句组合，只能单独成句
	所以即使是最简单的自增自减，也要单独写一行，因为它不是运算符

关系运算符
	和其他语句一样,共有6个关系运算符：<,<=,>,>=,==,!=
	结果返回一个bool值

逻辑运算符
	和C语言一样，有三个运算符，分别是：&&，||，！
	&&：逻辑与，python中的and
	||：逻辑或，python中的or
	！：逻辑否，python中的not

	作用：用于连接两个返回值为bool值的语句
	有短路现象，返回的也是bool值

赋值运算符
	复合赋值，和其他语言一样

位运算符
	&: 有0得0，全1为1
	|: 有1得1，全0为0
	^: 相同得0，不同得1
	>>: 右移
	<<：左移

其他运算符
	sizeof：求变量类型的大小
	&：取地址运算符：返回变量的地址
	*：解引用运算符：指向变量的指针

*/

// 运算符的优先级
/*
	总的来说和C语言运算符的优先级是一样的、
	如果混合运算中有不确定的关系，可以通过加括号提高优先级
*/
package main

import "fmt"

func main() {
	var v1, v2 = 63, 23

	// 运算运算符
	fmt.Println(v1 + v2)
	fmt.Println(v1 - v2)
	fmt.Println(v1 * v2)
	fmt.Println(v1 / v2)
	fmt.Println(v1 % v2)
	v1++
	fmt.Println(v1)
	v2--
	fmt.Println(v2)

	// 关系运算符
	fmt.Println(v1 > v2)
	fmt.Println(v1 >= v2)
	fmt.Println(v1 < v2)
	fmt.Println(v1 <= v2)
	fmt.Println(v1 == v2)
	fmt.Println(v1 != v2)

	// 逻辑运算符
	//有false就false
	fmt.Println("这是&&逻辑运算符")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	// 有true就true
	fmt.Println("这是||逻辑运算符")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	// 逻辑相反
	fmt.Println("这是！逻辑运算符")
	fmt.Println(!true)
	fmt.Println(!false)

	// 赋值运算符
	v1 += 2
	fmt.Println(v1)
	v1 *= 2
	fmt.Println(v1)

	// 位运算符
	flag1 := 2 // 0010
	flag2 := 4 // 0100

	fmt.Println(flag1 & flag2) // 有0得0，全1为1
	fmt.Println(flag1 | flag2) // 有1得1，全0为0
	fmt.Println(flag1 ^ flag2) // 相同得0，不同得1
	fmt.Println(flag1 >> 2)    // 右移
	fmt.Println(flag2 << 2)    // 左移 10000
}
