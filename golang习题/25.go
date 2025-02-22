
// 题目
// 跳跃游戏
// 非负整数数组 nums，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
// 例如：nums = [3,2,1,0,4]
// 从第0个位置开始，即使可以跳跃3步至第3个位置，但第3个位置的值为0，因此不能再往后跳。
// 结果：无法到达最后一个位置,因此result返回值为：False。
// 现将以下几个数组进行跳跃游戏判断，判断你是否能够到达最后一个下标，如果可以，返回 True ；否则，返回 False 。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(nums []int) bool {
    // start下面可以改动

    

    // end上面可以改动
}

func main() {
    nums1 := []int{1, 2, 3, 4, 5}
    nums2 := []int{2, 3, 1, 1, 4}
    nums3 := []int{1, 1, 1, 0, 1}
    fmt.Println(TEST_DO_NOT_CHANGE(nums1))
    fmt.Println(TEST_DO_NOT_CHANGE(nums2))
    fmt.Println(TEST_DO_NOT_CHANGE(nums3))
}
