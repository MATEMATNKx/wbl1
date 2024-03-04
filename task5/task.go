package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	N = 2
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала
 и выводят в stdout.
Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.

*/

func main() {
	ch := make(chan string)

	//Создаю данные и помещаю их в канал

	for i := 0; i < N; i++ {
		go func(i int) {
			for value := range ch {
				fmt.Printf("[%d], %s \n", i, value)
			}
		}(i)
	}
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
