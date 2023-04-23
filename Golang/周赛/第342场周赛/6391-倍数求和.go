/*
*@Author: Wenqiang
*@Date:   2023/4/23
*@Time:   10:49
 */
package main

func sumOfMultiples(n int) int {
	count := 0
	for i := 1; i < n+1; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			count += i
		}
	}
	return count
}
