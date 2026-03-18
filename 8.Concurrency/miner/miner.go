package miner

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func miner(ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int) {

	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтер номер:", n, "Мой рабочий день закончился!")
			return
		default:
			fmt.Println("Я шахтер номер:", n, "Начал добывать уголь!")
			time.Sleep(time.Second * 1)
			fmt.Println("Я шахтер номер:", n, "Добыл уголь!")

			transferPoint <- power
			fmt.Println("Я шахтер номер:", n, "Передал уголь:", power)
		}

	}

}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	transferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(ctx, wg, transferPoint, i, rand.Intn(10))
	}

	go func() {
		wg.Wait()
		close(transferPoint)
	}()

	return transferPoint
}
