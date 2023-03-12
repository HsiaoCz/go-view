package main

// 插入排序
// 1,4,3,12,6,9,24,15，
// 第一趟选择4，它前面没有值，所有放着不变
// 第二趟选择1，和4比较，1小，所以1放在前面
// 每次选择一个和前面的进行比较
// 有点像打扑克牌

func InsertSort(a []int, n int) {
	if n <= 1 {
		return
	}

	// 开始排序
	// 从索引为1开始，也就是第二个元素开始，和前面的进行比较
	for i := 0; i < n; i++ {
		value := a[i]
		j := i - 1
		// 查找要插入的位置并移动数据
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

func main() {
	a := []int{4, 1, 3, 6, 5, 9, 7}
	InsertSort(a, len(a))
}
