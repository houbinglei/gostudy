package main

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
