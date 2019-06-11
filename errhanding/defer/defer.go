package main

import (
	"LearnGo/functional/fib"
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
}

func writeFile(filename string) {
	// file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
		// panic(err)
	}
	// 打开文件后 立即 defer 关闭，这样即便在后面程序当中出现错误，file.Close() 也能够执行
	defer file.Close()

	// 先将内容写到 buf 内存当中
	writer := bufio.NewWriter(file)
	// 将内存中的内容写到文件当中
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	writeFile("fib.txt")
}
