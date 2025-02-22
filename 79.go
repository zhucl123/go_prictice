
// 题目
// 给定一个数组 prices ，它的第 i 个元素prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择某一天买入这只股票，并选择在未来的某一个不同的日子卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(prices []int) int {
    fmt.Println(prices)
    rlt := 0
    // start下面可以改动
    

    // end上面可以改动
    return rlt
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE([]int{7, 1, 5, 3, 6, 4}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{7, 6, 4, 3, 1}))
}
