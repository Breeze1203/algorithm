package main

/*
Problem Description
合并两个有序数组
*/

// merge 使用双指针 合并两个非递减数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 定义指针
	p1, p2, p := m-1, n-1, m+n-1
	// 当 nums1 和 nums2 都有元素时比较
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
