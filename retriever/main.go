package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func downloda(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

func session(s RetrieverPoster) string {
	// s.Get(url)
	s.Post(url, map[string]string{"contents": "another faked imooc.com"})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

func main() {
	var r Retriever
	// r = &real.Retriever{
	// 	UserAgent: "Mozilla/5.0",
	// 	TimeOut:   time.Minute,
	// }
	// inspect(r)
	// realRetriever := r.(*real.Retriever)
	// fmt.Println(realRetriever.TimeOut)

	r = &mock.Retriever{"this is a fake imooxc.com"}
	// fmt.Println(downloda(r))
	inspect(r)

	// retriever := &mock.Retriever{"this is a fake imooc.com"}
	// fmt.Println(session(retriever))
}
