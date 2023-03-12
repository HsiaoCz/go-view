package main

// 选择排序
// 选择排序，每次选择最小的，或者最大的放在已排序集合的后面
// 9 5 2 7 12 4
// 比如按升序排列
// 第一扫描，首先将9标记为最小的，然后往后扫描，发现5比9小，将5标记为最小的
// 继续往后扫描，发现2比5小，将2标记为最小的，当第一躺扫描完毕
// 发现最小的是2，将2和9交换顺序
// 继续这种扫描，知道全部拍好序

func SelectionSort(a []int, n int) {
	if n <= 1 {
		return
	}

	// 开始便利
	for i := 0; i < n; i++ {
		// 查找最小的值
		// 先声明一个值，代表最小的值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}

		// 交换数据
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

func main() {
	a := []int{12, 3, 6, 4, 17, 25, 5}
	SelectionSort(a, len(a))
}
