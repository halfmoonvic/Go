package main

import (
	"fmt"
	"time"
)

type atomicInt int

func (a *atomicInt) incremennt() {
	*a++
}

func (a *atomicInt) get() int {
	return int(*a)
}

func main() {
	var a atomicInt
	a.incremennt()
	go func() {
		a.incremennt()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
