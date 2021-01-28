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
// 注意事项：

// 函数外的每个语句都必须以关键字开始（var、const、func等）
// :=不能使用在函数外。
// _多用于占位，表示忽略值。

// 常量
// const (
//     pi = 3.1415
//     e = 2.7182
// )

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：示例中，常量n1、n2、n3的值都是100。
// const (
//     n1 = 100
//     n2
//     n3
// )

// iota

// 1.iota在const关键字出现时将被重置为0。
// 2.const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)

// 使用_ 跳过某些值
// const (
// 	n1 = iota //0
// 	n2        //1
// 	_
// 	n4 //3
// )

// func main() {
// 	fmt.Println(n1, n2, n4) //0 1 3
// }

// const (
// 	n1 = iota //0
// 	n2 = 100  //100
// 	n3 = iota //2
// 	n4        //3
// )
// const n5 = iota //0

// func main() {
// 	fmt.Println(n1, n2, n3, n4, n5)
// }

// 定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，
// 也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）

// const (
// 	_  = iota
// 	KB = 1 << (10 * iota)
// 	MB = 1 << (10 * iota)
// 	GB = 1 << (10 * iota)
// 	TB = 1 << (10 * iota)
// 	PB = 1 << (10 * iota)
// )

// func main() {
// 	fmt.Println(KB, MB, GB)
// }

// 多个iota定义在一行,每一行，iota 新增1
// const (
// 	a, b = iota + 1, iota + 2 //1,2
// 	c, d                      //2,3
// 	e, f                      //3,4
// )
