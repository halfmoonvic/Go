package main

import (
	"Learngo/test"
	"fmt"
)

type retrieve interface {
	Get(string) string
}

func getRetriever() retrieve {
	return test.Retriever{}
}

func main() {
	var retriever retrieve
	retriever = getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com"))
}
