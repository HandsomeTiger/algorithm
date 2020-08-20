package main

import (
	"fmt"
)

func main() {
	var arr = []int{7, 4, 6, 13, 5}
	fmt.Println("initial : ", arr)
	BubbleSort(arr)
	InsertSort2(arr)
	//return
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
			if arr[i] < arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

// 快速排序
func quickSort(arr []int) {

}
