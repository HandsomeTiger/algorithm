package main

import "fmt"

func main() {
	//arr := []float64{1, 3, 4, 67, 78, 99}
	//fmt.Println(find(98, arr))
	arr2 := []int{1, 2, 3, 4, 67, 78, 99}
	fmt.Println(BinaryFind(2, arr2, 0, len(arr2)-1))
}

// BinaryFind 二分查找法
func BinaryFind(targetVal int, arr []int, leftKey, rightKey int) (k, v int) {
	if leftKey > rightKey {
		fmt.Println("未找到")
		return
	}
	middleKey := (leftKey + rightKey) / 2
	if arr[middleKey] < targetVal {
		return BinaryFind(targetVal, arr, middleKey+1, rightKey)
	} else if arr[middleKey] == targetVal {
		return middleKey, arr[middleKey]
	} else {
		return BinaryFind(targetVal, arr, leftKey, middleKey-1)
	}
}

// 二分查找
func find(pattern float64, arr []float64) (s float64) {
	fmt.Println(arr[0])
	if len(arr) <= 1 {
		if arr[0] == pattern {
			return arr[0]
		} else {
			return 0.0
		}
	}
	var middle int
	middle = (len(arr) - 1) / 2
	fmt.Printf("arr = %v,len= %v,middle= %v\n", arr, len(arr), middle)
	if pattern > arr[middle] {
		s = find(pattern, arr[middle+1:len(arr)])
	} else if pattern == arr[middle] {
		s = arr[middle]
	} else {
		s = find(pattern, arr[0:middle+1])
	}
	return s
}
