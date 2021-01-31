// // slice 得append()方法
// // 调用append 函数 必须用原来得切片变量接收返回值，append 没有新创建一个数组，是修改了原来得数组。

// ss := []string{"北京"， "西安","广州"}
// s1 := []string{"武汉","长沙"}
// // s1 添加给ss
// ss = append(ss, s1...)  // ... 表示拆开，把s1 拆开

// // copy 复制切片，相当于重新开辟了一个内存空间，改原来得，不影响新copy 得切片

// // 切片删除， go 没有专门得方法删除切片得元素，只能自己写
// a1 = append(a1[:1],a1[2:]...)  // 删除了index 为2 得元素

package main

import (
	"fmt"
)

func main() {
	// var a = make([]int, 5, 10)
	// // fmt.Println(a)
	// for i := 0; i < 10; i++ {
	// 	// a := append(a, i)  这事错得，因为又重新定义了一个变量，应该用= 号，
	// 	a = append(a, i)

	// }
	// fmt.Println(a)
	// var a1 = [...]int{1, 6, 3, 5, 9}
	// sort.Ints(a1[:]) // 对切片进行排序
	// fmt.Println(a1)

	x1 := [...]int{1, 3, 5} //数组
	s1 := x1[:]             // 切片
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Printf("%p\n", &s1[0])
	// s1 = append(s1[:1], s1[2:]...)
	s1 = append(s1, 7)
	fmt.Printf("%p\n", &s1[0]) // append 超过容量之后，内存地址变了，没有超过，内存地址不变
}

// 0xc0000b0140
// 0xc0000da030
