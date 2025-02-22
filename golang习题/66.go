
// 题目：校验密码强度

// 编写一个函数，用于校验用户输入的密码强度。密码强度的规则如下：
// 1. 密码长度必须至少为8个字符。
// 2. 密码必须包含至少一个大写字母、一个小写字母、一个数字和一个特殊字符（例如：!@#$%^&*）。
// 3. 密码不能包含连续的3个字符，如 "123" 或 "abc"。

package main

import (
	"fmt"
	"regexp"
)

func TEST_DO_NOT_CHANGE(password string) bool {
	var result bool
	//start 下面可以改动
	

	//end 上面可以改动
	return result
}

func main() {
	fmt.Println(TEST_DO_NOT_CHANGE("Abc123!@"))    // False
	fmt.Println(TEST_DO_NOT_CHANGE("password1"))    // False
	fmt.Println(TEST_DO_NOT_CHANGE("Abcde12345"))  // False
}
