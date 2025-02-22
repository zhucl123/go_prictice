
// 题目
// 复原IP地址：
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。
// 你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
// 例如：s = "101023"输出的结果为['1.0.10.23', '1.0.102.3', '10.1.0.23', '10.10.2.3', '101.0.2.3']
// 请分别复原以下几个字符串：
// s1 = "25525511135"
// s2 = "0000"
// s3 = "192168101"

package main

import (
	"fmt"
)

func TEST_DO_NOT_CHANGE(s string) []string {
	fmt.Println(s)

	//start 下面可以改动

	//end 上面可以改动 
	return result
}

func main() {
	s1 := "25525511135"
	s2 := "0000"
	s3 := "192168101"
	fmt.Println(TEST_DO_NOT_CHANGE(s1))
	fmt.Println(TEST_DO_NOT_CHANGE(s2))
	fmt.Println(TEST_DO_NOT_CHANGE(s3))
}
```

请注意，为了使代码完整并能编译运行，你需要导入 `strings` 和 `strconv` 包。以下是完整的导入部分：

```go
import (
	"fmt"
	"strconv"
	"strings"
)
