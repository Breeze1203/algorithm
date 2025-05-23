package main

import "strings"

/*
Problem Description
给你一个字符串 s ，请你反转字符串中 单词 的顺序。
单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，
单词间应当仅用单个空格分隔，且不包含任何额外的空格

思路：根据空格将字符串切割得到数组，反转数组
*/
func reverseWords(s string) string {
	// 1. 使用 strings.Fields() 将字符串按一个或多个空白字符分割成单词切片。
	//    这个函数会自动处理前导/尾随空格以及单词间的多个空格，
	//    结果是一个只包含实际单词的切片。
	words := strings.Fields(s)

	// 2. 翻转单词切片中的单词顺序。
	//    使用双指针法原地翻转切片。
	left, right := 0, len(words)-1
	for left < right {
		// 交换 left 和 right 指针指向的单词
		words[left], words[right] = words[right], words[left]
		// 移动指针
		left++
		right--
	}

	// 3. 使用 strings.Join() 将翻转后的单词用单个空格连接起来。
	//    如果 words 切片为空 (例如输入 s 为 "" 或 "   ")，Join 会返回 ""。
	//    如果 words 切片只有一个单词，Join 会返回该单词本身。
	return strings.Join(words, " ")
}
