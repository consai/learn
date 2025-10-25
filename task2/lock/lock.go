package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var mutex sync.Mutex

func main() {

	test_lock()
	test_atomic()
}

func test_lock() {
	num := 0
	for i := 0; i < 9; i++ {
		go lockadd(&num)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(num)
}

func lockadd(j *int) {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		*j++
		mutex.Unlock()
	}
}

func test_atomic() {
	var num int32 = 0
	for i := 0; i < 9; i++ {
		go atomicadd(&num)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(num)
}

func atomicadd(j *int32) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(j, 1)
	}
}
