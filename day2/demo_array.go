// // 数组  存放元素得容器，必须指定存放得元素得类型和容量

// package main

// func main(){
// 	var a1 [3]bool // 变量叫a1， 长度为3得数组，放的是bool值，数组得长度是数组类型得一部分
// 	// 初始化方式一
// 	a1 := [3]bool{true,true,true}
// 	// 方式二,自动推断有多少个数
// 	a2 := [...]int{1,2,3,4,5}
// 	// 方式三
// 	a3 := [5]int{1,2,} // 后三位默认是0
// 	a3 := [5]int{0:1,4:2,} // 根据索引初始化，其余默认是0

// 	// 数组得遍历
// 	// 1. 根据索引遍历
// 	for i:=0;i<len(a3);i++ {
// 		fmt.println(a3[i])

// 	}
// 	// for range 遍历
// 	// 多维数组, 数组里只能放相同类型得。

// 	// 数组是值类型，复制了一份另外得，改变另外得，不改变原来得值

// 	// 引用类型，相当于软链

// }

package main

import "fmt"

func main() {
	// 求 [1 2 3 4 5] 所有元素得和
	// a1 := [...]int{1, 2, 3, 4, 5}
	// sum := 0
	// for _, v := range a1 {
	// 	// fmt.Printf("%d", v)
	// 	sum = sum + v  // = 是加法运算， := 是又定义了一个变量

	// }
	// fmt.Print(sum)
	// [1 3 5 7 8] 中找出和为8 得两个元素得下标
	a1 := [...]int{1, 3, 5, 7, 8}
	// print(len(a1))
	for i := 0; i < len(a1); i++ {
		// fmt.Println(a1[i])
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%v--%v)\n", i, j)
			}
		}
	}
	// for i, v := range a1 {
	// 	for i1, v1 := range a1 {
	// 		if v+v1 == 8 {
	// 			fmt.Println(i, i1)
	// 		} else {
	// 			continue
	// 		}
	// 	}
	// }

}
