package main

/*
字符串及其反转中是否存在同一子字符串
给你一个字符串 s ，请你判断字符串 s 是否存在一个长度为 2 的子字符串，在 s 反转后的字符串中也出现。
如果存在这样的子字符串，返回 true；如果不存在，返回 false
步骤 1：提取字符串 s 中所有长度为 2 的子字符串。
步骤 2：生成反转后的字符串 reverse(s)。
步骤 3：检查每个长度为 2 的子字符串是否出现在 reverse(s) 中。
为了高效实现，我们可以使用集合（set）来存储 s 的所有长度为 2 的子字符串，然后检查这些子字符串是否出现在 reverse(s) 中。
由于只需要检查长度为 2 的子字符串，我们可以直接遍历 s 的相邻字符对（如 s[i:i+2]
*/

func isSubstringPresent(s string) bool {
	if len(s) < 2 {
		return false
	}

	// 初始化一个map结构，保存提取对字符串
	substrings := make(map[string]bool)
	for i := 0; i < len(s)-1; i++ {
		substr := s[i : i+2]
		substrings[substr] = true
	}
	// 检查map结构里面是否存在反转后对字符串
	for substr := range substrings {
		// 反转子字符串
		reversedSubstr := string(substr[1]) + string(substr[0])
		if substrings[reversedSubstr] {
			return true
		}
	}
	return false

}
