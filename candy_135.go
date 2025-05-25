package main

/*
Problem Description
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
你需要按照以下要求，给这些孩子分发糖果：
每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目

思路：
初始化：

每个孩子最初都至少分到 1 个糖果。为此，我们创建一个与 ratings 数组等长的 candies 数组，并将所有元素初始化为 1。
第一次遍历（从左到右） ➡️：

我们从第二个孩子开始（索引为 1），一直遍历到最后一个孩子。
对于当前孩子 i，如果他的评分 ratings[i] 高于他左边孩子 ratings[i-1] 的评分，那么当前孩子 i 得到的糖果数 candies[i]
必须比左边孩子的糖果数 candies[i-1] 多一个。所以，更新 candies[i] = candies[i-1] + 1。
这次遍历保证了所有向右的递增评分序列都满足了糖果递增的条件。
第二次遍历（从右到左） ⬅️：

我们从倒数第二个孩子开始（索引为 n-2，其中 n 是孩子总数），一直遍历到第一个孩子。
对于当前孩子 i，如果他的评分 ratings[i] 高于他右边孩子 ratings[i+1] 的评分，那么当前孩子 i 得到的糖果数 candies[i] 必须比右边孩子的糖果数 candies[i+1] 多一个。
关键点：在更新 candies[i] 时，它可能已经在第一次遍历中因为左边的孩子而获得了一个较高的糖果数。我们需要同时满足左边和右边的条件。
因此，candies[i] 应该是在其当前值和 candies[i+1] + 1 中的较大者。即 candies[i] = max(candies[i], candies[i+1] + 1)。
这次遍历保证了所有向左的递增评分序列（即向右的递减评分序列）也都满足了糖果递增的条件，并且不会破坏第一次遍历已经建立的规则。
计算总数 ➕：

完成两次遍历后，candies 数组中的每个值就是对应孩子应得的最小糖果数。
将 candies 数组中的所有数字相加，即可得到需要准备的最小糖果总数。
*/

// max 是一个辅助函数，返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	// 1. 初始化：每个孩子至少一个糖果
	candies := make([]int, n)
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	// 2. 从左到右遍历
	// 如果右边的孩子评分更高，则比左边的孩子多一个糖果
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 3. 从右到左遍历
	// 如果左边的孩子评分更高，则比右边的孩子多一个糖果
	// 同时要考虑从左到右遍历时已经分配的糖果，取较大值
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// 4. 计算总糖果数
	totalCandies := 0
	for _, c := range candies {
		totalCandies += c
	}

	return totalCandies
}
