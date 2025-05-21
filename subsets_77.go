package main

/*
Problem Description
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集

思路：
初始化一个空的结果列表 result，用来存放所有找到的子集。
定义一个递归函数，我们称之为 backtrack。这个函数将接收以下参数：
startIndex：一个整数，表示我们当前应该从 nums 数组的哪个索引开始考虑元素。
这有助于避免产生重复的子集，并确保每个元素只被考虑一次（或者说，确保子集中的元素是按原数组中的顺序出现的，
尽管子集本身是无序的，但这种处理方式能避免如 [1,2] 和 [2,1] 这样的重复组合，因为我们只考虑一个方向）。
currentSubset：一个列表，表示当前正在构建的子集。
在 backtrack 函数的每一次调用开始时，我们将 currentSubset 的一个副本添加到 result 列表中。
因为 currentSubset 本身就代表一个有效的子集。
然后，我们从 startIndex 开始遍历 nums 数组到末尾
*/

func subsets(nums []int) [][]int {
	// 定义结果集
	result := [][]int{}

	backt_track(nums, 0, []int{}, &result)
	return result
}

func backt_track(nums []int, start int, currentSubset []int, result *[][]int) {
	// 1. 将当前子集的副本添加到结果中
	temp := make([]int, len(currentSubset))
	copy(temp, currentSubset)
	*result = append(*result, temp)

	// 开始遍历
	for i := start; i < len(nums); i++ {
		currentSubset = append(currentSubset, nums[i])
		backt_track(nums, i+1, currentSubset, result)
		currentSubset = currentSubset[:len(currentSubset)-1]
	}
}
