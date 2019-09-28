package main

import (
	"errors"
	"fmt"
)

// 使用数组来模拟栈
type Stack struct {
	MaxTop int    // 栈最大的数量
	Top    int    // 栈顶
	arr    [5]int // 数组模拟栈
}

func main() {
	stack := &Stack{
		MaxTop: 5,  // 栈最大可存放的数量
		Top:    -1, // 当栈顶为-1时表示栈为空
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.List()
	n, err := stack.Pop()
	if err != nil {
		return
	}
	fmt.Println(n)
	n, err = stack.Pop()
	if err != nil {
		return
	}
	fmt.Println(n)
	n, err = stack.Pop()
	if err != nil {
		return
	}
	fmt.Println(n)
	n, err = stack.Pop()
	if err != nil {
		return
	}
	fmt.Println(n)
	n, err = stack.Pop()
	if err != nil {
		return
	}
	fmt.Println(n)

}

// Push 入栈
func (this *Stack) Push(val int) error {
	// 判断栈是否满
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满")
		return errors.New("stack full")
	}
	this.Top++
	this.arr[this.Top] = val
	return nil
}

// 遍历栈，从栈顶开始遍历
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
	}

	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func (this *Stack) Pop() (int, error) {
	if this.Top == -1 {
		return 0, errors.New("stack pop failed : empty")
	}
	tmp := this.arr[this.Top]
	this.arr[this.Top] = 0
	this.Top--
	return tmp, nil
}
