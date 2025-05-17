package main

/*
Problem Description
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数
*/

// 三步旋转 反转整个，反转前k个，反转后n-k个
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 { // 如果数组为空，则无需操作
		return
	}

	k = k % n
	if k == 0 {
		return
	}

	// 1. 反转整个数组 (元素从下标 0 到 n-1)
	for i := 0; i < n>>1; i++ {
		temp := nums[i]
		nums[i] = nums[n-1-i]
		nums[n-1-i] = temp
	}

	// 2. 反转前k个元素 (元素从下标 0 到 k-1)
	for i := 0; i < k>>1; i++ {
		temp := nums[i]
		nums[i] = nums[k-1-i]
		nums[k-1-i] = temp
	}

	// 3. 反转后n-k个元素 (元素从下标 k 到 n-1)
	for i := 0; i < (n-k)>>1; i++ {
		temp := nums[k+i]
		nums[k+i] = nums[n-1-i]
		nums[n-1-i] = temp
	}
}
