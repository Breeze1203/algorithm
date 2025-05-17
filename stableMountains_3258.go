package main

// 对于下标 i（i > 0）的山，如果其左侧相邻山的高度 height[i-1] 严格大于 threshold，则该山是稳定的。
// 下标为 0 的山始终不是稳定的，因此我们从下标 1 开始检查
func stableMountains(height []int, threshold int) []int {
	result := make([]int, 0)
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			result = append(result, i)
		}
	}
	return result
}
