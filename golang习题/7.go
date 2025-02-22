
// 题目：已知一个3*3矩阵A,A的元素依次为1-9的平方
//   ｜1^2 2^2 3^2｜ 
// A=｜4^2 5^2 6^2｜，求该矩阵主对角线元素之和。
//   ｜7^2 8^2 9^2｜
// 程序分析：利用双重for循环控制输入二维数组，再将a[i][i]累加后输出。

package main

import "fmt"

func TEST_DO_NOT_CHANGE() int {
    rlt_sum := 0
    A := [][]int{
        {1 * 1, 2 * 2, 3 * 3},
        {4 * 4, 5 * 5, 6 * 6},
        {7 * 7, 8 * 8, 9 * 9},
    }
    // start下面可以改动




    // end上面可以改动
    return rlt_sum
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE())
    // 下面可以添加测试语句
}
