package main

/*
Problem Description
你在一个果园里，面前有一排果树 🌳🌳🌳🌳🌳，每棵树上都长着一种水果 🍎🍊🍇🍌🍉。
你有两个篮子 🧺🧺，但每个篮子只能装一种水果。
你的规则是：
✅ 你可以从任何一棵树开始摘水果，一直往右走 ➡️。
✅ 每棵树都必须摘一个水果 🍏🍊🍇。
✅ 最多只能装两种水果，如果遇到第三种水果 🚫，就必须停下！
你的目标 🎯：
👉 看看最多能摘多少个水果！ 🍎🍊🍇
举个例子：
如果果树是：
[🍏, 🍊, 🍏]
✅ 你可以摘 全部 3 颗树，因为你篮子里只装了 🍏 和 🍊 这两种水果。
✅ 答案是 3 🎉
如果果树是：
[🍌, 🍎, 🍇, 🍇]
如果从第一棵树（🍌）开始，你能摘 [🍌, 🍎]，只能摘 2 颗 ❌
但如果从 🍎 开始，可以摘 [🍎, 🍇, 🍇]，一共 3 颗！ ✅
✅ 答案是 3 🎉
如果果树是：
[🍏, 🍊, 🍇, 🍊, 🍊]
🚫 你不能摘 [🍏, 🍊, 🍇, 🍊]，因为出现了第三种水果 🍇！
✅ 但是你可以从 🍊 开始，摘 [🍊, 🍇, 🍊, 🍊]，最多摘 4 颗！
*/
func totalFruit(fruits []int) int {
	n := len(fruits)
	if n == 0 {
		return 0
	}
	// 最终摘取的果子数
	maxPicked := 0
	left := 0
	basket := make(map[int]int) // 用来存储当前窗口中水果种类及其数量 (fruitType -> count)
	for right := 0; right < n; right++ {
		currentFruitType := fruits[right]
		// 将当前水果放入篮子 (增加其计数)
		basket[currentFruitType]++
		for len(basket) > 2 {
			// 减少左边水果的计数
			leftFruitType := fruits[left]
			basket[leftFruitType]--
			if basket[leftFruitType] == 0 {
				// 如果这种水果数量为0，从篮子中移除该种类
				delete(basket, leftFruitType)
			}
			left++
		}
		maxPicked = max(maxPicked, right-left+1)
	}
	return maxPicked
}
