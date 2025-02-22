
// 题目
// 给定一个整数数组，判断是否存在重复元素。
// 如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(nums []int) bool {
    fmt.Println(nums) // Python的print语句在Go中对应的是fmt.Println，但这里不需要输出nums
    var result bool
    // start下面可以改动



    // end上面可以改动
    return result
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE([]int{1, 2, 3, 1}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{1, 2, 3, 4}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
