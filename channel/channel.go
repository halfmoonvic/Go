package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	// done chan bool
	// wg *sync.WaitGroup
	done func()
}

func doWorker(id int, c chan int, w worker) {
	for n := range c {
		// n, ok := <-c
		// if !ok {
		// 	break
		// }
		fmt.Printf("worker %d received %c\n", id, n)
		w.done()
		// go func() {
		// 	done <- true
		// }()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		// done: make(chan bool),
		// wg: wg,
		done: func() {
			wg.Done()
		},
	}

	go doWorker(id, w.in, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
		// <-workers[i].done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		// <-workers[i].done
	}

	// for _, worker := range workers {
	// 	<-worker.done
	// 	// <-worker.done
	// }
	wg.Wait()

	// time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	// bufferedChannel()
	// channelClose()
}
