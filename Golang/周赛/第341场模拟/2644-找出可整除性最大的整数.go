/*
*@Author: Wenqiang
*@Date:   2023/4/19
*@Time:   21:46
 */
package main

import "fmt"

// 复杂度有点高，后续思考优化
func find(nums, divisors []int) int {

	countMax := 0
	currentMinNum := 0
	for j := 0; j < len(divisors); j++ {
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i]%divisors[j] == 0 {
				count++
			}
		}

		if count > countMax {
			countMax = count
			currentMinNum = divisors[j]
		}
		if count == countMax && divisors[j] < currentMinNum {
			currentMinNum = divisors[j]
		}

	}
	if currentMinNum == 0 {
		current := divisors[0]
		for i := 1; i < len(divisors); i++ {
			if divisors[i] < current {
				current = divisors[i]
			}
		}
		return current
	} else {
		return currentMinNum
	}

}

func main() {
	a := []int{1}
	b := []int{999999}
	fmt.Println(find(a, b))
}
