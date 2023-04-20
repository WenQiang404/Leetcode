/*
*@Author: Wenqiang
*@Date:   2023/4/20
*@Time:   22:12
 */
package main

import "fmt"

// 方法一：直接暴力法：
// 执行用时：
// 4 ms
// , 在所有 Go 提交中击败了
// 72.15%
// 的用户
// 内存消耗：
// 3 MB
// , 在所有 Go 提交中击败了
// 28.74%
// 的用户
func intersection(nums1, nums2 []int) []int {
	resultNum := make([]int, 0)
	mapAppear := map[int]int{}

	for i := 0; i < len(nums1); i++ {
		_, ok := mapAppear[nums1[i]]
		if !ok {
			for j := 0; j < len(nums2); j++ {
				if nums1[i] == nums2[j] {
					resultNum = append(resultNum, nums2[j])
					break
				}
			}
			mapAppear[nums1[i]]++
		} else {
			continue
		}

	}
	return resultNum
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{1, 1}
	fmt.Println(intersection(nums1, nums2))
}
