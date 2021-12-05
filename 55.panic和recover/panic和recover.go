/*
Go语言中引入两个内置函数panic和recover来触发和终止异常处理流程，同时引入关键字defer来延迟执行defer后面的函数
直等到包含defer语句的函数执行完毕时，defer后的函数才会被执行
不管包含defer语句的函数是通过return的正常结束，还是由于panic导致的异常结束

panic和recover
	panic让程序报错，相当于其他语言的try语句
	recover让程序恢复，相当于其他语言的except语句
	panic和recover是处理异常的，属于流程控制

语法
	panic函数
		作用：让程序中断
		参数：任意类型，是一个空接口
		执行：不管panic被嵌套在第几层函数，都会回到main函数，已经被defer的函数都是要执行的
	recover函数
		作用：捕获panic，用来恢复程序
		参数：无
		返回值：panic的参数

Go语言的错误和异常
	总的来说就是error和panic，什么时候error，什么时候使用panic都是视情况而定
	以下场景适合异常（panic）
		空指针的引用
		下标越界
		除数为0
		不应该出现分支，比如default
		输入不应该引起函数错误
	其他场景使用错误处理
总结
	panic：让当前程序进入恐慌，中断程序的执行
	recover：让程序恢复，必须在defer函数中执行
	简单来讲∶go中可以抛岀一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理
 */
package main

func main() {

}
