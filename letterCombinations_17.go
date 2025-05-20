package main

/*
Problem Description
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母

思路：
初始化映射表: 创建一个从数字字符到对应字母字符串的映射，例如 {'2': "abc", '3': "def", ...}。
初始化一个空列表 results 来存储所有找到的组合。
处理空输入: 如果输入的 digits 字符串为空，直接返回空列表 results。
调用回溯辅助函数 backtrack，初始时 index 为 0，currentPath 为空
*/
// mapping 存储数字到字母的映射
var phoneMappings = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// backtrackHelper 是核心的回溯辅助函数
// index: 当前处理的 digits 字符串的索引
// digits: 输入的数字字符串
// currentPath: 当前构建的字母路径 (使用 []byte 以提高效率)
// results: 指向结果列表的指针
func backtrackHelper(index int, digits string, currentPath []byte, results *[]string) {
	//说明一个完整组合已经产生
	if index == len(digits) {
		*results = append(*results, string(currentPath))
		return
	}
	// 获取对应下标的数字的key
	currentDigit := digits[index]
	// 获取对应key的value
	possibleLetters := phoneMappings[currentDigit]
	for i := 0; i < len(possibleLetters); i++ {
		letter := possibleLetters[i]
		// 选择当前字母
		currentPath = append(currentPath, letter)
		// 递归进行下一个决策
		backtrackHelper(index+1, digits, currentPath, results)
		// 撤销
		currentPath = currentPath[:len(currentPath)-1]
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var results []string
	// 初始路径为空，从 digits 的第一个字符开始 (index 0)
	backtrackHelper(0, digits, []byte{}, &results)
	return results
}
