package main

import (
	"fmt"
	"strings"
)

func minWindow(s string, t string) string {
	//将字符串s存入数组中
	sNum := make([]string, 0)
	finalString := ""
	for v := range s {
		sNum = append(sNum, s[v:v+1])
	}
	i := 0
	mapT := make(map[string]int)
	mapTCopy := make(map[string]int) //复制后的t串map
	mapTCopy = mapT
	mapMinStringMap := make(map[int]int) //存储找到的满足条件的字串长度和起始索引
	for v := range t {
		newStr := t[v : v+1] //将字符串t分解存入map中，键为对应下标
		mapT[newStr] = i
		i++
	}

	for i := 0; func() bool {
		_, ok := mapT[sNum[i]]
		return ok
	}(); i++ {
		length := 0  //字串长度
		var left int //滑动窗口左端

		if i == len(s)-len(t)+1 {
			return finalString
		}
		left = i

		for right := left; right < len(sNum); right++ {
			deleteStr := sNum[right]    //待删除元素
			delete(mapTCopy, deleteStr) //若复制后的t串map中存在窗口内的元素，则删除
			if mapTCopy == nil {        //全部删除则表示全部存在，字串被全部包含
				length = right - left + 1
				mapMinStringMap[length] = left
				break
			}

		}

	}
	minLength := 6
	minIndex := 0
	for k, v := range mapMinStringMap {
		if k < minLength {
			minLength = k
			minIndex = v
		}
	}

	finalString = strings.Join(sNum[minIndex:minIndex+minLength], finalString)
	return finalString
}

func main() {
	s := "hello"
	t := "he"
	fmt.Println(minWindow(s, t))
}
