/*
*@Author: Wenqiang
*@Date:   2023/4/19
*@Time:   14:43
 */
package main

import "fmt"

func isAnagram(s, t string) bool {
	//if len(s) != len(t) {
	//	return false
	//}
	sMap := make(map[string]int)

	tMap := make(map[string]int)
	//将字符串元素与出现频度存入map
	for i := 0; i < len(s); i++ {
		count := 1
		for j := i + 1; j < len(s); j++ {
			if s[j:j+1] == s[i:i+1] {

				count++
			}
		}
		_, ok := sMap[s[i:i+1]]
		if ok {
			continue
		} else {

			sMap[s[i:i+1]] = count
		}

	}
	//重复将t存入
	for i := 0; i < len(t); i++ {
		count := 1
		for j := i + 1; j < len(t); j++ {
			if t[j:j+1] == t[i:i+1] {

				count++
			}
		}
		_, ok := tMap[t[i:i+1]]
		if ok {
			continue
		} else {

			tMap[t[i:i+1]] = count
		}

	}

	//判断两个map是否相等
	if len(sMap) != len(tMap) {
		return false
	}
	for key, val := range sMap {
		if tVal, ok := tMap[key]; ok || tVal != val {
			return false
		}
	}
	for key, val := range tMap {
		if sVal, ok := sMap[key]; ok || sVal != val {
			return false
		}
	}
	return true
}

func main() {
	a := make(map[int]int)
	b := make(map[int]int)
	a[0] = 1
	b[0] = 1
	fmt.Println(isAnagram("anagram", "nagaram"))

}
