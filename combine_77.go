package main

/*
Problem Description
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案

思路：路径 (Path)：我们用一个列表（比如 currentCombination）来存储当前正在构建的组合。
选择列表 (Choices)：在每一步，我们需要从可用的数字中选择一个加入到 currentCombination。为了避免重复的组合（例如 [1,2] 和 [2,1] 对于组合来说是相同的）并确保每个数字只使用一次，我们通常按升序选择数字。这意味着，如果我们已经选择了数字 x，那么下一个数字必须从大于 x 的数中选择。
结束条件 (Base Case for Recursion)：
当 currentCombination 的长度达到 k 时，我们就找到了一个有效的组合。将其加入到最终的结果列表中。
如果我们没有足够的数字可选来填满 k 个位置，也应该停止当前路径的探索
*/

func combine(n int, k int) [][]int {
	// 用于存储所有组合结果
	var results [][]int
	// 特殊情况处理
	if k == 0 {
		return [][]int{{}}
	}
	// 如果k为负数或k大于n，则没有有效组合
	if k < 0 || k > n {
		return [][]int{{}}
	}
	back_track(1, n, k, []int{}, &results)
	return results
}

func back_track(startIndex int, n int, k int, currentPath []int, results *[][]int) {
	// 如果当前组合的长度达到了k
	if len(currentPath) == k {
		// 创建 currentPath 的副本并添加到结果中
		// 必须创建副本，因为 currentPath 在后续的回溯中会被修改
		comboCopy := make([]int, len(currentPath)) // 更安全地使用 len(currentPath)
		copy(comboCopy, currentPath)
		*results = append(*results, comboCopy)
		return
	}
	// 计算当前可选数组的上限
	// 我们还需要k-len(currentPath)
	// 这些数字将从i开始选择
	// 为了能够选够，从i到n至少要有k-len(currentPath)个数字
	// 即n-i+1>=k-len(currentPath)
	// 所有i<=n-(k-len(currentPath)+1)
	limit := n - (k - len(currentPath)) + 1
	for i := startIndex; i <= limit; i++ {
		//选择当前数组
		currentPath = append(currentPath, i)
		//递归探索下一个数字
		back_track(i+1, n, k, currentPath, results)
		// 撤销选择（撤销选择 (回溯)，恢复 currentPath 以便进行其他分支的探索）
		currentPath = currentPath[:len(currentPath)-1]
	}
}
