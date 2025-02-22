
// 题目
// 实现一个函数，该函数接受一个字符串 email 作为参数，返回一个布尔值表示该字符串是否为有效的电子邮箱地址。
// 以下是有效电子邮箱地址的一些基本规则（这只是一个简化版本，实际的电子邮箱验证可能更加复杂）：
// 必须包含一个"@"符号。
// "@"前必须有字符。
// "@"后必须至少有一个"."。
// "."后面必须有至少2个到5个字符。
// 不允许有空格。

package main

import (
	"fmt"
	"regexp"
)

func TEST_DO_NOT_CHANGE(email string) bool {
	fmt.Println(email)
	// start下面可以改动
	
	
	// end上面可以改动
	return result
}

func main() {
	fmt.Println(TEST_DO_NOT_CHANGE("example@email.com"))
	fmt.Println(TEST_DO_NOT_CHANGE("example@email"))
	fmt.Println(TEST_DO_NOT_CHANGE("example.com"))
	fmt.Println(TEST_DO_NOT_CHANGE("example@.com"))
}
