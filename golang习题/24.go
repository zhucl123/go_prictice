
// 题目
// 计算并显示一段文本中每个单词的出现频率。
// 例如："Hello world! This is a test. This is only a test."，应输出{'hello': 1, 'world': 1, 'this': 2, 'is': 2, 'a': 2, 'test': 2, 'only': 1}
// 计算并显示以下文本中每个单词的出现频率。
// text1 = "Far from eye, far from heart."
// text2 = "Fields have eyes, and woods have ears."
// text3 = "Good for good is natural, good for evil is manly."

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func TEST_DO_NOT_CHANGE(text string) map[string]int {
	fmt.Println(text)

	// start下面可以改动


	// end上面可以改动
	return result
}

func main() {
	text1 := "Far from eye, far from heart."
	text2 := "Fields have eyes, and woods have ears."
	text3 := "Good for good is natural, good for evil is manly."
	fmt.Println(TEST_DO_NOT_CHANGE(text1))
	fmt.Println(TEST_DO_NOT_CHANGE(text2))
	fmt.Println(TEST_DO_NOT_CHANGE(text3))
}
