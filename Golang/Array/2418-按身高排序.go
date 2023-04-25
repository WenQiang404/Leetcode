/*
*@Author: Wenqiang
*@Date:   2023/4/25
*@Time:   20:36
 */
package main

import "fmt"

//通过索引排序然后输出相应名字，只通过了8个用例不说，巨复杂

//func sortPeople(names []string, heights []int) []string {
//	indexArray := make([]int, 0)
//	index := 0
//	count := 0
//	for count < len(heights)-1 {
//		maxElement := heights[0]
//		for i := 0; i < len(heights); i++ {
//			if heights[i] > maxElement {
//				maxElement = heights[i]
//				index = i
//			}
//		}
//		count++
//		indexArray = append(indexArray, index)
//		heights[index] = -1
//	}
//	for j := 0; j < len(heights); j++ {
//		if heights[j] != -1 {
//			indexArray = append(indexArray, j)
//			break
//		}
//	}
//
//	finalArray := make([]string, 0)
//	for m := 0; m < len(indexArray); m++ {
//		finalArray = append(finalArray, names[indexArray[m]])
//	}
//	return finalArray
//}

func sortPeople(names []string, heights []int) []string {
	nameTallMap := make(map[int]string, 0)
	for i := 0; i < len(names); i++ {
		nameTallMap[heights[i]] = names[i]
	}
	max := 0
	maxStr := ""
	count := 0
	finalArray := make([]string, 0)
	for count < len(nameTallMap) {
		for key, val := range nameTallMap {
			if key > max && val != "" {
				max = key
				maxStr = val
			}
		}
		finalArray = append(finalArray, maxStr)
		count++
		nameTallMap[max] = ""
		max = 0
	}
	return finalArray

}
func main() {
	stringa := []string{"a", "b", "c"}
	het := []int{155, 185, 150}
	fmt.Println(sortPeople(stringa, het))
}
