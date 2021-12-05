/*
切片的理解
	切片存储的是元素的信息和指向元素的指针，实际上元素可能存放在不连续的地方
	类似于链表的数据结构，只是把他们的地址都集合在一起了
	明白了这点之后，很多操作都可以理解了

	但是扩容前后数据是会发生复制的，因为扩容前后切片的内存地址发生了改变

	总结来说：切片的这些性质都要因为切片是引用类型，不是值类型

1. 索引
	和数组索引一样
	注意：如果是通过make创建的切片，只有被默认初始化的元素才能被索引
		如果没有初始化元素，用索引是会报错的
		也就是实际长度是len,cap是需要增加元素的时候才会被启用的

2. 切片
	和数组的切片一样
	返回值是切片类型
	切片的切片返回值是切片类型
	存储：通过切片操作只是创建了一个存储这部分信息的切片，内部存储并不会复制数据
		也就是如果通过这部分切片修改，是会影响原来的切片的

3. 追加
	return = append(name,value)
	可以追加多个数据
	也可以把一个切片的内容追加到原来的切片后面，但是后面需要加三个点
		return = append(name,[]type{value}...)

4. 删除
	实际上没有删除操作，删除的效果是通过append构造出来的
	原理：要删除的数据把切片分成两块，把要删除的元素后面的切片追加到待删除元素的位置上
	但是实际上并没有真正在内存上删除这个元素，因为这样的删除操作实际上只是把这个元素的信息和指针不放在这个切片上
	如果把切片的结果赋值给一个新的变量，原理的变量还是可以访问到被删除的元素的

5. 插入
	也没有插入的操作，也是通过append构造出来的
	可以先把要插入数据的位置后面的数据先切片出来
	然后按照插入数据，最后再把前面切片出来数据插入回去

6. 切片的复制
	copy()函数
	把一个切片的数据复制到另外一个空间中，两个切片是独立的
	copy(复制到的切片，被复制的切片)

8. 切片的嵌套
	和数组的嵌套一样

7. 循环
总结
	无论是添加，删除还是插入，实际操作的都不是实际的数据，而是他们的"缩影"
	用一个"缩影"表示完整的数据，操作的时候，就不用操作实际的数据
	切片是引用类型
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 创建切片
	gender := make([]string,2,3)

	// 索引
	gender[0] = "male"
	gender[1] = "female"

	// 切片
	sex := gender[:2]
	fmt.Println(sex,reflect.TypeOf(sex))

	// 切片的存储
	fmt.Printf("%p\n",&gender[0])
	fmt.Printf("%p\n",&sex[0])

	// 修改会影响原来的切片
	sex[0] = "男"
	sex[1] = "女"
	fmt.Println(gender,sex)

	// 删除元素
	age := make([]int,7,10)
	// 切片赋值
	age = []int{16,18,19,22,24,21,30}
	fmt.Println(age)
	deleteIndex := 4
	age = append(age[:deleteIndex],age[deleteIndex + 1:]...)
	fmt.Println(age)

	// 切片的插入
	/*
	注意：切片操作的是数据的地址
	但是append数据，操作的是实际的数据
	如果不新建切片，append进来的数据会把原来数据覆盖掉
	即使你把原来数据的地址复制出来了，但是append进来的数据会替代原来数据的地址
	所以会出现两个append进来数据
	 */
	insetIndex := 3
	result := make([]int,0,len(age) + 1)
	result = append(result,age[:insetIndex]...)
	result = append(result,99)
	result = append(result,age[insetIndex:]...)
	fmt.Println(result)


	// 切片的复制
	newAge := make([]int,0,len(age))
	copy(age,newAge)
	fmt.Printf("%p\n",&age)
	fmt.Printf("%p\n",&newAge)

	// 循环
	for i := 0 ; i < len(age); i++ {
		fmt.Printf("%d,",age[i])
	}

	fmt.Println()
	for _,value := range age {
		fmt.Printf("%d,",value)
	}
}
