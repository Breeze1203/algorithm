package main

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	//如果k为1 ，每个元素自身就是窗口最大值
	if k == 1 {
		return nums
	}

	return []int{}

}
