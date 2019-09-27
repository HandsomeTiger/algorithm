package main

import (
	"fmt"
)

func main() {
	var arr = []int{7, 4, 6, 13, 5}
	fmt.Println("initial : ", arr)
	SelectSort(arr)
	fmt.Println("select sort:", arr)
	arr = []int{7, 4, 6, 13, 5}
	fmt.Println("initial : ", arr)
	InsertSort(arr)
	fmt.Println("insert sort:", arr)
}

// 选择排序法，依次找到最小值按顺序排序
func SelectSort(arr []int) {
	length := len(arr)
	fmt.Println(length)
	for i := 0; i < length-1; i++ {
		var maxKey int = i
		fmt.Println("----", i)
		for j := i + 1; j <= length-1; j++ {
			if arr[maxKey] < arr[j] {
				maxKey = j
			}
		}
		arr[i], arr[maxKey] = arr[maxKey], arr[i]
		fmt.Println(arr)
	}
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
					fmt.Println("curr i !< ", arr[cur], arr[cur-1])
					arr[cur] = insertVal
					break
				}
				fmt.Println("current = ", arr)
			}
		}
		fmt.Println(arr)
	}

}
