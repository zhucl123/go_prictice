
// 题目
// 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
//
// 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(nums []int) int {
    fmt.Println(nums)
    var rlt int
    // start下面可以改动
    

    
    // end上面可以改动
    return rlt
}

func isSorted(arr []int) bool {
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] >= arr[i+1] {
            return false
        }
    }
    return true
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    nums := []int{1, 3, 5, 4, 7}
    fmt.Println(TEST_DO_NOT_CHANGE(nums))
    // #######下面可以添加测试语句
}
