package main

import (
	"fmt"
)

func main() {
	var arr = []int{7, 4, 6, 13, 5}
	fmt.Println("initial : ", arr)
	BubbleSort(arr)
	arr = []int{7, 4, 6, 13, 5}
	InsertSort2(arr)
	arr = []int{7, 4, 6, 13, 5, 20, 1, 2, 99, 8}
	quickSortA(arr, 0, len(arr)-1)
	fmt.Println(arr)
	return
	SelectSort(arr)
	InsertSort(arr)

	//arr2 := []int{-9, 9, 3, -159, 65, 24}
	//quickSort(arr2)
	//fmt.Println(arr2)
}

// BubbleSort 冒泡排序法
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

// 选择排序法，依次找到最小值按顺序排序
func SelectSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		var maxKey int = i
		for j := i + 1; j <= length-1; j++ {
			if arr[maxKey] < arr[j] {
				maxKey = j
			}
		}
		arr[i], arr[maxKey] = arr[maxKey], arr[i]
	}
	fmt.Println(arr)
}

func InsertSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		var insertIndex = i + 1
		var insertVal = arr[insertIndex]
		fmt.Println("insertIndex = ", insertIndex, "insert val = ", insertVal)
		if arr[insertIndex] > arr[insertIndex-1] {
			fmt.Println("i > i-1", arr[insertIndex], arr[insertIndex-1])
			continue
		} else {
			cur := insertIndex
			fmt.Println("i !> i-1", arr[insertIndex], arr[insertIndex-1])
			for cur >= 0 {
				fmt.Println("curr i = ", cur)
				if cur-1 >= 0 && insertVal < arr[cur-1] {
					fmt.Println("curr i < ", arr[cur], arr[cur-1])
					arr[cur] = arr[cur-1]
					cur--
				} else {
					arr[cur] = insertVal
					break
				}
				fmt.Println("current = ", arr)
			}
		}
		fmt.Println(arr)
	}
}

func InsertSort2(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

// 快速排序，不考虑空间利用率
func quickSortSimple(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	pivot := array[0] // divided from middle

	var low []int
	var pivots []int
	var high []int

	for _, v := range array {
		if v < pivot {
			low = append(low, v)
		} else if v == pivot {
			pivots = append(pivots, v)
		} else {
			high = append(high, v)
		}
	}

	var result []int
	result = append(result, quickSortSimple(low)...)
	result = append(result, pivots...)
	result = append(result, quickSortSimple(high)...)
	return result
}

/**
快速排序：分治法+递归实现
随意取一个值A，将比A大的放在A的右边，比A小的放在A的左边；然后在左边的值AA中再取一个值B，将AA中比B小的值放在B的左边，将比B大的值放在B的右边。以此类推
*/
func quickSort(arr []int, first, last int) {
	flag := first
	left := first
	right := last

	if first >= last {
		return
	}
	// 将大于arr[flag]的都放在右边，小于的，都放在左边
	for first < last {
		// 如果flag从左边开始，那么是必须先从有右边开始比较，也就是先在右边找比flag小的
		for first < last {
			if arr[last] >= arr[flag] {
				last--
				continue
			}
			// 交换数据
			arr[last], arr[flag] = arr[flag], arr[last]
			flag = last
			break
		}
		for first < last {
			if arr[first] <= arr[flag] {
				first++
				continue
			}
			arr[first], arr[flag] = arr[flag], arr[first]
			flag = first
			break
		}
	}

	quickSort(arr, left, flag-1)
	quickSort(arr, flag+1, right)
}

func quickSortA(arr []int, left, right int) {
	flag := left
	first := left + 1
	last := right

	for first <= last {
		for first <= last {
			if arr[flag] > arr[last] {
				arr[flag], arr[last] = arr[last], arr[flag]
				flag = last
				last--
				break
			}
			last--
		}
		for first <= last {
			if arr[flag] < arr[first] {
				arr[flag], arr[first] = arr[first], arr[flag]
				flag = first
				first++
				break
			}
			first++
		}
	}

	quickSort(arr, left, flag-1)
	quickSort(arr, flag+1, right)

}
