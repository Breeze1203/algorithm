package main

import (
	"fmt"
	"strings"
)

/*
Problem Description
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
P   A   H   N
A P L S I I G
Y   I   R

解题思路：
处理特殊情况：
如果 numRows == 1，Z 字形就变成了一行，输出与输入相同。
如果 numRows >= len(s)，这意味着每个字符都会在 Z 字形的“竖直”部分，且每一行最多只有一个字符
（或者说，字符串的长度不足以填满一个完整的 "V" 字形下降和上升）。这种情况下，输出也与输入相同，因为逐行读取的结果就是原字符串。
如果输入字符串 s 为空，直接返回空字符串。

模拟填充：
我们可以创建 numRows 个字符串构建器（strings.Builder），每一个代表 Z 字形排列中的一行。
遍历输入字符串 s 中的每个字符。
维护当前字符应该被放置的行号 currentRow 和移动方向 step（向下 +1 或向上 -1）。
初始时，currentRow = 0，step = 1（向下移动）。
将当前字符追加到 rowBuilders[currentRow]。
更新 currentRow：
如果 currentRow 到达最顶行（0），则下一步应该向下移动，所以 step = 1。
如果 currentRow 到达最底行（numRows - 1），则下一步应该向上移动，所以 step = -1。
然后，currentRow += step 来确定下一个字符的行号。

合并结果：
遍历完所有字符后，按顺序（从 rowBuilders[0] 到 rowBuilders[numRows-1]）将每个字符串构建器的内容连接起来，得到最终的 Z 字形排列字符串
*/

func convert(s string, numRows int) string {
	// 特殊情况处理
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	// 虽然上面的条件会处理 len(s) == 0 的情况 (当 numRows >= 0),
	// 但明确写出会更清晰，假设 numRows >= 1。
	if len(s) == 0 {
		return ""
	}

	// 创建 numRows 个 strings.Builder
	rowBuilders := make([]strings.Builder, numRows)
	currentRow := 0
	step := 1 // 1 表示向下, -1 表示向上
	for _, char := range s {
		rowBuilders[currentRow].WriteRune(char)
		// 判断是否到达顶部或底部，需要改变方向
		if currentRow == 0 {
			step = 1 // 到达顶部，开始向下
		} else if currentRow == numRows-1 {
			step = -1 // 到达底部，开始向上
		}
		// 更新当前行
		currentRow += step
	}

	// 合并所有行
	var result strings.Builder
	for i := 0; i < numRows; i++ {
		result.WriteString(rowBuilders[i].String())
	}
	return result.String()
}

func main() {
	s1 := "PAYPALISHIRING"
	numRows1 := 3
	fmt.Printf("Input: s = \"%s\", numRows = %d\n", s1, numRows1)
	fmt.Printf("Output: %s\n", convert(s1, numRows1)) // Expected: PAHNAPLSIIGYIR
}
