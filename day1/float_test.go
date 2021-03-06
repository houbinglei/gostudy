// Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准：
//  float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。
//   float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。

// 打印浮点数时，可以使用fmt包配合动词%f，代码如下：
// package main

// import (
// 	"fmt"
// 	"math"
// )

// go语言默认浮点型是float64
// func main() {
// 	fmt.Printf("%.10f\n", math.Pi)
// 	fmt.Printf("%.2f\n", math.Pi)
// }

// bool 类型
// Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

// 注意：

// 布尔类型变量的默认值为false。
// Go 语言中不允许将整型强制转换为布尔型.
// 布尔型无法参与数值运算，也无法与其他类型进行转换。
