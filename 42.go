
// 题目：移动零元素
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(nums []int) {
    fmt.Println(nums)
    // start下面可以改动



    // end上面可以改动
}

func main() {
    TEST_DO_NOT_CHANGE([]int{1, 2, 1, 0, 3, 0, 2, 5})
    TEST_DO_NOT_CHANGE([]int{1, 0, -1})
    TEST_DO_NOT_CHANGE([]int{0, 1})
}
