package main

import . "algorithm/common"

/*
二叉树中的链表
给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。如果在二叉树中，存在一条一直向下的路径，
且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，
那么请你返回 True ，否则返回 False 。一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径
*/

/*
Definition for singly-linked list.
*/

/*
Definition for a binary tree node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	// 如果链表为空，说明匹配成功
	if head == nil {
		return true
	}
	// 如果node为空，则匹配是吧
	if root == nil {
		return false
	}
	//检查当前节点是否与 head匹配
	if checkPath(head, root) {
		return true
	}
	return checkPath(head, root) ||
		isSubPath(head, root.Left) ||
		isSubPath(head, root.Right)
}

// 检查从当前树节点开始的路径是否匹配链表
func checkPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	// 如果树为空或值不匹配，失败
	if root == nil || root.Val != head.Val {
		return false
	}
	//如果匹配的话，比对下一个
	//左节点
	left := checkPath(head.Next, root.Left)
	//右节点
	right := checkPath(head.Next, root.Right)

	return left || right
}
