// make函数创建切片
s1 := make([]int, 5, 10) // 默认值5个0，长度5，容量10
// 切片就是一个框，框住一块连续得内存， 只能保存相同类型得变量

// 切片之间不能比较得， 切片唯一合法得比较操作是和nil 比较
// 一个nil 值得切片没有底层数组，就是没有开辟内存空间
// 但是长度和容量都是0得切变，不一定是nil，可能开辟了内存空间
// make 函数就是说开辟一个内存空间，开始干活了
s2 := make([]int, 0) // len(s2)=0,cap(s2)=0,s2!=nil
// 所以判断一个切片是否是空，要用len函数==0 来判断，不应该用s==nil 判断

// 切片得赋值拷贝
s3 := []int{1, 3, 5}
s4 := s3 // s3和s4都指向了底层数组，切片不保存值，值都再底层数组中保存得

// 切片得遍历
// 1.索引遍历
// 2. range循环

//vscode：ctrl shitt p  搜snippet，configure user snippet --> go --> $0 是光标停止得位置