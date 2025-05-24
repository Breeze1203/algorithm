package main

import "fmt"

/*
Problem Description
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积

思路：
想象一下你正在数组中从左到右旅行，每到一个数字，你都要思考：“如果我的连续子数组必须在这里（当前数字）结束，
那么我能得到的最大乘积是多少呢？”
核心挑战：负负得正
如果所有数字都是正数，事情很简单：你只需要一直乘下去，如果乘积变小了就从当前数重新开始。但负数改变了游戏规则：
一个负数乘以一个正数，结果是负数。
一个负数乘以一个负数，结果是正数！
这意味着，一个当前看起来很糟糕的“最小乘积”（比如一个很大的负数），如果遇到下一个负数，
就可能摇身一变成为一个非常大的“最大乘积”。
*/

// maxInt 返回两个整数中的较大者
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// minInt 返回两个整数中的较小者
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		// 根据题目要求，数组通常非空，但作为防御性编程可以加上
		return 0 // 或者根据具体错误处理策略返回错误或 panic
	}

	// overallMaxProduct 存储遍历到目前为止找到的全局最大乘积
	overallMaxProduct := nums[0]
	// maxEndingHere 存储以当前元素结尾的子数组的最大乘积
	maxEndingHere := nums[0]
	// minEndingHere 存储以当前元素结尾的子数组的最小乘积
	// (需要最小乘积是因为当乘以一个负数时，最小可能变成最大)
	minEndingHere := nums[0]

	for i := 1; i < len(nums); i++ {
		currentNum := nums[i]

		// 当我们计算新的 maxEndingHere 和 minEndingHere 时，
		// 需要用到上一轮的 maxEndingHere。
		// 如果不先保存上一轮的 maxEndingHere，
		// 它在计算新的 maxEndingHere 时会被覆盖，
		// 导致计算新的 minEndingHere 时使用的是错误的（当轮更新后的）maxEndingHere。
		prevMaxEndingHere := maxEndingHere

		// 更新以当前元素结尾的最大乘积
		// 它可以是:
		// 1. 当前数字本身 (开始一个新的子数组)
		// 2. 上一个最大乘积 * 当前数字
		// 3. 上一个最小乘积 * 当前数字 (如果当前数字是负数，这可能产生新的最大值)
		maxEndingHere = maxInt(currentNum, maxInt(prevMaxEndingHere*currentNum, minEndingHere*currentNum))

		// 更新以当前元素结尾的最小乘积
		// 它可以是:
		// 1. 当前数字本身
		// 2. 上一个最大乘积 * 当前数字 (使用 prevMaxEndingHere)
		// 3. 上一个最小乘积 * 当前数字
		minEndingHere = minInt(currentNum, minInt(prevMaxEndingHere*currentNum, minEndingHere*currentNum))
		// 更新全局最大乘积
		overallMaxProduct = maxInt(overallMaxProduct, maxEndingHere)
	}
	return overallMaxProduct
}

func main() {
	// 测试用例
	nums1 := []int{2, 3, -2, 4}
	fmt.Printf("输入: %v, 最大乘积: %d (预期: 6)\n", nums1, maxProduct(nums1))
	// 记录当前最大值
}
