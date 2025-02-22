
// 题目：定义一个函数，实现输入一句话，单词之间使用空格隔开，统计出其中各单词的词频数，
// 并以keyword:count的格式存在一个dict中，返回这个dict。
// 注意 jkl; 不是一个单词，jkl 是一个单词，要把“; + ”等非字符符号去掉
// 即输入：“abc fjf jkl+ abc abc jkl;” 结果是 {"abc":3,"fjf":1,"jkl":2}

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func TEST_DO_NOT_CHANGE(str_line string) map[string]int {
	wordDict := make(map[string]int)
	// 开始作答

	

	// 结束作答  

	return wordDict
}

func main() {
	strLine := "abc fjf jkl+ abc abc jkl;"
	fmt.Println(TEST_DO_NOT_CHANGE(strLine))
}
