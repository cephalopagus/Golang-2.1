package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	transferPoint := make(chan int)

	go func() {
		iteration := 3 + rand.Intn(4)
		fmt.Println("Iteration:", iteration)
		for i := 0; i < iteration; i++ {
			transferPoint <- 10
			time.Sleep(time.Millisecond * 1000)
		}
		close(transferPoint)
	}()
	coal := 0
	// for {
	// 	v, ok := <-transferPoint
	// 	if !ok {
	// 		fmt.Println("FINISH")
	// 		break
	// 	}
	// 	coal += v
	// 	fmt.Println("Coal:", coal)

	// }
	for i := range transferPoint {
		coal += i
		fmt.Println("Coal:", coal)
	}
	fmt.Println("All coals:", coal)
}
