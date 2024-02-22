package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

var m sync.Mutex

func main() {
	var cm sync.Mutex
	cm = m  //copy of the mutex

	go increment("Goroutine 1", &m)
	go increment("Goroutine 2", &cm)

	time.Sleep(time.Second)
}

func increment(name string, m *sync.Mutex){
	for i:= 1; i <= 3; i++{
		m.Lock()
		counter += 1
		fmt.Println(name + ": ", counter)
		m.Unlock()
	}
}
