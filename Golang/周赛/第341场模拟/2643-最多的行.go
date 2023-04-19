/*
*@Author: Wenqiang
*@Date:   2023/4/19
*@Time:   19:56
 */
package main

func rowAndMaximumOnes(mat [][]int) []int {
	maxOneNum := 0

	index := 0
	for i := 0; i < len(mat); i++ {
		numOne := 0
		length := len(mat[0])
		for j := 0; j < length; j++ {
			if mat[i][j] == 1 {
				numOne++
			}
		}
		if numOne > maxOneNum {
			index = i
			maxOneNum = numOne
		}
	}
	result := make([]int, 0)
	result = append(result, index)
	result = append(result, maxOneNum)
	return result
}

//	maxNum := 0
//	index := 0
//	indexArray := make([]int, 0)
//	for key, val := range mapCount {
//		if val > maxNum {
//			indexArray = nil
//			maxNum = val
//			index = key
//			indexArray = append(indexArray, index)
//		} else if val == maxNum {
//			indexArray = append(indexArray, val)
//		}
//	}
//	minIndex := 100
//	for k := range indexArray {
//		if k < minIndex {
//			minIndex = k
//		}
//	}
//	result := []int{}
//	result = append(result, maxNum)
//	result = append(result, minIndex)
//	return result
//}
