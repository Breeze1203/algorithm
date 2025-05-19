package main

/*
Problem Description
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点

思路：初始化快慢指针： 创建两个指针，fast 和 slow，都初始化为 dummy。
快指针先走 n 步： 让 fast 指针先向前移动 n 步。这样，fast 和 slow 之间就相隔了 n 个节点。
同时移动快慢指针： 接着，同时向前移动 fast 和 slow 指针，直到 fast.Next 为 nil (即 fast 指向链表的最后一个节点)。
此时，slow 指针指向的就是要删除节点的前一个节点。
理解关键点： 当 fast 指向链表末尾时，由于 fast 比 slow 多走了 n 步，
所以 slow 刚好停在倒数第 n+1 个节点的位置（或者说，倒数第 n 个节点的前一个节点）。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 创建哨兵节点，简化边界处理（尤其是删除头节点的情况）
	dummy := &ListNode{Val: 0, Next: head}
	fast := dummy
	slow := dummy

	// 1. 快指针先向前移动 n 步
	// 注意：题目保证了 n 是有效的，所以 fast 不会越界到 nil 之前
	// fast 要比 slow 多 n 个节点，所以 fast 从 dummy 开始，走 n+1 步才能让 fast 和 slow 之间有 n 个节点
	// 或者理解为，fast 要指向第 n 个节点（从 dummy 开始算起）
	for i := 0; i < n; i++ {
		if fast.Next == nil { // 处理 n 大于链表长度的情况，虽然题目通常保证 n 有效
			return head // 或者根据题目要求返回特定错误或行为
		}
		fast = fast.Next
	}

	// 2. 同时移动快慢指针，直到快指针到达链表末尾 (fast.Next == nil)
	// 此时 slow 指向的是待删除节点的前一个节点
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 3. 删除目标节点
	// slow.Next 就是要删除的节点
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}

	// 4. 返回新的头节点
	return dummy.Next
}
