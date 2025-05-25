package main

/*
Problem Description
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。
你从其中的一个加油站出发，开始时油箱为空。
给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，
否则返回 -1 。如果存在解，则 保证 它是 唯一 的

思路：
全局判断：首先，如果所有加油站的总油量小于总消耗量，那么无论从哪个加油站出发，都不可能完成环路。这种情况下，直接返回 -1。
局部判断与起点更新：我们可以尝试从一个加油站 start 出发，一路向后行驶，记录当前油箱的油量 current_gas。
每经过一个加油站 i，油量增加 gas[i]。
开往下一个加油站 i+1 (或者如果是最后一个加油站，则开往第一个加油站)，油量减少 cost[i]。
如果在任何时刻 current_gas 变为负数，说明从 start 出发，无法到达当前这个点。更重要的是，从 start 到当前点 i 之间的任何一个加油站出发，
也无法到达点 i 的下一个点。为什么呢？因为如果从 start 到 i 之间的某个点 k 出发，到达 i 时油箱里的油量肯定会比从 start 出发时更少（因为少了从 start 到 k-1 积累的油量）。
所以，如果从 start 出发在 i 点油量为负，那么下一个可能的起点至少是 i+1。
基于以上两点，我们可以设计一个贪心算法：

初始化三个变量：
total_tank：记录总的净油量 (总 gas - 总 cost)。用于最终判断是否有解。
current_tank：记录从当前尝试的起点出发，油箱的实时油量。
starting_station：记录当前尝试的起点，初始化为 0。
遍历所有的加油站 i 从 0 到 n-1：
更新 total_tank += gas[i] - cost[i]。
更新 current_tank += gas[i] - cost[i]。
如果 current_tank < 0：
这意味着从当前的 starting_station 出发，无法到达 i+1。
因此，我们将下一个可能的起点更新为 i+1。
重置 current_tank = 0，因为我们从新的起点开始计算。
遍历结束后：
如果 total_tank < 0，说明总油量不足以支持全程，返回 -1。
否则，starting_station 就是那个唯一的解。因为题目保证如果存在解，则解是唯一的。
我们的贪心策略确保了如果存在解，starting_station 会停在那个正确的起点上。当 total_tank >= 0 时，
意味着至少存在一个起点可以完成旅程。我们的算法在 current_tank 变为负数时，将下一个可能的起点设为 i+1。
这个过程实际上排除了所有不可能的起点。如果最终 total_tank >= 0，那么最后确定的 starting_station 一定是可行的
*/
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	totalTank := 0
	currentTank := 0
	startingStation := 0

	for i := 0; i < n; i++ {
		diff := gas[i] - cost[i]
		totalTank += diff
		currentTank += diff

		// 如果当前油量为负，意味着从 startingStation 出发无法到达 i+1
		// 所以，下一个可能的起点是 i+1
		// 并重置 currentTank
		if currentTank < 0 {
			startingStation = i + 1
			currentTank = 0
		}
	}

	// 如果总油量小于总消耗量，则无解
	if totalTank < 0 {
		return -1
	}
	return startingStation
}
