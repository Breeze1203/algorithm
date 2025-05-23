package main

import "strconv"

/*
Problem Description
给你一个字符串数组 tokens ，表示一个根据 逆波兰表示法 表示的算术表达式。
请你计算该表达式。返回一个表示表达式值的整数

思路：逆波兰表示法（Reverse Polish Notation, RPN），也称为后缀表示法，其核心特点是操作符位于其操作数之后。
例如，普通的中缀表达式 (3 + 4) * 5 对应的逆波兰表示法是 3 4 + 5 *。
计算逆波兰表达式的算法通常依赖于栈数据结构：
初始化一个空栈：这个栈将用来存储运算过程中的操作数。
遍历 Token：从左到右逐个遍历输入字符串数组 tokens 中的每个元素（token）。
如果 Token 是数字：将其转换为整数，并压入栈顶。
如果 Token 是操作符 (+, -, *, /)：
从栈顶弹出两个操作数。注意顺序：先弹出的数是右操作数（operand2），后弹出的数是左操作数（operand1）。
根据当前操作符，执行相应的运算（operand1 operator operand2）。
将运算结果压回栈顶。
处理除法：题目明确指出“两个整数之间的除法总是 向零截断”。
Go 语言中的整数除法 / 对于正负数的行为恰好是向零截断（例如，5 / 2 = 2，-5 / 2 = -2，5 / -2 = -2），所以可以直接使用。
最终结果：当遍历完所有 token 后，栈中将只剩下一个元素，这个元素就是整个表达式的最终计算结果
*/
func evalRPN(tokens []string) int {
	stack := []int{} // 使用切片作为栈
	for _, token := range tokens {
		num, err := strconv.Atoi(token) // 尝试将 token 转换为整数
		if err == nil {
			// 如果是数字，压入栈中
			stack = append(stack, num)
		} else {
			// 如果是操作符
			// 弹出栈顶的两个操作数
			// 注意：先弹出的是右操作数 (operand2)
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2] // 移除已弹出的两个数

			var result int
			switch token {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				result = operand1 / operand2 // Go 整数除法向零截断
			}
			// 将运算结果压回栈中
			stack = append(stack, result)
		}
	}

	// 最终栈中剩下的唯一元素即为结果
	return stack[0]
}
