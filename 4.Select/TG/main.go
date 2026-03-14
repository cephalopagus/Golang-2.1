package main

import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	messCh1 := make(chan Message)
	messCh2 := make(chan Message)

	go func() {
		for {
			messCh1 <- Message{
				Author: "Vika",
				Text:   "Привет",
			}
			time.Sleep(2 * time.Second)
		}

	}()
	go func() {
		for {
			messCh2 <- Message{
				Author: "Isma",
				Text:   "Салам",
			}
			time.Sleep(500 * time.Millisecond)
		}

	}()

	for {
		select {
		case msg1 := <-messCh1:
			fmt.Printf("Author - %s, text - %s\n", msg1.Author, msg1.Text)
		case msg2 := <-messCh2:
			fmt.Printf("Author - %s, text - %s\n", msg2.Author, msg2.Text)
		}
	}

}
