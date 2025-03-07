
// 题目
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 例如：grid = [
//     [1,3,1,5],
//     [1,5,1,3],
//     [4,2,1,5],
//     [2,5,3,1]
//     ]，求得的最短路径数字总和为11。（1-3-1-1-1-3-1）
// 因此，result输出值为 11
// 有以下几个网格数据，请分别求出最短路径数字总和：
// grid1 = [[2,3,1,4],[4,5,1,5],[1,7,4,9],[4,7,2,1]]
// grid2 = [[1,4,3,2],[1,5,1,5],[2,1,3,1],[1,7,3,8]]
// grid3 = [[3,1,1,2],[6,5,2,1],[5,7,1,2],[6,2,2,3]]

package main

import "fmt"

func TEST_DO_NOT_CHANGE(grid [][]int) int {
    fmt.Println(grid)
    //start 下面可以改动

    
    //end 上面可以改动 
    return 0
}

func main() {
    grid1 := [][]int{{2, 3, 1, 4}, {4, 5, 1, 5}, {1, 7, 4, 9}, {4, 7, 2, 1}}
    grid2 := [][]int{{1, 4, 3, 2}, {1, 5, 1, 5}, {2, 1, 3, 1}, {1, 7, 3, 8}}
    grid3 := [][]int{{3, 1, 1, 2}, {6, 5, 2, 1}, {5, 7, 1, 2}, {6, 2, 2, 3}}
    fmt.Println(TEST_DO_NOT_CHANGE(grid1)) 
    fmt.Println(TEST_DO_NOT_CHANGE(grid2))  
    fmt.Println(TEST_DO_NOT_CHANGE(grid3))  
}
