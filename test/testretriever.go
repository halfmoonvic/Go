package test

import "fmt"

type Retriever struct{}

func (Retriever) Get(url string) string {
	fmt.Println(url)
	return "fake retrieve"
}
