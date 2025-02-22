
// 请只在 "
// ########## 下面可以改动
//
// ########## 上面可以改动 "
// 字典wordList 中从单词 beginWord和 endWord 的 转换序列 是一个按下述规格形成的序列：
//
// 序列中第一个单词是 beginWord 。
// 序列中最后一个单词是 endWord 。
// 每次转换只能改变一个字母。
// 转换过程中的中间单词必须是字典wordList 中的单词。
// 给你两个单词 beginWord和 endWord 和一个字典 wordList ，找到从beginWord 到endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。

package main

import (
	"fmt"
)

func TEST_DO_NOT_CHANGE(beginWord string, endWord string, wordList []string) int {
	fmt.Println(beginWord, endWord, wordList)
	var rlt int
	// ########## 开始作答
	

	// ########## 结束作答  
	return rlt
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Println(TEST_DO_NOT_CHANGE(beginWord, endWord, wordList))
}
