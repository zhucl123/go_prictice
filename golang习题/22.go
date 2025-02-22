
// 题目
// 使用冒泡排序方法对以下列表进行排序：
// 提示：冒泡排序的步骤：
// 1、比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 2、对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 3、针对所有的元素重复以上的步骤，除了最后一个。
// 4、持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
// 尝试将以下列表用冒泡法进行排序，结果直接输出排序后的列表
// 例如 arr3 = [67,23,8,54,31,49]
// 排序后输出的result形式为： [8, 23, 31, 49, 54, 67]

package main

import "fmt"

func TEST_DO_NOT_CHANGE(arr []int) []int {
    fmt.Println(arr)
    // start下面可以改动
    

    // end上面可以改动
    return result
}

func main() {
    arr1 := []int{51, 23, 79, 11, 19, 59, 66, 8, 97, 72, 36, 49, 81}
    arr2 := []int{3, 78, 34, 8, 56, 12, 9, 24, 45}
    arr3 := []int{67, 23, 8, 54, 31, 49}

    fmt.Println(TEST_DO_NOT_CHANGE(arr1))
    fmt.Println(TEST_DO_NOT_CHANGE(arr2))
    fmt.Println(TEST_DO_NOT_CHANGE(arr3))
}
