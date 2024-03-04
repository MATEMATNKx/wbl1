package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//Создаю данные и помещаю их в канал

	for i := 0; i < N; i++ {
		go func(i int) {
			for value := range ch {
				fmt.Printf("[%d], %s \n", i, value)
			}
		}(i)
	}
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-interrupt
		fmt.Printf("Ctrl-C was pressed")
	}()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- "Some text: " + strconv.Itoa(i)
			time.Sleep(1 * time.Second)

		}

	}

}
