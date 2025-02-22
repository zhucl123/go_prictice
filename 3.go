
// 题目
// 使用26个字母和10个数字，产生一个随机组合的6位字符串
// 提示
// 运行一下以下，体会random的强大
// random.choice('tomorrow')
// random.randint(1,10)
// random.shuffle([1,3,5,6,7])

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func TEST_DO_NOT_CHANGE() string {
	chars_ := []rune("abcdefghijklmnopqrstuvwxyz")
	nums_ := []rune("0123456789")
	random_str := ""
	// start下面可以改动 

	


	// end上面可以改动
	return random_str
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(TEST_DO_NOT_CHANGE())
}
