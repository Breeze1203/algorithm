package main

/*
Problem Description
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，
评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。
注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况

思路：
检测是否有环：
我们使用两个指针，slow 和 fast，都从链表的头节点 head 开始。
slow 指针每次移动一步 (slow = slow.Next)。
fast 指针每次移动两步 (fast = fast.Next.Next)。
如果链表中没有环，fast 指针或 fast.Next 最终会到达 nil（链表尾部）。
如果链表中有环，fast 指针最终会追上 slow 指针，即它们会指向同一个节点。此时，我们可以确定链表中有环。

找到环的入口节点：
当 slow 和 fast 相遇时，我们知道链表中存在环。现在需要找到环的起始节点。
一个关键的结论是：当 slow 和 fast 相遇后，我们将其中一个指针（例如 slow）重新指向链表的头节点 head。
然后，我们将 fast 指针留在相遇点，并将两个指针都以相同的速度（每次一步）向前移动。
它们下一次相遇的节点就是环的入口节点
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	// 第一阶段：检测是否有环
	hasCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
