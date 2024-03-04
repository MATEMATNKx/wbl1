package main

import (
	"fmt"
	"strconv"
	"time"
)

// go mod init main
func main() {
	ch := make(chan string)
	//Горутина на чтение
	go func(ch <-chan string) {
		for value := range ch {
			fmt.Println(value)
		}
	}(ch)

	//Горутина на запись
	go func(ch chan<- string) {
		for i := 0; ; i++ {

			ch <- fmt.Sprintf("Some text: %d", i)
		}
	}(ch)
	stop := time.After(1 * time.Second)
	for i := 0; ; i++ {
		select {
		case <-stop:
			return
		default:
			ch <- "Some text: " + strconv.Itoa(i)
		}
	}
}
