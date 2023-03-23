package main

import (
	"fmt"
)

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	character := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}
	numberChannel := make(chan int)
	characterChannle := make(chan string)

	go func() {
		for i := range number {
			numberChannel <- i
		}
	}()

	go func() {
		for j := range character {
			characterChannle <- j
		}
	}()
}
