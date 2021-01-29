//字符用单引号（’）包裹起来

package main

import "fmt"

func main() {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // ctrl + shift +u 变成大写字母， A 是65
	// fmt.Println(len(s))               // 11
	for i := 0; i < len(s); i++ { // , byte
		// fmt.Printf("%v(%c)", s[i], s[i])
		fmt.Printf("%v\n", s[i]) // %c 是英文字符，
	}

	// for _, r := range s { //rune
	// 	fmt.Printf("%v(%c) ", r, r)
	// }
	// fmt.Println()

	// for i := 65; i <= 90; i++ {
	// 	fmt.Println(i)
	// }

}
