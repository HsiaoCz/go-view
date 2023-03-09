package main

import "fmt"

func NewSlice(s []int) []int {
	s = append(s, 100)
	return s
}

func ChangeSlice(s []int) {
	s[1] = 12
}

func AppendSliceByPoint(s *[]int) []int {
	return append(*s, 100)
}

func AppendSliceByPointAndSendOldSlice(s *[]int) *[]int {
	*s = append(*s, 10)
	return s
}

func main() {
	s := []int{1, 2, 3, 4}
	newS := NewSlice(s)
	fmt.Println("原slice:s", s)       //原slice:s [1 2 3 4]
	fmt.Println("新slice:newS", newS) //新slice:newS [1 2 3 4 100]

	// 一个思考，这里为什么没有改变原来的slice呢？
	// 原因在于追加操作使得切片发生了扩容，而传入的参数是原参数的值拷贝而已
	// 这份值拷贝迁移到了新的内存地址

	// 再看一个例子
	ChangeSlice(s)
	fmt.Println("使用新的函数再看:s", s) // 使用新的函数再看:s [1 12 3 4]

	// 这里的值就发生了改变，为什么这里的值发生了改变呢？
	// 因为切片是引用类型，虽然这里传的是切片的值拷贝，但是它们指向的底层数组是一样的
	// 所以在函数里改变值，外边的slice也能看到

	// 再看一个例子
	newSS := AppendSliceByPoint(&s)
	fmt.Println("这一次呢:s", s)            // 这一次呢:s [1 12 3 4]
	fmt.Println("新的slice:NewSS", newSS) // 新的slice:NewSS [1 12 3 4 100]

	// 这里为什么值没有改变?
	// append会返回一个新的切片，这个切片并没有赋值给原来的切片
	// 这个切片赋值给了go自己帮我们起的名字的一个slice里面了

	// 这里一个新的例子
	newSSS := AppendSliceByPointAndSendOldSlice(&s)
	fmt.Println("这一次追加呢?,s:", s) //这一次追加呢?,s: [1 12 3 4 10]
	fmt.Println("新的slice:newSSS:", newSSS)
	// 这一次发生了改变，基于指针的操作
}
