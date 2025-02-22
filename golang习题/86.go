
// 题目
// 加油站问题：
// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
// 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(gas []int, cost []int) int {
    var result int
    
    // start下面可以改动

    // end上面可以改动
    return result
}

func main() {
    gas1 := []int{2, 3, 4, 5, 1}
    cost1 := []int{5, 1, 2, 3, 4}

    gas2 := []int{3, 3, 4, 0, 2}
    cost2 := []int{2, 3, 4, 0, 1}

    gas3 := []int{5, 1, 2, 3, 4}
    cost3 := []int{4, 4, 1, 5, 1}

    gas := []int{1, 2, 3, 4, 5}
    cost := []int{3, 4, 5, 1, 2}

    fmt.Println(TEST_DO_NOT_CHANGE(gas1, cost1))
    fmt.Println(TEST_DO_NOT_CHANGE(gas2, cost2))
    fmt.Println(TEST_DO_NOT_CHANGE(gas3, cost3))
    fmt.Println(TEST_DO_NOT_CHANGE(gas, cost))
}
