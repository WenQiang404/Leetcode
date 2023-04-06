/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，
同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/
package main

import (
	"fmt"
)

func moveZeroes(nums []int) (num []int) {
	slow := 0
	tmp := 0
	for fast := 0; fast < len(nums); fast++ {
		if len(nums) <= 1 {
			return nums
			break
		}
		if nums[fast] != 0 {
			tmp = nums[fast]
			nums[fast] = nums[slow]
			nums[slow] = tmp
			slow++
		} else {
			continue
		}
	}
	return nums

}

func main() {
	var testNum = []int{1, 2, 0, 3, 0, 0, 4}
	numFinal := moveZeroes(testNum)
	fmt.Println(numFinal)
}
