
// 题目
// 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。
// 找出只出现一次的那两个元素。你可以按任意顺序返回答案。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(nums []int) []int {
    fmt.Println(nums)
    result := make([]int, 0)
    //start 下面可以改动



    //end 上面可以改动
    return result
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE([]int{1, 2, 1, 3, 2, 5}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{-1, 0}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{0, 1}))
}
