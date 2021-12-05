/*
嵌套分为key的嵌套和value的嵌套，但是map要求key必须是可哈希类型
所以key的类型只能是，整型，浮点型，字符串，数组,即使key是数组，里面的元素也不能是可不可哈希类型(python可以)
也就是说key无论有没有嵌套，里面的类型都只能是可哈希类型
value里面还可以和其他类型嵌套，要会分析什么和什么嵌套

技巧：map和第一个中括号是key的类型，后面的都是value的类型
	map后面能有两个连续的中括号，说明key是数组

1. map和数组的嵌套
	value是数组 map[type][size]type
	key是数组   map[[size]type]type

2. map和切片的嵌套
	key不可以是切片，因为切片是不和哈希类型
	value是切片 map[type][]type

3. map和map嵌套
	key不可以是map
	value是map map[type]map[type]type

4. value中的嵌套

 */
package main

import "fmt"

func main() {
	//value是数组
	name := make(map[string][2]int)
	// 赋值
	name["小红"] = [2]int{89,75}
	name["小明"] = [2]int{92,78}
	for key,value := range name {
		fmt.Println(key,value)
	}

	// key是数组
	name1 := make(map[[2]int]string)
	// 赋值
	name1[[2]int{89,69}] = "小明"
	name1[[2]int{64,78}] = "小红"
	for key,value := range name1 {
		fmt.Println(key,value)
	}

	// map和切片的嵌套
	name2 := make(map[string][]int)
	// 赋值
	name2["小红"] = []int{67,60,79}
	name2["小明"] = []int{71,79,62}
	for key,value := range name2 {
		fmt.Println(key,value)
	}

	// map和map嵌套
	name3 := make(map[string]map[string]bool)
	// 赋值
	name3["小红"] = map[string]bool{"female":true}
	name3["小明"] = map[string]bool{"male":true}
	for key,value := range name3 {
		fmt.Println(key,value)
	}

	// value中的嵌套
	data := make(map[[2]int][3]map[string]int)
	// 赋值
	// key是一个长度为2的一维数组，元素类型是int
	// value是一个长度为3的一维数组，元素是map,map的类型是string:int
	data[[2]int{1,1}] = [3]map[string]int{{"小明":89},{"小红":97},{"小刚":78}}
	data[[2]int{2,2}] = [3]map[string]int{{"小爱":85},{"小布":70},{"小溪":56}}
	for key,value := range data {
		fmt.Println(key,value)
	}

	fmt.Println()
	for _,value := range data {
		for key,item := range value {
			fmt.Println(key,item)
		}
	}
}
