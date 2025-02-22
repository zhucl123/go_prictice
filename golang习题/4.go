
// 题目：x=10, y=20,  将两个变量值互换  x=20，y=10。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(x, y int) (int, int) {
    fmt.Println(x, y)
    // start下面可以改动

    

    // end上面可以改动
    return x, y
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE(10, 20))
    // 下面可以添加测试语句
}
