
// 题目
// 给你两个没有重复元素的数组nums1和nums2其中nums1是nums2的子集。
//
// 请你找出nums1中每个元素在nums2中的下一个比其大的值。
//
// nums1中数字x的下一个更大元素是指x在nums2中对应位置的右边的第一个比x大的元素。如果不存在，对应位置输出 -1 。
// 输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
// 输出: [-1,3,-1]
// 解释:
//     对于num1中的数字4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
//     对于num1中的数字1 ，第二个数组中数字1右边的下一个较大数字是 3 。
//     对于num1中的数字2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(nums1 []int, nums2 []int) []int {
    fmt.Println(nums1, nums2)
    lst_rlt := []int{}
    // start下面可以改动






    // end上面可以改动
    return lst_rlt
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
