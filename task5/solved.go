package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

// go mod init main
// Строка - не изменяемый тип. Любая работа с ней - новое выделение памяти
func main() {
	s := ""
	start := time.Now()
	for i := 0; i < 100000; i++ {
		s += strconv.Itoa(i)
	}
	fmt.Println(time.Since(start))

	var b bytes.Buffer
	start = time.Now()
	for i := 0; i < 100000; i++ {
		b.WriteString(strconv.Itoa(i))
	}
	fmt.Println(time.Since(start))

}
