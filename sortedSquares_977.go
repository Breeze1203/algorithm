package main

/*
Problem Description
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序

解题思路：因为平方后最大的数一定来自原数组的两端（绝对值最大的数），利用双指针比较两端谁大
*/

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1
	r := n - 1 // 指向结果数组的末尾，从后往前填充

	for left <= right {
		leftSquare := nums[left] * nums[left]    // 3. 计算平方
		rightSquare := nums[right] * nums[right] // 3. 计算平方

		// 比较平方值
		if leftSquare > rightSquare {
			res[r] = leftSquare
			left++
		} else {
			res[r] = rightSquare
			right--
		}
		r--
	}
	return res
}
