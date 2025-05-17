package main

import "math"

/*
Problem Description
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr]
，并返回其长度。如果不存在符合条件的子数组，返回 0

思路，双指针（滑动窗口），比较双指针区间的和与target的值
*/

func minSubArrayLen(target int, nums []int) int {
	// 左指针
	left := 0
	n := len(nums)
	minLength := math.MaxInt32
	currentSum := 0
	for right := 0; right < n; right++ {
		// 计算和
		currentSum += nums[right]
		// 与target比较
		for currentSum >= target {
			minLength = min(minLength, right-left+1)
			// 和减去左指针指向的元素
			currentSum -= nums[left]
			// 移动左指针
			left++
		}
	}
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}
