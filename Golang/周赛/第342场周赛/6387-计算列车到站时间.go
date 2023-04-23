/*
*@Author: Wenqiang
*@Date:   2023/4/23
*@Time:   10:46
 */
package main

import "fmt"

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}

func main() {
	fmt.Println(findDelayedArrivalTime(13, 11))
}
