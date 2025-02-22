
// 题目
// 给你一个链表，删除列表的倒数第 n 个结点，并且返回列表。
//
// 示例：
// 输入：lst = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]

package main

import "fmt"

func TEST_DO_NOT_CHANGE(lst []int, n int) []int {
    fmt.Println(lst, n)
    lst_rlt := []int{}
    // start下面可以改动



    // end上面可以改动
    return lst_rlt
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE([]int{1, 2, 3, 4, 5}, 2))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{1}, 1))
    fmt.Println(TEST_DO_NOT_CHANGE([]int{1, 2}, 1))
}
