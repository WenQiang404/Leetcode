/*给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。
不能使用任何内置的库函数，如  sqrt 。
*/
package main

import (
	"fmt"
)

func find(x int) bool {
	start := 1
	end := x

	for start <= end {
		mid := (start + end) / 2
		if mid*mid == x {
			return true

		} else if mid*mid > x {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}
	return false

}

func main() {
	fmt.Println(find(15))
}
