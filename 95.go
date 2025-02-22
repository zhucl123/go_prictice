
// 题目
// 给定一个字符串s和一些长度相同的单词words。找出 s 中恰好可以由words 中所有单词串联形成的子串的起始位置。
// 注意子串要与words 中的单词完全匹配，中间不能有其他字符，但不需要考虑words中单词串联的顺序。

package main

import (
	"fmt"
)

func TEST_DO_NOT_CHANGE(s string, words []string) []int {
	var lst_rlt []int
    // 开始作答
	
    
	
    // 结束作答  
	return lst_rlt
}

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "word"}
	fmt.Println(TEST_DO_NOT_CHANGE(s, words))
	s = "barfoothefoobarman"
	words = []string{"foo", "bar"}
	fmt.Println(TEST_DO_NOT_CHANGE(s, words))
}
