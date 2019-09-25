package main

import (
	"errors"
	"fmt"
)

// 队列
func main() {
	queue := &Queue{
		maxSize: 4,
		array:   [4]int{},
		front:   -1,
		rear:    -1,
	}
	queue.Show()
	queue.AddQueue(100)
	queue.Show()
	queue.AddQueue(200)
	queue.Show()
	queue.AddQueue(300)
	queue.AddQueue(400)
	queue.AddQueue(500)
	queue.GetQueue()
	queue.GetQueue()
	queue.GetQueue()
	queue.GetQueue()
	queue.GetQueue()
}

// Queue 队列管理
type Queue struct {
	maxSize int
	array   [4]int // 数组模拟队列
	front   int    // 表示队列的首部
	rear    int    // 表示队列的尾部
}

// AddQueue 往队列添加元素
func (this *Queue) AddQueue(val int) error {
	// 判断队列是否已满 rear>=maxSize
	if this.rear == this.maxSize-1 {
		fmt.Println("超出队列最大规模")
		// 当前的队尾（最后一个元素）是最大值得时候，再添加就会超出队列的最大值
		return errors.New("超出队列最大规模")
	}
	// 队尾+1 并给对应的赋值
	this.rear++
	this.array[this.rear] = val
	fmt.Printf("队列尾部%d添加元素%d成功\n", this.rear, val)
	return nil
}

// Show 显示队列元素
func (this *Queue) Show() {
	fmt.Println("队列展示：")
	// 设计思路（自定义）：this.front表示队首的位置，但是队列不包含队首该元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d] = %d\n", i, this.array[i])
	}
}

// GetQueue 从队列中取值，先进先出
func (this *Queue) GetQueue() (int, error) {
	if this.rear == this.front {
		fmt.Println("queue empty")
		return 0, errors.New("queue empty")
	}
	this.front++
	val := this.array[this.front]
	fmt.Printf("从队列取出%d的值是%d\n", this.front, val)
	return val, nil
}
