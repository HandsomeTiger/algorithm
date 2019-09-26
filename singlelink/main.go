package main

import "fmt"

// 单链表
func main() {
	// 1. 先创建一个头结点（初始）
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "李逵",
		nickname: "黑旋风",
	}
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	ListHeroNode(head)
	DeleteHeroNode(head, 1)
	ListHeroNode(head)
	//InsertHeroNode(head, hero1)
	//ListHeroNode(head)
	DeleteHeroNode2(head, 2)
	ListHeroNode(head)
}

// HeroNode 链表结点
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode // 指向下一个结点
}

// 给链表插入一个结点
// 方式1：在单链表的`最后`加入，传入一个头结点和一个要插入的结点
func InsertHeroNode(head *HeroNode, insertNode *HeroNode) {
	// 1. 先根据头结点找到该链表当前的最后一个结点
	var last *HeroNode = head
	for last.next != nil {
		last = last.next
	}
	fmt.Printf("当前最后一个结点是%v,插入新结点\n", last)
	last.next = insertNode
}

func ListHeroNode(head *HeroNode) {
	out := []*HeroNode{}
	for current := head; current != nil; current = current.next {
		fmt.Printf("当前结点：%v\n", current)
		out = append(out, current)
	}

	for _, v := range out {
		fmt.Println(v)
	}
}

// InsertHeroNodeSort 按照编号从小到大排序插入
func InsertHeroNodeSort(head *HeroNode) {
	// 判断no大于..
}

func DeleteHeroNode(head *HeroNode, no int) {
	temp := head
	for {

		if temp.next.no == no {

			temp.next = temp.next.next
			break
		} else {

			temp = temp.next

		}
	}
}

func DeleteHeroNode2(head *HeroNode, no int) {
	temp := head
	temp.no = 7
}
