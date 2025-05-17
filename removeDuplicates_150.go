package main

import "fmt"

// removeDuplicates 移除有序数组中的重复元素，返回唯一元素的数量 k
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 初始化慢指针 slow，指向当前唯一元素的末尾位置
	slow := 0
	// 快指针 i 用于检查每个元素，决定是否需要保留
	for i := 0; i < len(nums); i++ {
		// 如果当前元素 nums[i] 与慢指针位置的元素 nums[slow] 不同
		// 说明 nums[i] 是一个新的唯一元素
		if nums[i] != nums[slow] {
			// 慢指针向前移动一位，准备放置新的唯一元素
			slow++
			nums[slow] = nums[i]
		}
	}

	// 返回唯一元素的数量
	// slow 是最后一个唯一元素的索引，因此数量为 slow+1
	fmt.Printf("最终结果: k=%d, nums=%v\n", slow+1, nums)
	return slow + 1
}

func main() {
	// 测试用例 1: nums=[1,1,2,2,3,3]
	nums1 := []int{1, 1, 2, 2, 3, 3}
	fmt.Println("测试用例 1: nums=[1,1,2,2,3,3]")
	k1 := removeDuplicates(nums1)
	fmt.Printf("k=%d, 前 %d 个元素: %v\n\n", k1, k1, nums1[:k1])

	// 测试用例 2: nums=[1,1,1]
	nums2 := []int{1, 1, 1}
	fmt.Println("测试用例 2: nums=[1,1,1]")
	k2 := removeDuplicates(nums2)
	fmt.Printf("k=%d, 前 %d 个元素: %v\n\n", k2, k2, nums2[:k2])

	// 测试用例 3: nums=[1,2,3]
	nums3 := []int{1, 2, 3}
	fmt.Println("测试用例 3: nums=[1,2,3]")
	k3 := removeDuplicates(nums3)
	fmt.Printf("k=%d, 前 %d 个元素: %v\n\n", k3, k3, nums3[:k3])

	// 测试用例 4: nums=[]
	nums4 := []int{}
	fmt.Println("测试用例 4: nums=[]")
	k4 := removeDuplicates(nums4)
	fmt.Printf("k=%d, 前 %d 个元素: %v\n", k4, k4, nums4[:k4])
}
