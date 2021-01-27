package main

import "fmt"

//  在一个const声明里，第一次出现iota，是0 ， 新增一行变量声明，iota 加1
const (
	n1 = iota
	n2
	_
	n3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 1 变成 10000000000（二进制） 转换成10进制 是1024
	MB = 1 << (10 * iota)
)

func main() {
	fmt.Println("hello world")
	fmt.Println(n1, n2, n3)
}
