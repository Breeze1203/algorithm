package main

import "fmt"
import . "algorithm/common"

/*
Problem Description
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 基本情况处理：如果链表为空或者 k 小于等于 1 (无需翻转)，直接返回原链表头
	if head == nil || k <= 1 {
		return head
	}
	// 创建一个哨兵节点（dummy node），它的 Next 指向真正的头节点。
	dummy := &ListNode{Next: head}
	// groupPrev 指向每一组需要翻转的 k 个节点的前一个节点 初始时，它指向哨兵节点
	groupPrev := dummy
	for { // 无限循环，直到内部条件使其退出
		// 1. 找到当前组的第 k 个节点 (kthNode)
		kthNode := groupPrev
		for i := 0; i < k; i++ {
			kthNode = kthNode.Next // 向下移动指针
			if kthNode == nil {
				// 如果在找到第 k 个节点之前就到达了链表末尾 (kthNode == nil)，
				// 说明剩余的节点不足 k 个，不需要翻转，直接返回哨兵节点的 Next。
				return dummy.Next
			}
		}
		// 2. 记录第 k 个节点后面的节点 (groupNext)
		// 这将是下一组节点的开始，或者如果这是最后一组，则为 nil。
		groupNext := kthNode.Next
		// 3. 断开当前 k 个节点组成的子链表与后面部分的连接
		// 将 kthNode.Next 置为 nil，使得这 k 个节点成为一个独立的链表段，方便翻转。
		kthNode.Next = nil
		// 4. 翻转这 k 个节点组成的子链表
		//    当前 k-group 的头节点是 groupPrev.Next。
		//    调用 reverseList 函数进行翻转。
		//    翻转后，原 kthNode 成为这个翻转后子链表的头，
		//    而原 groupPrev.Next (即这 k 个节点的第一个节点) 成为这个翻转后子链表的尾。
		// dumy--->2--->1--->nil
		reversedGroupHead := revers(groupPrev.Next)
		// 5. 重新连接链表：
		// groupPrev执行翻转前的头节点，保存下来
		originalGroupHead := groupPrev.Next //(1)

		groupPrev.Next = reversedGroupHead //2--->1连接：前一部分 -> 翻转后的头
		originalGroupHead.Next = groupNext // b) 连接：翻转后的尾 -> 后一部分
		// 6. 移动 groupPrev 指针，为处理下一组做准备。
		groupPrev = originalGroupHead
		// 如果 groupNext 是 nil，说明已经处理到链表的末尾，没有下一组了。
		if groupNext == nil {
			break // 退出 for 循环
		}
	}
	return dummy.Next // 返回哨兵节点指向的头节点，即为修改后的链表头
}

// 翻转链表
func revers(head *ListNode) *ListNode {
	prev := &ListNode{}
	cur := head
	for cur != nil {
		// 定义一个节点暂存后面节点
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// printList 辅助函数：打印链表 (用于测试)
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	fmt.Println("测试用例 1:")
	// 手动创建链表节点
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}

	// 连接节点形成链表: 1 -> 2 -> 3 -> 4 -> 5
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	// 打印原始链表，确认创建正确
	fmt.Printf("原始链表: ")
	printList(&node1) // 传入头节点指针
	reverseKGroup(&node1, 2)
	printList(&node1)
}
