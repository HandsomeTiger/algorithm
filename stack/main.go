package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 使用数组来模拟栈
type Stack struct {
	MaxTop int    // 栈最大的数量
	Top    int    // 栈顶
	arr    [5]int // 数组模拟栈
}

func main() {
	//example()
	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	// 符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "3+2*6-2"
	index := 0
	for index < len(exp) {
		ch := exp[index : index+1]
		//fmt.Println(ch)
		// 把字符转换成对应的ascii编号
		chAscii := int([]byte(ch)[0])
		if operStack.IsOper(chAscii) {
			// 是符号
			// 1.如果符号栈是一个空栈，直接压入栈
			if operStack.Top == -1 {
				operStack.Push(chAscii)
			} else {
				// 2.如果不是一个符号栈不是一个空栈
				if operStack.Priority(operStack.arr[operStack.Top]) >= chAscii {
					// 2.1 如果符号栈的栈顶的优先级大于当前的优先级
					// 从符号栈弹出一个符号，在从数字栈弹出两个数字，进行运算，在把运算结果压入数字栈
					num1, _ := numStack.Pop()
					num2, _ := numStack.Pop()
					res := operStack.Calc(num1, num2, chAscii)
					// 将运算结构压入数栈
					numStack.Push(res)
					// 再把当前符号栈压入符号栈
					operStack.Push(chAscii)
				} else {
					operStack.Push(chAscii)
				}
			}

		} else {
			// 是数字 直接压入数栈
			num, _ := strconv.Atoi(ch)
			numStack.Push(num)
		}
		index++
	}
	// 如果表达式 扫描入栈完成，依次从符号栈取出一个符号，然后从数栈取出两个数字，进行运算
	// 运算的结果压入数栈,循环直到数栈只剩一个元素
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ := numStack.Pop()
		num2, _ := numStack.Pop()
		oper, _ := operStack.Pop()
		res := operStack.Calc(num1, num2, oper)
		numStack.Push(res)
	}

	// 数栈里面最后一个数就是最终的计算结果
	res, _ := numStack.Pop()
	fmt.Printf("表达式的计算结果是：%d", res)
}

// 利用ascii判断是不是符号
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 编写方法，返回运算符的优先级
func (this *Stack) Priority(oper int) int {
	if oper == 42 || oper == 47 {
		return 1
	} else if oper == 43 || oper == 45 {
		return 2
	} else {
		fmt.Println("运算符有误")
	}
	return 0
}
func (this *Stack) Calc(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符有误")
	}
	return res
}

func example() {
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
