package main

import (
	"fmt"
	"math"
)

//func minWindow(s string, t string) string {
//	if len(s) < len(t) {
//		return ""
//	}
//	//将字符串s存入数组中
//	sNum := make([]string, 0)
//	var flag bool
//	for v := range s {
//		sNum = append(sNum, s[v:v+1])
//	}
//	i := 0
//	mapT := make(map[int]string)
//	mapTCopy := make(map[int]string)     //复制后的t串map
//	mapMinStringMap := make(map[int]int) //存储找到的满足条件的字串长度和起始索引
//	num := len(t)                        //字串长度
//	for v := range t {
//		newStr := t[v : v+1] //将字符串t分解存入map中，键为对应下标
//		mapT[i] = newStr
//		i++
//	}
//	//生成mapCopy
//	for k, v := range mapT {
//		mapTCopy[k] = v
//	}
//
//	for i := 0; i < len(sNum); i++ {
//		length := 0  //子串长度
//		var left int //滑动窗口左端
//
//		left = i
//		//
//		for _, v := range mapTCopy {
//			if v == sNum[i] {
//				flag = true
//				break
//			} else {
//				continue
//			}
//		}
//		if flag {
//			for right := left; right < len(sNum); right++ {
//				for _, v := range mapTCopy {
//					if v == sNum[right] {
//						num -= 1
//					}
//				}
//
//				if num == 0 { //全部删除则表示全部存在，字串被全部包含
//					length = right - left + 1
//					mapMinStringMap[length] = left
//					for k, v := range mapT {
//						mapTCopy[k] = v
//					}
//					break
//				}
//
//			}
//
//		} else {
//			continue
//		}
//
//	}
//	return getMinStr(sNum, mapMinStringMap)
//}
//func getMinStr(num []string, mapStr map[int]int) string {
//	var str string = ""
//	var minLength = 999999999
//	var minIndex = 0
//	for k, v := range mapStr {
//		if k < minLength {
//			minLength = k
//			minIndex = v
//		}
//	}
//	str = strings.Join(num[minIndex:minIndex+minLength], str)
//	return str
//}

// answer
func minWindowAnswer(s, t string) (ans string) {
	if len(s) < len(t) {
		return ""
	}
	need, window := map[byte]int{}, map[byte]int{}
	left, right, valid := 0, 0, 0 //窗口
	start, length := 0, math.MaxInt32
	for i := range t {
		need[t[i]]++
	}
	for right < len(s) {
		//新加入的数据
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
			//左侧窗口是否要收缩
			for valid == len(need) {
				if right-left < length {
					start = left
					length = right - left
				}
				//待移出的元素
				d := s[left]
				//左移窗口
				left++
				//更新窗口中的数据
				if need[d] > 0 {
					if window[d] == need[d] {
						valid--
					}
					window[d]--
				}
			}
		}

	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

func main() {
	var s string = "aabcvgh"
	var t string = "bc"
	fmt.Println(minWindowAnswer(s, t))
}
