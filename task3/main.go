package main

import "fmt"

func main() {
	arr := []int64{2, 4, 6, 8, 10}
	ch := make(chan int64, len(arr))
	var result int64
	result = 0
	for _, value := range arr {
		go func(value int64, ch chan<- int64) {
			ch <- value * value
		}(value, ch)
	}
	for i := 0; i < len(arr); i++ {
		result += <-ch

	}
	fmt.Println(result)
}
