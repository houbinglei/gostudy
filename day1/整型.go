// 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
// uint	32位操作系统上就是uint32，64位操作系统上就是uint64
// int	32位操作系统上就是int32，64位操作系统上就是int64
// uintptr	无符号整型，用于存放一个指针
// package main

// import "fmt"

// var a = '3'

// func main() {
// 	fmt.Printf("%d\n", a)  // 字符是十进制类型得数返回。 51
// 	fmt.Print(a) 			// 51

// }
package main

import "fmt"

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF
}
