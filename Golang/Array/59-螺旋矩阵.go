/*
*@Author: Wenqiang
*@Date:   2023/4/25
*@Time:   10:40
 */
package main

func generateMatrix(n int) [][]int {
	startx, starty := 0, 0 //行列的起始位置
	var loop int = n / 2   //循环次数
	offset := 1            //终止下标
	count := 1
	center := n / 2         //若为奇数，中心元素的下标
	res := make([][]int, n) //保存结果的数组
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := startx, starty
		//行不变列变
		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}
		//列不变行变，此时列数为j
		for i = startx; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		//行不变，列从右往左变
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		//列不变，行从下往上递减
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}
	if n%2 == 1 {
		res[center][center] = n * n
	}
	return res
}
