/*
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5
*/
package main

import "fmt"

// 解法一：暴力法
func pow(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		for count := 1; count < 999999999; count++ {
			if count*count <= x && (count+1)*(count+1) > x {
				return count
			}

		}
	}
	return -1
}

// 解法二：二分搜索法
func pow_2(x int) int {
	start := 1
	end := x
	for start <= end {
		mid := (start + end) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return (start + end) / 2
}
func main() {
	fmt.Println(pow_2(8))
}
