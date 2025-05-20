package main

import "fmt"

/*
Problem Description
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合

思路：
对于生成括号问题，我们可以这样思考：
状态表示：在构建括号字符串的过程中，我们需要知道以下信息：
当前已经构建的字符串 currentString。
当前已经使用的左括号数量 openCount。
当前已经使用的右括号数量 closeCount。
总共需要生成的括号对数 n。
决策（如何添加括号）：
什么时候可以添加左括号 ( ？ 只要我们还有未使用的左括号（即 openCount < n），我们就可以添加一个左括号。
什么时候可以添加右括号 ) ？ 只要我们还有未使用的右括号，并且当前已经添加的右括号数量小于已经添加的左括号数量（即 closeCount < openCount），
我们就可以添加一个右括号。这个条件是为了保证括号的有效性，即不能在没有对应左括号的情况下提前闭合。

递归终止条件（什么时候算构建完成一个有效的组合？）：
当当前构建的字符串 currentString 的长度达到 2 * n（因为 n 对括号总共有 2n 个字符），并且 openCount == n 且 closeCount == n 时，
我们就找到了一个有效的括号组合。将其添加到结果列表中。
回溯过程：
我们从一个空字符串开始。
在每一步，我们尝试添加一个 ( (如果条件允许)。添加后，我们递归调用函数，更新 currentString、openCount。
递归调用返回后（意味着基于当前 ( 的所有可能性都探索完了），我们需要撤销刚才的选择（即把 currentString 恢复到添加 ( 之前的状态），
以便尝试其他的可能性。不过，在 Go 中，如果字符串是值传递或者在递归调用时创建新字符串，则不需要显式撤销字符串本身，但计数器需要正确传递。通常我们通过构建新的字符串传递给下一层递归，或者使用 []byte / strings.Builder 并进行删除操作。
然后，我们尝试添加一个 ) (如果条件允许)。添加后，递归调用函数，更新 currentString、closeCount。
同样，递归返回后，撤销选择
*/
func generateParenthesis(n int) []string {
	var result []string
	backtrack(&result, []byte{}, 0, 0, n)
	return result
}

// result：存储最终所有有效的括号组合
// currentPath: 当前正在构建的括号路径
// openCount: 当前路径中左括号的数量
// closeCount: 当前路径中右括号的数量
// maxPairs: 需要生成的括号对数
func backtrack(result *[]string, currentPath []byte, openCount int, closeCount int, maxPairs int) {
	// 终止条件：当路径长度达到2*n时候，说明一个完整的组合已生成
	if len(currentPath) == maxPairs*2 {
		*result = append(*result, string(currentPath))
		return
	}
	// 尝试添加左括号 '('
	if openCount < maxPairs {
		// 做选择
		currentPath = append(currentPath, '(')
		// 进入下一层决策
		backtrack(result, currentPath, openCount+1, closeCount, maxPairs)
		// 撤销选择 (将 currentPath 恢复到添加 '(' 之前的状态)
		currentPath = currentPath[:len(currentPath)-1]
	}

	// 尝试添加右括号 ')'
	if closeCount < openCount {
		currentPath = append(currentPath, ')')
		// 进入下一层决策
		backtrack(result, currentPath, openCount, closeCount+1, maxPairs)
		// 撤销选择
		currentPath = currentPath[:len(currentPath)-1]
	}
}

func main() {
	n1 := 3
	fmt.Printf("n = %d, 括号组合: %v\n", n1, generateParenthesis(n1))
	// 预期输出: ["((()))","(()())","(())()","()(())","()()()"] (顺序可能不同)

}
