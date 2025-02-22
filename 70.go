
// 题目：单词拆分
//
// 给定一个包含非空单词的字典 wordDict，编写一个函数，判断一个非空字符串 s 是否可以被空格拆分成一个或多个字典中出现的单词。
//
// 假设字典为：wordDict = ["cats", "dog", "sand", "and", "cat"]
//
// 对于以下输入：
//
// s = "catsandog"
//
// 应该返回 False，因为 "catsandog" 无法被拆分成字典中的单词。

package main

import "fmt"

var wordDict = []string{"cats", "dog", "sand", "and", "cat"}

func TEST_DO_NOT_CHANGE(s string) bool {
    var output bool
    // start下面可以改动



    // end上面可以改动
    return output
}

func main() {
    fmt.Println(TEST_DO_NOT_CHANGE("catsanddog"))
    fmt.Println(TEST_DO_NOT_CHANGE("sandsdogcat"))
}
