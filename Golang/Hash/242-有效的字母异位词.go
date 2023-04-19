/*
*@Author: Wenqiang
*@Date:   2023/4/19
*@Time:   14:43
 */
package main

import (
	"fmt"
	"sort"
)

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

// -------------------------------------------------------
// 方法一：两个字符串排序后进行比较，若相等则true
func methodOne(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

// 方法二，哈希表：s维护一个频次数组，再在t中减去对应字母的频次，若出现负值则说明不相同
func methodTwo(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
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
	fmt.Println(methodOne("anagram", "nagaram"))
	fmt.Println(methodTwo("anagram", "nagaram"))

}
