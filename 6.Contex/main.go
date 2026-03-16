package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())

	childContext, childCancel := context.WithCancel(parentContext)

	go foo(parentContext)

	go boo(childContext)

	time.Sleep(time.Second * 1)
	parentCancel()

	time.Sleep(time.Second * 1)
	childCancel()

	time.Sleep(time.Second * 2)

}

func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("FOOOOO IS OVER(")
			return
		default:
			fmt.Println("FOOOOO is working")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
func boo(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("BOOO IS OVER(")
			return
		default:
			fmt.Println("BOOO is working")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
