package main

import (
	"fmt"
	"sync"
	"time"
)

var likes int = 0
var rmtx sync.RWMutex

func setLike(wg *sync.WaitGroup) {
	for i := 0; i < 100_000; i++ {
		rmtx.Lock()
		likes++
		rmtx.Unlock()
	}
	defer wg.Done()
}

func getLike(wg *sync.WaitGroup) {
	for i := 0; i < 100_000; i++ {
		rmtx.RLock()
		_ = likes
		rmtx.RUnlock()
	}
	defer wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	initTime := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go setLike(wg)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go getLike(wg)
	}

	wg.Wait()
	fmt.Println(time.Since(initTime))
}
