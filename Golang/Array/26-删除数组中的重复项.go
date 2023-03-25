/*
给你一个 升序排列 的数组 nums ，请你原地删除重复出现的元素使每个元素只出现一次 ，
返回删除后数组的新长度.元素的相对顺序应该保持一致 。
由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
将最终结果插入 nums 的前 k 个位置后返回 k 。
不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

删除前	[1,2,3,3,4,5,5,6]
删除后	[1,2,3,4,5,6]
*/
package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(nums []int) (int, []int) {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if numMap[nums[i]] == 0 {
			numMap[nums[i]] = nums[i]
		} else {
			continue
		}
	}
	length := len(numMap)
	var arraySort []int
	for _, value := range numMap {
		arraySort = append(arraySort, value)
	}
	//map中的value每次顺序都不一样，需要排序
	sort.Ints(arraySort)

	return length, arraySort
}

func main() {
	testArray := []int{0, 0, 7, 7, 8, 9, 10, 10, 11, 12}
	result, array := removeDuplicates(testArray)

	fmt.Println(result)
	fmt.Println(array)

}
