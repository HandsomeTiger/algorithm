package main

import (
	"errors"
	"fmt"
)

// 环形队列的思路可以通过 当前rear+1%maxSize 取模来实现
func main() {
	queue := &CircleQueue{
		maxSize: 4,
		head:    0,
		tail:    0,
	}
	queue.Push(100)
	queue.Push(200)
	queue.Push(300)
	queue.Push(400)
	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Push(400)
	queue.Push(400)
	queue.Pop()
	queue.Push(400)
}

type CircleQueue struct {
	maxSize int
	array   [4]int
	head    int //指向队首
	tail    int //指向队尾
}

// Push 入队列
func (this *CircleQueue) Push(val int) error {
	if this.IsFull() {
		fmt.Println("队列满了")
		return errors.New("full")
	}
	this.array[this.tail] = val
	fmt.Printf("队列%d添加元素%d\n", this.tail, val)
	this.tail = (this.tail + 1) % this.maxSize
	fmt.Println("当前size = ", this.Size())
	return nil
}

// Pop 出队列
func (this *CircleQueue) Pop() (int, error) {
	if this.IsEmpty() {
		fmt.Println("empty")
		errors.New("empty")
	}
	val := this.array[this.head]
	fmt.Printf("队列%d推出元素%d\n", this.head, val)
	this.head = (this.head + 1) % this.maxSize
	fmt.Println("当前size = ", this.Size())
	return val, nil
}

// 判断环形队列为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 判断环形队列是否满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// Size 该环形队列有多少个元素
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}
