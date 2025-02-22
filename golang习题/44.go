
// 题目
// 求最大子序列和，整数数组number如下，找到一个具有最大和值的连续子数组并返回该和值：
// 例如：number = [-2, 1, -3, 4, -1, 2, 1, -5, 4]，其中最大子序列和为[4,-1,2,1]，结果result 直接返回： 6
// 求以下数组的最大子序列和：
// number1 = [-1,2,3,-3,5,-5,1,5,-1,6,-4,3,-6]
// number2 = [1,-2,4,-3,4,-1,4,-7,-1,-6,5,4,-1]
// number3 = [2,-2,1,-1,3,4,-1,2,-1,3,-1,3,-2]

package main

import "fmt"

func TEST_DO_NOT_CHANGE(nums []int) int {
    fmt.Println(nums)
    var result int
    // start下面可以改动
    
    

    // end上面可以改动
    return result
}

func main() {
    number1 := []int{-1, 2, 3, -3, 5, -5, 1, 5, -1, 6, -4, 3, -6}
    number2 := []int{1, -2, 4, -3, 4, -1, 4, -7, -1, -6, 5, 4, -1}
    number3 := []int{2, -2, 1, -1, 3, 4, -1, 2, -1, 3, -1, 3, -2}
    number := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
    fmt.Println(TEST_DO_NOT_CHANGE(number1))
    fmt.Println(TEST_DO_NOT_CHANGE(number2))
    fmt.Println(TEST_DO_NOT_CHANGE(number3))
    fmt.Println(TEST_DO_NOT_CHANGE(number))
}
