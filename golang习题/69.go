
// 题目：最高海拔
// 有一个自行车手打算进行一场公路骑行，这条路线总共由 n + 1 个不同海拔的点组成。自行车手从海拔为 0 的点 0 开始骑行。
// 给你一个长度为 n 的整数数组 gain ，其中 gain[i] 是点 i 和点 i + 1 的 净海拔高度差（0 <= i < n）。请你返回 最高点的海拔 。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(input_ []int) int {
    output := 0
    // start下面可以改动


    // end上面可以改动
    return output
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE([]int{-5, 1, 5, 0, -7}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{-4, -3, -2, -1, 4, 3, 2}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{1, 2, 1, 0, 3, 0, 2, 5}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{1, 0, -1, 5, -6, 7}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{0, 1, 2, 3, -10, 7, 2}))
}
