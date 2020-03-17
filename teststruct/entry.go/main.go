package main

import (
	"fmt"
	"learngo/teststruct"
)

// type Triangle interface {
//  Area() int
//  Girth() int
// }
type MyTriangle struct {
	*teststruct.Triangle
	Flag bool
}

func (mt *MyTriangle) print() {
	fmt.Println(mt.Flag)
}
func main() {
	// var a Triangle
	a := MyTriangle{&teststruct.Triangle{3, 4, 5}, true}
	fmt.Println(a.Triangle)
	fmt.Println(a.Area())
	fmt.Println(a.Girth())
	a.print()
}
