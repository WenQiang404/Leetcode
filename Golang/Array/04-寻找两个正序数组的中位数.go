/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的中位数 。
算法的时间复杂度应该为 O(log (m+n))
*/
package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var index int
	for j := 0; j < len(nums2); j++ {
		nums1 = append(nums1, nums2[j])
	}
	if len(nums1)%2 == 0 {
		index = len(nums1)/2 - 1
	} else {
		index = len(nums1) / 2
	}
	indexA := index
	indexB := index + 1
	var midNum float64
	midNum = float64(nums1[indexA] + nums1[indexB]/2.0)
	return midNum
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 7}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
