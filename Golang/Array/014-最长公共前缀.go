package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	length := len(strs[0])
	if length == 0 {
		return ""
	}
	for index := 0; index < len(strs); index++ {

		if len(strs[index]) <= length {
			length = len(strs[index])
		}
	}
	index := 0
	str := ""
	strFinal := ""
	flag := 0
	for index < length {
		str = strs[0][index : index+1]
		for i := 1; i < length; i++ {
			if str == strs[i][index:index+1] {
				flag = 0
				continue
			} else {
				flag = 1
				break
			}
		}
		if flag == 0 {
			strFinal += str
			index += 1
		} else {
			break
		}

	}
	return strFinal
}

//func longestCommonPrefix(strs []string) string {
//	length := len(strs)
//	if length == 0 {
//		return ""
//	}
//
//	var res = len(strs[0])
//	for l := 0; l < length-1; l++ {
//		end := 0
//		i := len(strs[l+1])
//		if i > res {
//			i = res
//		}
//		for r := 0; r < i; r++ {
//			if strs[l][r] == strs[l+1][r] {
//				end++
//			} else {
//				break
//			}
//
//		}
//		if res >= end {
//			res = end
//		}
//
//	}
//	return strs[0][0:res]
//
//}

func main() {
	s := []string{"flower", "floo", "flight"}
	fmt.Println(longestCommonPrefix(s))
}
