package main

import "fmt"

// removeDuplicates 原地删除重复元素，使每个元素最多出现两次，返回新长度
// 参数：nums - 有序整数数组
// 返回值：新长度 k，nums 前 k 个元素包含最多出现两次的元素
func removeDuplicates_two(nums []int) int {
	// 获取数组长度
	n := len(nums)
	// 特殊情况：如果数组长度小于等于 2，直接返回原长度
	// 这是边界条件，确保长度 <= 2 时无需删除
	if n <= 2 {
		return n
	}

	// 初始化计数器 count，从 2 开始
	// 表示前两个元素（索引 0 和 1）默认保留（最多两次允许）
	count := 2
	fmt.Printf("初始状态: count=%d, nums=%v\n", count, nums)
	// 快指针 i 从索引 2 开始遍历
	// 检查从第三个元素开始的元素是否需要保留
	for i := 2; i < n; i++ {
		// 如果当前元素 nums[i] 与 count-2 位置的元素不同
		// 说明 nums[i] 不是第三次或以上出现，可以放入
		if nums[i] != nums[count-2] {
			// 将当前元素放入 count 位置，并将 count 增加 1
			nums[count] = nums[i]
			count++
		}
	}
	// 返回新长度
	return count
}

func main() {
	// 测试用例 1: nums=[1,1,1,2,2,3]
	nums1 := []int{1, 1, 1, 2, 2, 3}
	fmt.Println("测试用例 1: nums=[1,1,1,2,2,3]")
	k1 := removeDuplicates_two(nums1)
	fmt.Printf("k=%d, 前 %d 个元素: %v\n\n", k1, k1, nums1[:k1])
}
