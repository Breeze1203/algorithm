package main

/*
Problem Description
单项链表（节点值 Val，指向下一节点 Next）设计链表
双向链表（节点值 Val，指向下一节点 Next，指向前一节点 Prev）
*/

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	dummyHead := &ListNode{Val: -1, Next: nil}
	return MyLinkedList{head: dummyHead, size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head.Next // Start from the first actual node
	for i := 0; i < index; i++ {
		if cur == nil { // Should not happen if size is correct and index is valid
			return -1
		}
		cur = cur.Next
	}
	if cur == nil { // Should not happen
		return -1
	}
	return cur.Val
}

// Standard parameter order: index, val
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode := &ListNode{Val: val, Next: cur.Next}
	cur.Next = newNode
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val) // index = this.size, value = val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val) // index = 0, value = val
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur.Next != nil { // Node to delete exists
		cur.Next = cur.Next.Next
		this.size--
	}
}
