package main

import "fmt"

// 有一点要注意的是
// for range采用的值覆盖的方式
// 值的地址不会发生改变
// 可以看到v的地址没有发生任何变化
func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4, 5)
	for _, v := range s {
		fmt.Println("v的值:", v)
		fmt.Printf("v的地址:%v\n", &v)
	}

}
