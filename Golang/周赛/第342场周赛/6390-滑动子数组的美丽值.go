/*
*@Author: Wenqiang
*@Date:   2023/4/23
*@Time:   10:53
 */
package main

import (
	"fmt"
)

// 时间复杂度过高
func getSubarrayBeauty(nums []int, k int, x int) []int {
	start, end := 0, 0
	end = start + k - 1
	resultArray := make([]int, 0)
	numsCopy := make([]int, len(nums))
	for end <= len(nums)-1 {
		copy(numsCopy, nums)
		selectionSort(numsCopy[start : end+1])
		//sort.Slice(numsCopy[start:end+1], func(i, j int) bool {
		//	return numsCopy[i] < numsCopy[j]
		//})
		if numsCopy[start+x-1] < 0 {
			resultArray = append(resultArray, numsCopy[start+x-1])
		} else {
			resultArray = append(resultArray, 0)
		}

		start += 1
		end = start + k - 1

	}
	return resultArray
}

func selectionSort(arr []int) {
	n := len(arr)

	// 选择排序算法
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// 交换 arr[i] 和 arr[minIdx]
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	num := []int{1, -1, -3, -2, 3}
	fmt.Println(getSubarrayBeauty(num, 3, 2))
}
