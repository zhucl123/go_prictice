
// 题目
// 将两个升序列表合并为一个新的升序列表并返回。
//
// 示例：
// 输入：lst_1 = [1,2,4], lst_2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
//
// 输入：lst_1 = [], lst_2 = [0]
// 输出：[0]

package main

import "fmt"

func TEST_DO_NOT_CHANGE(lst_1 []int, lst_2 []int) []int {
    fmt.Println(lst_1, lst_2)
    var lst_rlt []int
    // start下面可以改动



    // end上面可以改动
    return lst_rlt
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE([]int{1, 2, 4}, []int{1, 3, 4}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{}, []int{}))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{}, []int{0}))
}
