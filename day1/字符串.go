// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
// }

// 多行字符串, 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。

// s1 := `第一行
// 第二行strings
// 第三行
// `
// fmt.Println(s1)

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "binglei"
	str := "c:\\Code\\lesson1\\go.exe"
	fmt.Println(len(s1))                      // 7
	ret := strings.Split(str, "\\")           // [c: Code lesson1 go.exe]
	fmt.Printf("%s+%s", s1, str)              //binglei+c:\Code\lesson1\go.exe
	fmt.Println(strings.Contains(s1, "bing")) // true
	fmt.Println(strings.HasPrefix(s1, "bi"))
	fmt.Println(strings.Join(ret, "+"))

}
