package main

/*
Problem Description
移除数组中目标value
*/
func removeElement(nums []int, val int) int {
	// 定义慢指针 slow，指向下一个不等于 val 元素应该放置的位置
	slow := 0
	// 定义快指针 fast，遍历数组
	for fast := 0; fast < len(nums); fast++ {
		// 如果当前元素不等于 val，将其放入 slow 位置
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
