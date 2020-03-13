package main

import (
	"LearnGo/queue"
	"fmt"
)

func main() {

	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	// fmt.Println(q.Pop())
	// fmt.Println(q.Pop())

	// fmt.Println(len(q))
	// fmt.Println(q.IsEmpty())

	// q.Push("abc")
	// q.Push(false)
	q.Pop()
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q)
	fmt.Println(q.IsEmpty())

}
