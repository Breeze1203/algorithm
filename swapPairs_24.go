package main

/*
Problem Description
两两交换链表中的节点
方法一：利用栈的特性
方法二：迭代
*/

func swapPairs(head *ListNode) *ListNode {
	// 创建虚拟头节点
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		// 定义两个要交换的链表
		node1 := cur.Next
		node2 := cur.Next.Next
		// 开始交换
		// 1. prev 的 Next 指向 node2
		cur.Next = node2
		// 2. node1 的 Next 指向 node2 原来的 Next
		node1.Next = node2.Next
		// 3. node2 的 Next 指向 node1
		node2.Next = node1
		// 移动 prev 指针到下一对要交换的节点的前一个节点
		// 交换后，node1 变成了第二个节点，所以 prev 移动到 node1
		cur = node1
	}
	return dummy.Next

}
