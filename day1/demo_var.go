package main

// 标准声明
// var 变量名 变量类型
// var name string
// var age int
// var isok bool

// 批量声明
// var（
// 	a string
// 	b int
// 	c bool
// ）

// 变量得初始化
// var name1 string = "binglei.hou"

//类型推导,这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。
// var name2 = "binglei.hou"
// var name3, age = "Q1mi", 20

// 短变量声明
// var m = 100
// func main() {
// 	n := 100
// 	m := 200
// 	fmt.Println(n, m)
// }

// 匿名变量_
// 注意事项：有些变量要接收，但是又不使用得时候，就可以用_来接收这个值。

// 函数外的每个语句都必须以关键字开始（var、const、func等）
// :=不能使用在函数外。
// _多用于占位，表示忽略值。

// for i := range s {
// 	fmt.println(i) // i 只返回索引
// }

// for i, v := range s{
// 	fmt,println(i,v) // 返回索引 和值得asc码
// 	fmt.printf("%d,%c\n",i,v) // 索引是int 类型， 值是字符类型， %c
// }

// for _, v := range s{
// 	fmt.printf("%c",v) // 只返回值
// }

import "fmt"

func main() {
	s := "侯冰雷"
	// 返回中文得字符得原型值
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}
}
