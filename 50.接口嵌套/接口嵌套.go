/*
嵌套就是继承，接口不但是允许继承，而且允许多继承

再次理解接口
 多态  接口A要求：只要是男的就可以进入队伍
	  接口B要求：只要是女的就可以进入队伍
	  接口C要求：如果你是男的，而且会说话就可以进入队伍
	  结构体D条件：男人，会说话，会洗衣服，会做饭

	  结构体D即可以选择进入接口A，也可以进入接口C，因为他满足了接口A，C的所有要求
	  这个男人还有其他技能，但是如果它进入了某个接口，就只能使用该接口的功能，自己的其他功能就无法发挥了

	  接口只是制定了要求，只要满足要求就能来，但是接口并没有实现具体的方法，要进来的人自己实现

	  这样理解的话结构体就是满足多态了，因为多态是针对子类的（结构体），
	      这个结构体可以选择进入不同的接口，每个接口都是一种形态，实现了同一个事物多种形态

	  结构体也满足鸭子类型了，你即使有其他功能，但是你不能使用，你现在完全满足接口，你就可以叫做接口的实现类

 嵌套  接口E：私营企业
		  接口F：腾讯公司
			  接口G：QQ音乐
			  接口H：微信
		  接口J：阿里
			  接口K：支付宝
			  接口L：淘宝

	  结构体M实现了接口K，那么结构体也能实现接口J和接口E，此时结构体实现那个接口，就拥有那个接口的属性和方法
 */
package main

import "fmt"

func main() {

	//Cat结构体实现C接口
	//C接口继承了A接口和B接口

	// 创建Cat对象
	var cat = Cat{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("---------------------")

	// 创建接口A对象
	var a1 A = cat
	a1.test1()      // 只有test1方法可以调用

	fmt.Println("---------------------")

	// 创建接口B对象
	var a2 B = cat
	a2.test2()     // 只有test2方法可以调用

	fmt.Println("---------------------")

	// 创建接口C对象
	var a3 C = cat
	a3.test1()
	a3.test2()
	a3.test3()

	/*
	实现类都是cat对象
	cat能实现什么方法，关键看给cat赋值那个接口

	如果cat要实现接口C，要把A，B,C的接口的方法都实现才能实现接口C
	 */
}
type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type Cat struct {

}

//实现了接口A
func (c Cat) test1() {
	fmt.Println("test1()")
}

//实现了接口B
func (c Cat) test2() {
	fmt.Println("test2()")
}

//实现了接口C
func (c Cat) test3() {
	fmt.Println("test3()")
}