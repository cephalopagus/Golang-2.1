package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	chInt := make(chan int)
	chString := make(chan string)

	go func() {
		i := 1
		for {
			chInt <- i
			i++
			time.Sleep(time.Second * 1)
		}

	}()

	go func() {
		i := 1
		for {
			chString <- "num - " + strconv.Itoa(i)
			i++
			time.Sleep(time.Millisecond * 500)
		}

	}()
	for {
		select {
		case num := <-chInt:
			fmt.Println(num)
		case str := <-chString:
			fmt.Println(str)
		}
	}

}
