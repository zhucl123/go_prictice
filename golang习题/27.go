
// 题目
// 整数反转，给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 例如： 整数为12345，反转后result 输出：54321
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−2**31,  2**31 − 1] ，就返回 0。

package main

import (
	"fmt"
	"strconv"
)

func TEST_DO_NOT_CHANGE(x int) int {
	var result int

    // start下面可以改动
	

    // end上面可以改动
	return result
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseInt(x int) int {
	s := strconv.Itoa(x)
	reversed, _ := strconv.Atoi(reverseString(s))
	return reversed
}

func main() {
	fmt.Println(TEST_DO_NOT_CHANGE(1589456))
	fmt.Println(TEST_DO_NOT_CHANGE(565893))  
	fmt.Println(TEST_DO_NOT_CHANGE(-2147483648))
}
