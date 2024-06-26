package main

import (
	"fmt"
	"time"
)

func timesThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem * 3
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int, len(arr))

	go func(arr []int, ch chan int) {

		for _, elem := range arr {
			ch <- elem * 3
		}
	}(arr, ch)
	time.Sleep(2 * time.Second)
//instead of just using an anonymous function
// if i use go func(...) it does not return.
// even if i wait with time.Sleep(..)
// go func doesn't return, why?
	func(ch chan int) {
		for i := 0; i < cap(ch); i++ {
			fmt.Println(<-ch)
		}
	}(ch)
}
