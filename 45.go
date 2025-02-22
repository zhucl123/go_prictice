
// 题目
// 找出两个字符串之间的最长公共子字符串。
// 例如：现有2个字符串，s1,s2;
// s1 = "ABABC"
// s2 = "BABCA"
// s1和s2之间的最长公共子字符串即为"BABC"
// 在result中只返回字符串： "BABC"。

// 现有以下3组字符串，请找出每组字符串的最长公共子字符串。
// a1 = "ABABCGRFGGJDJFKLSJDFKLJSLG"
// a2 = "BABCACGJISEJGIOSDJGIOSDGJG"

// b1 = "SDASQWERTYUIOPVJDKKVHNJKDN"
// b2 = "QWERTYUIOPFDSFKSJFFNJSDFFF"

// c1 = "RERFZXCVBNMJKLJKLJKLJKLJKL"
// c2 = "WZXCVBNMFDUSHFUSDHFFHJDSFH"

package main

import (
	"fmt"
)

func TEST_DO_NOT_CHANGE(s1, s2 string) string {
	fmt.Println(s1, s2)

	// start下面可以改动
	

	// end上面可以改动 
	return result
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || contains(s[1:], substr))
}

func main() {
	a1 := "ABABCGRFGGJDJFKLSJDFKLJSLG"
	a2 := "BABCACGJISEJGIOSDJGIOSDGJG"

	b1 := "SDASQWERTYUIOPVJDKKVHNJKDN"
	b2 := "QWERTYUIOPFDSFKSJFFNJSDFFF"

	c1 := "RERFZXCVBNMJKLJKLJKLJKLJKL"
	c2 := "WZXCVBNMFDUSHFUSDHFFHJDSFH"

	fmt.Println(TEST_DO_NOT_CHANGE(a1, a2))
	fmt.Println(TEST_DO_NOT_CHANGE(b1, b2))
	fmt.Println(TEST_DO_NOT_CHANGE(c1, c2))
}
