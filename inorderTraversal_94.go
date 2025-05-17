package main

import (
	"fmt"
)

// Node 定义二叉树节点
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// inorderTraversal 使用递归实现中序遍历
func inorderTraversal(root *Node) []int {
	result := make([]int, 0, 100) // 预分配容量
	inorderHelper(root, &result)
	return result
}

func inorderHelper(node *Node, result *[]int) {
	if node == nil {
		return
	}
	inorderHelper(node.Left, result)
	*result = append(*result, node.Val)
	inorderHelper(node.Right, result)
}

// inorderTraversalMorris 使用 Morris 遍历实现中序遍历，带中文调试信息
func inorderTraversalMorris(root *Node) []int {
	result := make([]int, 0, 100) // 预分配容量
	var pre *Node                 // 前驱节点
	current := root               // 当前节点

	fmt.Println("开始 Morris 中序遍历")
	for current != nil {
		// 打印节点信息，突出左右子节点
		leftVal := "<nil>"
		rightVal := "<nil>"
		if current.Left != nil {
			leftVal = fmt.Sprintf("%d", current.Left.Val)
		}
		if current.Right != nil {
			rightVal = fmt.Sprintf("%d", current.Right.Val)
		}
		fmt.Printf("当前节点: %d (地址: %p, 左子节点: %s, 右子节点: %s)\n",
			current.Val, current, leftVal, rightVal)

		if current.Left == nil {
			fmt.Printf("  无左子树，访问节点 %d\n", current.Val)
			result = append(result, current.Val)
			fmt.Printf("  当前结果: %v\n", result)
			current = current.Right
			fmt.Printf("  移动到右子树: %v\n", current)
		} else {
			pre = current.Left
			fmt.Printf("  存在左子树，从 %d (地址: %p) 开始寻找前驱\n", pre.Val, pre)
			for pre.Right != nil && pre.Right != current {
				pre = pre.Right
				fmt.Printf("    移动到前驱的右子节点: %d (地址: %p)\n", pre.Val, pre)
			}

			if pre.Right == nil {
				fmt.Printf("  前驱 %d 无右子节点，链接到当前节点 %d\n", pre.Val, current.Val)
				pre.Right = current
				current = current.Left
				fmt.Printf("  移动到左子树: %v\n", current)
			} else {
				fmt.Printf("  前驱 %d 已链接到当前节点 %d，恢复树结构并访问\n", pre.Val, current.Val)
				pre.Right = nil
				result = append(result, current.Val)
				fmt.Printf("  已访问节点 %d，当前结果: %v\n", current.Val, result)
				current = current.Right
				fmt.Printf("  移动到右子树: %v\n", current)
			}
		}
	}
	fmt.Printf("遍历完成，结果: %v\n", result)
	return result
}

// 验证树是否恢复
func verifyTreeRestoration(root *Node) bool {
	if root == nil {
		return true
	}
	visited := make(map[*Node]bool)
	var dfs func(node *Node) bool
	dfs = func(node *Node) bool {
		if node == nil {
			return true
		}
		if visited[node] {
			return false // 检测到循环
		}
		visited[node] = true
		return dfs(node.Left) && dfs(node.Right)
	}
	return dfs(root)
}

// 辅助函数：从数组构建二叉树
func buildTree(values []interface{}, index *int) *Node {
	if *index >= len(values) || values[*index] == nil {
		*index++
		return nil
	}
	val, ok := values[*index].(int)
	if !ok {
		panic(fmt.Sprintf("无效的节点值: %v", values[*index]))
	}
	node := &Node{Val: val}
	*index++
	node.Left = buildTree(values, index)
	node.Right = buildTree(values, index)
	return node
}

// 辅助函数：打印树结构
func printTree(node *Node, level int) {
	if node == nil {
		fmt.Printf("%s[空]\n", indent(level))
		return
	}
	fmt.Printf("%s%d\n", indent(level), node.Val)
	printTree(node.Left, level+1)
	printTree(node.Right, level+1)
}

func indent(level int) string {
	indentStr := ""
	for i := 0; i < level; i++ {
		indentStr += "  "
	}
	return indentStr
}

// 测试函数
func runTestCase(testName string, values []interface{}) {
	fmt.Printf("测试用例 %s: %v\n", testName, values)
	index := 0
	root := buildTree(values, &index)
	fmt.Println("树结构:")
	printTree(root, 0)
	fmt.Println("\n运行递归中序遍历...")
	resultRecursive := inorderTraversal(root)
	fmt.Printf("递归结果: %v\n", resultRecursive)
	fmt.Println("\n运行 Morris 中序遍历...")
	resultMorris := inorderTraversalMorris(root)
	fmt.Printf("Morris 结果: %v\n", resultMorris)
	fmt.Printf("树结构是否恢复: %v\n\n", verifyTreeRestoration(root))
}

func main() {
	// 测试用例 1: [4,2,6,1,3,5,7]
	runTestCase("1", []interface{}{4, 2, 1, 3, 6, 5, 7})

	// 测试用例 4: [5,3,7,2,4,6,8,1,null,null,null,null,null,null,9]
	runTestCase("4", []interface{}{5, 3, 7, 2, 4, 6, 8, 1, nil, nil, nil, nil, nil, nil, 9})

	// 测试用例 5: [10,5,15,3,7,12,18,2,4,6,8]
	runTestCase("5", []interface{}{10, 5, 15, 3, 7, 12, 18, 2, 4, 6, 8})

	// 测试用例 6: [10,5,15,3,7,12,18,2,4,6,8,11,13,16,19]
	runTestCase("6", []interface{}{10, 5, 15, 3, 7, 12, 18, 2, 4, 6, 8, 11, 13, 16, 19})
}
