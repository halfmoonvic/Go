package main

import (
	"fmt"
	"time"
)

//反例,直接报错
func chan0() {
	c := make(chan int)
	c <- 1
	c <- 2
	// for i := 0; i < 100; i++ {
	// 	c <- i //阻塞了main的goroutine
	// }
	// go func() {
	// for data := range c {
	// fmt.Println(data) }
	// }()
	go func() {
		for {
			// n := <-c
			// fmt.Println(n)
			fmt.Printf("worder %d\n", <-c)
		}
	}()
}

/**
  必须先启动goroutine接收,才能进行发送,否则报错,
  除非发送也是新启动了一个goroutine.
*/
func chan1() {
	c := make(chan int)
	go func() { //先启动一个新的goroutine接收
		for data := range c {
			fmt.Println(data)
		}
	}()
	//main 中发送,不报错
	for i := 0; i < 100; i++ {
		c <- i
	}
	//上面新的go func... 放到这里就报错.和chan0()函数一样
	//go func() {
	//  for   {
	//    fmt.Println(<-c)
	//  }
	//
	//}()
	time.Sleep(time.Millisecond * 10)
}

/**
  而如果在main中定义一个chan,然后新启动一个发送者的goroutine,那么就不报错,即使没有接收者
*/
func chan2() {
	c := make(chan int)
	go func() { //可以在chan定义的goroutine中发送数据,也可以定义新的goroutine发送数据
		fmt.Println("发送者", time.Now())
		for i := 0; i < 100; i++ {
			fmt.Println("发送数据:", i)
			c <- i
		}
		fmt.Println("关闭chan")
		close(c) //关闭chan,如此,接受者就会感应到,从而接收者会退出接收chan的循环
	}()
	time.Sleep(time.Second * 10)
	//go func() { //接收goroutine中的数据必须在一个新的goroutine中接收数据
	//  fmt.Println("接受者",time.Now())
	//  for data := range c {
	//    fmt.Println(data)
	//  }
	//
	//}()
	time.Sleep(time.Millisecond * 10)
}

/**
  在非main的goroutine中定义的chan,即使没有接收者,即使定以后不再启动一个新的goroutine也不报错
  **老师: 为什么在main中定义的chan,不启动一个新的goroutine就马上报错,而不再main的goroutine中定义,就不报错?**
*/
func chan3() {
	go func() {
		c := make(chan int)
		fmt.Println("发送数据:", 1)
		c <- 1
	}()
	time.Sleep(time.Second * 10)
}
func main() {
	fmt.Println("主程序", time.Now())
	chan1()
	//chan2()
	// chan3()
	time.Sleep(time.Millisecond * 10)
}
