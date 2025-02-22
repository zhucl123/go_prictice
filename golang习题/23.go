
// 题目
// 请将数据集data归一化到0和1之间
// 提示：归一化值x = (x-x(min))/(x(max)-x(min))
// 例如data = [5, 20, 35, 50, 65]，归一化后的值为[0.0, 0.25, 0.5, 0.75, 1.0]
// 请将下列数据集进行归一化并输出结果,结果请保留小数点后2位：
// data1 = [2,8,16,32,64]
// data2 = [3,9,27,45,72]
// data3 = [6,18,42,48,78]

package main

import (
	"fmt"
	"math"
)

func TEST_DO_NOT_CHANGE(data []float64) []float64 {
	fmt.Println(data)
	var result []float64
	// start下面可以改动
	

	// end上面可以改动
	return result
}

func main() {
	data1 := []float64{2, 8, 16, 32, 64}
	data2 := []float64{3, 9, 27, 45, 72}
	data3 := []float64{6, 18, 42, 48, 78}
	fmt.Println(TEST_DO_NOT_CHANGE(data1))
	fmt.Println(TEST_DO_NOT_CHANGE(data2))
	fmt.Println(TEST_DO_NOT_CHANGE(data3))
}
