package main

import "fmt"

// if条件判断基本写法
// if 表达式1 {
//     分支1
// } else if 表达式2 {
//     分支2
// } else{
//     分支3
// }

// Go语言规定与if匹配的左括号{必须与if和表达式放在同一行，{放在其他位置会触发编译错误。 同理，与else匹配的{也必须与else写在同一行，else也必须与上一个if或else if右边的大括号在同一行。

func main() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// for 初始语句;条件表达式;结束语句{
//     循环体语句
// }
// for循环的初始语句和结束语句都可以省略，例如：
// func forDemo3() {
// 	i := 0	// for循环用完了， 这个变量还再。
// 	for i < 10 { // 没有分号了都
// 		fmt.Println(i)
// 		i++
// 	}
// }

// 无限循环
// for {
//     循环体语句
// }
