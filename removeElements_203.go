package main

/*
Problem Description
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点
*/

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	// 如果整个链表都被删除了
	if head == nil {
		return nil
	}

	// 2. 处理链表中间和尾部的节点
	// 此时 head.Val != val (或者 head 是 nil，但上面已经处理了)
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next // cur.Next 不需要删除，cur 前进
		}
	}

	return head
}
