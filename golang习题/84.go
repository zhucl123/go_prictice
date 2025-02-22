
// 题目
// 整数数组number，判断number中是否存在三个元素a，b，c，使得a + b + c = 0？找出所有满足条件且不重复的三元组。
// 例如：
// number = [-1, 0, 1, 2, -1, -4, 4, -3, 7, -5]
// result 形式如下：
// [[-5, 1, 4], [-4, -3, 7], [-4, 0, 4], [-3, -1, 4], [-3, 1, 2], [-1, -1, 2], [-1, 0, 1]]
// 现有如下几个数组：
// number1 = [-1, 0, 1, 2, -1, -4, 3]
// number2 = [-2, 0, 1, 1, -1, 3, 4]
// number3 = [0, -1, 2, -3, 1, 2, 4]

package main

import (
	"fmt"
)

func TEST_DO_NOT_CHANGE(nums []int) [][]int {
	fmt.Println(nums)
	var result [][]int
	// start下面可以改动



	// end上面可以改动
	return result
}

func main() {
	number1 := []int{-1, 0, 1, 2, -1, -4, 3}
	number2 := []int{-2, 0, 1, 1, -1, 3, 4}
	number3 := []int{0, -1, 2, -3, 1, 2, 4}
	fmt.Println(TEST_DO_NOT_CHANGE(number1))
	fmt.Println(TEST_DO_NOT_CHANGE(number2))
	fmt.Println(TEST_DO_NOT_CHANGE(number3))
}
