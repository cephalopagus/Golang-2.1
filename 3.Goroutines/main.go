package main

import (
	"fmt"
	"time"
)

func main() {
	coal := 0

	transferPoint := make(chan int)

	tm := time.Now()
	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint

	fmt.Println(coal, time.Since(tm))
}

func mine(transferPoint chan int, n int) {
	fmt.Println("Поход начался - ", n)
	time.Sleep(time.Second * 1)
	fmt.Println("Поход закончился - ", n)
	transferPoint <- 10
}
