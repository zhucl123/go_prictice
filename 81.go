
// 题目：逆波兰表达式计算器

// 逆波兰表达式是一种数学表达式的表示方法，其中操作符在操作数的后面。例如，表达式 2 + 3 的逆波兰表示法是 2 3 +。
// 逆波兰表达式具有以下特点：
// 1. 逆波兰表达式不需要使用括号，因为操作符的位置已经明确表示了运算的优先级。
// 2. 逆波兰表达式可以直接由计算机进行计算，而无需解析表达式。

// 编写一个函数 TEST_DO_NOT_CHANGE(expression)，它接受一个逆波兰表达式字符串 expression 作为参数，返回表达式的计算结果。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(s string) int {
    output := 0
    // start下面可以改动

    
    // end上面可以改动
    return output
}

func isSpace(c byte) bool {
    return c == ' '
}

func isDigit(c byte) bool {
    return c >= '0' && c <= '9'
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE("2 3 +"))
    fmt.Println(TEST_DO_NOT_CHANGE("3 4 + 2 *"))
    fmt.Println(TEST_DO_NOT_CHANGE("3 4 2 / +"))
}
