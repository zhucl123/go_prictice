
// 题目
// 请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 例如矩阵matrix = [    [ 1, 2, 3 ],
//     [ 4, 5, 6 ],
//     [ 7, 8, 9 ]
// ]，返回的结果为[1, 2, 3, 6, 9, 8, 7, 4, 5]
// 请按照顺时针螺旋顺序 ，返回以下矩阵中的所有元素。
// matrix1 = [[2,3,1,4],[4,5,1,5],[1,7,4,9],[4,7,2,1]]
// matrix2 = [[4,6,8],[1,8,5],[5,1,7]]
// matrix3 = [[23,17,54,66],[22,77,34,79]]

package main

import "fmt"

func TEST_DO_NOT_CHANGE(matrix [][]int) []int {
    fmt.Println(matrix)
    var result []int
    //start 下面可以改动
    
    
    
    //end 上面可以改动 
    return result
}

func main() {
    matrix1 := [][]int{{2, 3, 1, 4}, {4, 5, 1, 5}, {1, 7, 4, 9}, {4, 7, 2, 1}}
    matrix2 := [][]int{{4, 6, 8}, {1, 8, 5}, {5, 1, 7}}
    matrix3 := [][]int{{23, 17, 54, 66}, {22, 77, 34, 79}}
    matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

    fmt.Println(TEST_DO_NOT_CHANGE(matrix1))
    fmt.Println(TEST_DO_NOT_CHANGE(matrix2))
    fmt.Println(TEST_DO_NOT_CHANGE(matrix3))
    fmt.Println(TEST_DO_NOT_CHANGE(matrix))
}
