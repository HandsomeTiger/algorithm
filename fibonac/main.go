package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibo(10))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci2(10))
	//fmt.Println(arr)
}

// fibo 递归的斐波那契列 0,1,1,2,3,5,8,13,21 .... 第n个的值等于 n-1列的值加n-2列的值，在第一列的时候返回1
func fibo(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

// fibonacci 非递归实现，借助数组
func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		arr := make([]int, n+1)
		arr[0] = 0
		arr[1] = 1
		for i := 2; i <= n; i++ {
			arr[i] = arr[i-1] + arr[i-2]
		}
		fmt.Println(arr)
		return arr[n]
	}
}

// fibonacci2
func fibonacci2(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		first := 0
		second := 1
		res := 1
		for i := 2; i <= n; i++ {
			res = first + second
			first = second
			second = res
		}
		return res
	}
}
