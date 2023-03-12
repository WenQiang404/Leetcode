/*给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。
请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/
package main

import (
	"fmt"
)

//方法一：使用二分法找到一个target，然后向前向后遍历找到起始位置并记录索引下标
func searchRange(num []int, target int) []int {
	var start, end = 0, 0
	var index = -1 //num中初始位置的下标
	null := []int{-1, -1}
	if len(num) == 0 {
		return null
	} else {
		left := 0
		right := len(num) - 1

		for left <= right {

			mid := left + (right-left)/2
			if target == num[mid] {
				index = mid
				break
			} else if target < num[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		}

	}
	if index == -1 {
		return null
	} else {

		for i := index; i >= 0; i-- {
			if num[i] == target {
				start = i
			} else {
				break //由于是有序数组，第一个元素匹配不到就可以break，这样可以减少时间开销
			}
		}
		for j := index; j <= len(num)-1; j++ {
			if num[j] == target {
				end = j
			} else {
				break
			}
		}

	}
	return []int{start, end}

}

//方法二 哈希表

func main() {
	arrayTest := []int{5, 7, 7, 8, 10}
	fmt.Println(searchRange(arrayTest, 8))
}
