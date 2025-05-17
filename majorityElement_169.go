package main

import "fmt"

/*
Problem Description
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素

Ideas
正常思路是遍历每个元素，并在map里面统计元素出现次数

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题
*/

func majorityElement_boyerMoore(nums []int) int {
	candidate := 0 // 可以用 nums[0] 初始化，因为题目保证数组非空
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	// 因为题目保证多数元素总是存在，所以最后 candidate 一定是多数元素。
	// 如果不保证，则需要再遍历一遍数组，验证 candidate 的数量是否 > n/2。
	return candidate
}

func main() {
	nums1 := []int{3, 2, 3}
	fmt.Printf("输入: %v, 输出 (Boyer-Moore): %d\n", nums1, majorityElement_boyerMoore(nums1)) // 输出: 3

	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("输入: %v, 输出 (Boyer-Moore): %d\n", nums2, majorityElement_boyerMoore(nums2)) // 输出: 2

	nums3 := []int{1}
	fmt.Printf("输入: %v, 输出 (Boyer-Moore): %d\n", nums3, majorityElement_boyerMoore(nums3)) // 输出: 1

	nums4 := []int{6, 5, 5}
	fmt.Printf("输入: %v, 输出 (Boyer-Moore): %d\n", nums4, majorityElement_boyerMoore(nums4)) // 输出: 5
}
