package main

import "math"

// 冒泡排序
// a是数组，n表示数组的大小
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}
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
		// 如果没有数据交换，提前退出
		if !flag {
			break
		}
	}
}

// 插入排序
func InsertionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
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

// 选择排序
func SelectionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		// 查找最小的值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		// 交换
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

// 归并排序
func MergeSort(a []int, n int) {
	if n <= 1 {
		return
	}
	mergeSort(a, 0, n-1)
}

func mergeSort(a []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)
	merge(a, start, mid, end)
}

func merge(a []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0

	for ; i <= mid && j <= end; k++ {
		if a[i] < a[j] {
			tmpArr[k] = a[i]
			i++
		} else {
			tmpArr[k] = a[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = a[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = a[j]
		j++
	}
	copy(a[start:end+1], tmpArr)
}

// 快速排序
func QuickSort(a []int, n int) {
	separateSort(a, 0, n-1)
}

func separateSort(a []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(a, start, end)
	separateSort(a, start, i-1)
	separateSort(a, i+1, end)
}

func partition(a []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := a[end]

	i := start
	for j := start; j < end; j++ {
		if a[j] < pivot {
			if !(i == j) {
				// 交换位置
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}
	a[i], a[end] = a[end], a[i]
	return i
}

// 计数排序

func CountingSort(a []int, n int) {
	if n <= 1 {
		return
	}

	var max = math.MinInt32
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}

	c := make([]int, max+1)
	for i := range a {
		c[a[i]]++
	}
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}

	r := make([]int, n)
	for i := range a {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}
	copy(a, r)
}

// 堆排序

func HeapSort(arr []int, n int) {
	// 1. 建立一个大顶堆
	buildMaxHeap(arr, n)
	length := n
	// 2. 交换堆顶元素与堆尾，并对剩余的元素重新建堆
	for i := n - 1; i > 0; i-- {
		swap(arr, 0, i)
		length--
		heapify(arr, 0, length)
	}
	// 3. 返回堆排序后的数组
}

func buildMaxHeap(arr []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, i, n)
	}
}

// 从上至下堆化
func heapify(arr []int, i int, n int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, n)
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
