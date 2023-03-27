/*
给定一个字符串 s 请你找出其中不含有重复字符的最长子串的长度。
"abcabvc"
*/
package main

import "fmt"

//
//func lengthOfLongestSubstring(s string) int {
//	maxCount, p, index, count := 0, 0, 0, 0
//	numMap := make(map[uint8]int)
//
//	for index < len(s) {
//		for _, value := range s {
//			_, ok := numMap[s[value]]
//			if !ok {
//				numMap[s[value]] = index
//				count += 1
//				index += 1
//			} else {
//				if count > maxCount {
//					maxCount = count
//				}
//				p += 1
//				numMap = make(map[uint8]int)
//				index = p
//			}
//		}
//	}
//	return maxCount
//
//}

// 解答：滑动窗口
func lengthOfLongestSubstring2(s string) (res int) {
	chIdx, begin := make(map[rune]int, 26), 0
	for i, c := range s {
		//节省复杂度，根据新窗口的开始位置判断是否满足条件，不需要移除map中的元素
		if j, ok := chIdx[c]; ok && j >= begin {
			res = max(res, i-begin)
			begin = j + 1
		}
		chIdx[c] = i
	}
	res = max(res, len(s)-begin)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	s := "abcadc"
	fmt.Println(lengthOfLongestSubstring2(s))
}
