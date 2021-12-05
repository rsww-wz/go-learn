/*
思路
	自定义一个error接口，里面只有一个error方法，只有实现了这个error方法就可以调用这个接口


 */

package main

import (
	"fmt"
	"math"
)

func main() {
	radius := -3.0
	area,err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("圆形的面积是：",area)
}

// 定义结构体
type areaError struct {
	msg string
	radius float64
}

// 实现error接口，就是实现error方法
func (e *areaError) Error() string {
	return fmt.Sprintf("error:半径,%.2f,%s",e.radius,e.msg)
}

func circleArea(radius float64)(float64,error) {
	if radius < 0 {
		return 0,&areaError{"半径是非法的",radius}
	}
	return math.Pi * radius * radius,nil
}