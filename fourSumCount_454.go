package main

/*
Problem Description
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足

最直接的想法是使用四层嵌套循环，遍历所有可能的 (i, j, k, l) 组合，然后检查它们的和是否为 0。这种方法当 n 较大时，会导致超时。
优化，我们可以采用哈希表将问题分解。
预计算前两个数组的和：
遍历 nums1 和 nums2 中的所有元素对 (nums1[i], nums2[j])

查找匹配的后两个数组的和：
遍历 nums3 和 nums4 中的所有元素对 (nums3[k], nums4[l])。
计算它们的和 sum34 = nums3[k] + nums4[l]。
现在，我们希望找到 sum12 + sum34 == 0，即 sum12 == -sum34。
对于当前的 sum34，我们检查哈希表 sumMap 中是否存在键 -sum34。
如果存在，那么 sumMap[-sum34] 的值就是与当前 sum34 配对能够使得总和为 0 的 sum12 的数量。
我们将这个数量累加到最终结果 count 中。
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// sumMap 用于存储 nums1[i] + nums2[j] 的和及其出现的次数
	sumMap := make(map[int]int)
	count := 0
	n := len(nums1) // 假设所有数组长度相同

	// 1. 预计算前两个数组的和
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum12 := nums1[i] + nums2[j]
			sumMap[sum12]++
		}
	}

	// 2. 查找匹配的后两个数组的和
	for k := 0; k < n; k++ {
		for l := 0; l < n; l++ {
			sum34 := nums3[k] + nums4[l]
			target := -sum34
			if freq, ok := sumMap[target]; ok {
				count += freq
			}
		}
	}

	return count
}
