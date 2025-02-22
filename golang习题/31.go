
// 题目：实现选择排序法，对 lst_input 从小到大进行排序，排序结果保存到lst_rlt并输出。
// 程序分析：可以利用选择法，即从后9个比较过程中，选择一个最小的与第一个元素交换，下次类推，即用第二个元素与后8个进行比较，并进行交换。
// 请不要用 sort   

package main

import "fmt"

func TEST_DO_NOT_CHANGE(lst_input []int) []int {
    lst_rlt := make([]int, len(lst_input))
    copy(lst_rlt, lst_input)
    fmt.Println(lst_input)
    //start 下面可以改动
    
    
        
    //end 上面可以改动 "
    return lst_rlt
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE([]int{11, 43, 27, 41, 56, 79, 98, 75, 3, 66}))
    #######下面可以添加测试语句
}
