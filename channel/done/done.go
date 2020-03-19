package main

import (
	"fmt"
)

func doWorker(id int, c chan int, done chan bool) {
	// for 循环会无限的 接受 channel 中的数据，只要 向 channel 发送了数据就会 进行接受
	// for {
	// 	fmt.Printf("worker %d received %c\n", id, <-c)
	// }
	// range 则是 channel 中有多少数据，就接受多少数据
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		go func() {
			done <- true
		}()
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	// var c chan int // c == nil
	// fmt.Println(c)
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
		// go doWorker(i, workers[i])
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		// <-worker.done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		// <-worker.done
	}
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}

	// time.Sleep(time.Microsecond)
}

// func closeChannel() {
// 	ch := make(chan int)
// 	go doWorker(0, ch)
// 	ch <- 'a'
// 	ch <- 'b'
// 	ch <- 'c'
// 	close(ch)
// 	time.Sleep(time.Millisecond)
// }

func main() {
	chanDemo()
}
