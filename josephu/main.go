package main

import "fmt"

// 约瑟夫问题
func main() {
	first := addBoy(5)
	showBoy(first)
}

type boy struct {
	no   int
	next *boy
}

// num 环形列表boy的总数，head表示环形列表的头指针
func addBoy(num int) (head *boy) {
	first := &boy{}
	current := &boy{}
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	for i := 1; i <= num; i++ {
		boy := &boy{
			no: i,
		}
		if i == 1 {
			first = boy
			current = boy
			current.next = first
		} else {
			current.next = boy
			current = boy
			current.next = first
		}
	}
	return first
}

// 显示单向的环形列表
func showBoy(first *boy) {
	// 处理如果环形链表为空的情况
	if first.next == nil {
		fmt.Println("链表为空")
	}
	current := first
	for {
		fmt.Println("boy 编号：", current.no)
		if current.next == first {
			fmt.Println("链表遍历完成，退出遍历")
			break
		}
		current = current.next
	}
}
