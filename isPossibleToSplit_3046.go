package main

// 判断数组能否切割
func isPossibleToSplit(nums []int) bool {
	// 用map统计每个元素出现次数，如果大于2则不能切割
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num] > 2 {
			return false
		}
	}
	return true
}
