/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法
*/
package main

import (
	"fmt"
)

// O(n)
func FindIndex(array []int, target int) int {
	index := 0
	//若数组中存在target
	for i := 0; i < len(array); i++ {
		if target > array[len(array)-1] {
			index = len(array)
			break
		}
		if array[i] < target && array[i+1] >= target {
			index = i + 1
		}

	}
	return index

}

// O(log n)
func search(num []int, target int) int {
	high := len(num) - 1
	low := 0
	if num[len(num)-1] < target {
		return len(num)
	} else if num[0] > target {
		return 0
	} else {
		index := 0
		for low <= high {
			mid := low + (high-low)/2
			if num[mid] == target {
				return mid
			} else if num[mid] > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
			index = low
		}
		return index

	}

}
func main() {
	array := []int{1, 3, 5, 6, 7, 9, 10, 13}
	fmt.Println(search(array, 11))
	//test the git states
}
