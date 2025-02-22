
// 题目
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 （只去偷1个晚上），能够偷窃到的最高金额。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(nums []int) int {
    result := -1
    // 开始作答

    
    // 结束作答  
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    nums := []int{1, 2, 3, 1}
    fmt.Println(TEST_DO_NOT_CHANGE(nums))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{2, 3, 2}))
}
