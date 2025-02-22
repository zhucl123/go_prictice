
// 题目
// 已知m*n矩阵matrix,将该矩阵顺时针旋转90度，并输出新矩阵。
// 示例：
// matrix = [
//     [1, 2, 3],
//     [4, 5, 6],
//     [7, 8, 9]
// ]
// 正确输出结果为：
// matrix_90 = [[7, 4, 1],[8, 5, 2],[9, 6, 3]]

// 请将以下矩阵顺时针旋转90度，并输出结果。
// matrix1 = [[2,3,1,4],[4,5,1,5],[1,7,4,9],[4,7,2,1]]
// matrix2 = [[4,6,8],[1,8,5],[5,1,7]]
// matrix3 = [[23,17,54,66],[22,77,34,79]]

package main

import "fmt"

func TEST_DO_NOT_CHANGE(matrix [][]int) [][]int {
    result := nil
    // start下面可以改动
    
    
        
    // end上面可以改动
    return result
}

func main() {
    matrix1 := [][]int{{2, 3, 1, 4}, {4, 5, 1, 5}, {1, 7, 4, 9}, {4, 7, 2, 1}}
    matrix2 := [][]int{{4, 6, 8}, {1, 8, 5}, {5, 1, 7}}
    matrix3 := [][]int{{23, 17, 54, 66}, {22, 77, 34, 79}}
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    fmt.Println(TEST_DO_NOT_CHANGE(matrix1))
    fmt.Println(TEST_DO_NOT_CHANGE(matrix2))
    fmt.Println(TEST_DO_NOT_CHANGE(matrix3))
    fmt.Println(TEST_DO_NOT_CHANGE(matrix))
}
