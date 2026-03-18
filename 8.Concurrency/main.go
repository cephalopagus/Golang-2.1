package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64
	var mails []string
	mtx := sync.Mutex{}

	minerCtx, minerCncl := context.WithCancel(context.Background())
	postmanCtx, postmanCncl := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	go func() {
		time.Sleep(time.Second * 10)
		minerCncl()
	}()
	go func() {
		time.Sleep(time.Second * 10)
		postmanCncl()
	}()

	coalTransferPoint := miner.MinerPool(minerCtx, 200)
	mainTransferPoint := postman.PostmanPool(postmanCtx, 300)

	initTime := time.Now()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mainTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()
	wg.Wait()

	// isCoalCloased := false
	// isMailCloased := false

	// for !isCoalCloased || !isMailCloased {
	// 	select {
	// 	case c, ok := <-coalTransferPoint:
	// 		if !ok {
	// 			isCoalCloased = true
	// 			continue
	// 		}
	// 		coal += c

	// 	case m, ok := <-mainTransferPoint:
	// 		if !ok {
	// 			isMailCloased = true
	// 			continue
	// 		}

	// 		mails = append(mails, m)
	// 	}
	// }

	fmt.Println("coal:", coal.Load())
	mtx.Lock()
	fmt.Println("mail:", len(mails))

	mtx.Unlock()

	fmt.Println(time.Since(initTime))
}
