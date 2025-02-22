
// 题目：查找字符串 str_2 在 str_1 中起始位置(第一次出现), 如果 不存在 返回 -1,-1
//     如:在字符串abcdefg查找包含cde的字段，将起始位置输出 2,4

package main

import (
	"fmt"
)

func TEST_DO_NOT_CHANGE(str_1 string, str_2 string) (int, int) {
	i_start := -1
	i_end := -1
	// start下面可以改动


	
	// end上面可以改动
	return i_start, i_end
}

func main() {
	fmt.Println(TEST_DO_NOT_CHANGE("abcdefg", "cde"))
	// 下面可以添加测试语句
}
```

请注意，为了使代码能够正常工作，你需要导入 `strings` 包。以下是完整的代码：

```go
package main

import (
	"fmt"
	"strings"
)

func TEST_DO_NOT_CHANGE(str_1 string, str_2 string) (int, int) {
	i_start := -1
	i_end := -1
	// start下面可以改动
	
	if str_2 == "" || !strings.Contains(str_1, str_2) {
		return -1, -1
	}
	i_start = strings.Index(str_1, str_2)
	i_end = i_start + len(str_2) - 1

	// end上面可以改动
	return i_start, i_end
}

func main() {
	fmt.Println(TEST_DO_NOT_CHANGE("abcdefg", "cde"))
	// 下面可以添加测试语句
}
