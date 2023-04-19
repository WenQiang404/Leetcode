## 哈希表收获心得
#### 242-有效的字母异位词
    1.sort.Slice(x any, less func(i int, j int) bool))
    对切片进行排序，需要先将字符串转化为字节数组：s1 := []bytes(s)
    2.类型rune是int32的别名