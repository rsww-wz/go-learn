/*
声明
	var name [size][size]type
	var name = [size][size]type{{value},{value},...}
初始化
	name[index] = [size]type{value}
 */
package main

import "fmt"

func main() {
	// 声明二维数组
	var ages [2][2]int
	fmt.Println(ages)

	// 声明并初始化二维数组
	age := [2][3]int{{2,4,6},{1,3,5}}
	fmt.Println(age)

	// 选择性初始化
	var data [2][3]int
	data[0] = [3]int{11,22,33}
	fmt.Println(data)

	// 遍历二维数组
	for _,item := range age {
		for _,j := range item {
			fmt.Printf("%d,",j)
		}
	}
}
