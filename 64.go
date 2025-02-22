
// 题目：括号匹配

// 编写一个函数，用于检查给定的字符串中的括号是否匹配。字符串中允许包含小括号 ()、中括号 [] 和大括号 {}，并且括号可以嵌套。

// 检查时，要确保每个左括号都有与之匹配的右括号，并且括号的嵌套关系必须正确。例如，字符串 "()[]{}" 是括号匹配的，而字符串 "{[()]}" 也是括号匹配的，但字符串 "{[)]}" 不是括号匹配的。

// 请编写一个函数 TEST_DO_NOT_CHANGE(s)，其中 s 是输入的字符串，函数应该返回布尔值，表示括号是否匹配。

package main

import "fmt"

func TEST_DO_NOT_CHANGE(s string) bool {
    stack := []rune{}  // 使用栈来辅助判断括号匹配
    //start 下面可以改动

    
    //end 上面可以改动
    return len(stack) == 0  // 最后栈中应该为空
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE("()[]{}"))  // 应该输出 true
    fmt.Println(TEST_DO_NOT_CHANGE("{[()]}"))  // 应该输出 true
    fmt.Println(TEST_DO_NOT_CHANGE("{[)]}"))   // 应该输出 false
}
