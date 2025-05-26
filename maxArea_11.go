package main

import (
	"fmt"
)

/*
Problem Description
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器
*/

// maxArea 是解决盛最多水容器问题的主函数
func maxArea(height []int) int {
	// 初始化两个指针
	left := 0
	right := len(height) - 1
	maxWater := 0 // 用于存储最大水量

	// 循环直到两个指针相遇
	for left < right {
		// 计算容器的宽度
		width := right - left

		// 计算水的高度 (受限于较短的线)
		h := 0
		if height[left] < height[right] {
			h = height[left]
		} else {
			h = height[right]
		}

		// 计算当前水量
		currentWater := width * h

		// 如果当前水量大于已记录的最大水量，则更新最大水量
		if currentWater > maxWater {
			maxWater = currentWater
		}

		// 移动指向较短线的那个指针
		// 目标是找到一条更高的线来潜在地增加高度
		if height[left] < height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		} else {
			// 如果两条线高度相等，移动哪一个都可以。
			// 通常选择移动一个，例如 left++，以确保所有情况都被考虑到。
			left++
		}
	}

	return maxWater
}

func main() {
	example1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("示例 1 高度: %v, 最大水量: %d\n", example1, maxArea(example1)) // 预期: 49

	example2 := []int{1, 1}
	fmt.Printf("示例 2 高度: %v, 最大水量: %d\n", example2, maxArea(example2)) // 预期: 1

	example3 := []int{4, 3, 2, 1, 4}
	fmt.Printf("示例 3 高度: %v, 最大水量: %d\n", example3, maxArea(example3)) // 预期: 16 (由索引0的4和索引4的4构成)
}
