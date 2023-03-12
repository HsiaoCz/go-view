package main

// 冒泡排序：
// 1,4,3,12,6,9,24,15
// 第一次比较会确定最后一个位置的数

func BubbleSort(a []int, n int) {
	// 如果传进来的数组长度<=1直接返回
	if n <= 1 {
		return
	}

	// 否则开始排序
	for i := 0; i <= n; i++ {
		// 提前退出的标志
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				// 此次冒泡有数据交换
				flag = true
			}
		}
		if !flag {
			break
		}

	}
}

func main() {
	a := []int{1, 4, 3, 12, 6, 9, 24, 15}
	BubbleSort(a, len(a))
}
