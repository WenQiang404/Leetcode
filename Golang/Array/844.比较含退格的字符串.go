package main

import (
	"fmt"
	"reflect"
)

func backspaceCompare(s string, t string) bool {
	slow_s := 0
	slow_t := 0
	var array_s = make([]string, len(s))
	var array_t = make([]string, len(s))
	for _, v := range s {
		if v != '#' {
			array_s[slow_s] = string(v)
			slow_s++
		}
		if v == '#' && slow_s == 0 {
			continue
		}
		if v == '#' && slow_s != 0 {
			slow_s--
			array_s[slow_s] = ""
		}
	}
	for _, v := range t {
		if v != '#' {
			array_t[slow_t] = string(v)
			slow_t++
		}
		if v == '#' && slow_t == 0 {
			continue
		}
		if v == '#' && slow_t != 0 {
			slow_t--
			array_t[slow_t] = ""
		}
	}
	if len(array_s) == len(array_t) {
		return reflect.DeepEqual(array_s, array_t)
	} else {
		return false
	}
}

func main() {
	s := "ab#c"
	t := "ad#c"
	fmt.Println(len(s))
	fmt.Println(backspaceCompare(s, t))
}
