package main

/*
Problem Description
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样
*/

// reverse 用于原地反转 runes 切片从 start 到 end (包含 end) 的部分
func reverse(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

func reverseStr(s string, k int) string {
	runes := []rune(s) // 将字符串转换为 rune 切片以便修改
	n := len(runes)

	// 以 2k 为步长遍历字符串
	for i := 0; i < n; i += 2 * k {
		// 判断从当前位置 i 开始，到字符串末尾的剩余字符数
		// 1. 如果剩余字符少于 k 个 (n - i < k  或者等价于 i + k > n)
		//    则反转从 i 到末尾的所有字符
		if i+k > n {
			reverse(runes, i, n-1)
		} else {
			// 2. 否则 (剩余字符大于或等于 k 个)
			//    这包含了两种情况：
			//    a) 剩余字符大于等于 2k 个 (标准2k块处理): 反转前 k 个
			//    b) 剩余字符小于 2k 但大于等于 k 个: 依然是反转前 k 个
			//    所以，统一反转从 i 开始的前 k 个字符
			reverse(runes, i, i+k-1)
		}
	}

	return string(runes) // 将修改后的 rune 切片转回字符串
}
