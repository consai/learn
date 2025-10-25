package main

import (
	"fmt"
	"time"
)

func main() {
	chan10()

	chan100()
	time.Sleep(100)

}

func chan10() {

	ch := make(chan int)
	go send(ch)
	go recv(ch)

}
func send(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func recv(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func chan100() {
	ch := make(chan int, 10)
	go send100(ch)
	go recv100(ch)
}
func send100(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func recv100(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
