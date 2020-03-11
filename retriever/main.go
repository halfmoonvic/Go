package main

import (
	"fmt"
	"learngo/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// func download(r Retriever) string {
// 	return r.Get(url)
// }

// func post(poster Poster) {
// 	poster.Post(url,
// 		map[string]string{
// 			"name":   "ccmouse",
// 			"course": "golang",
// 		})
// }

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post("http://www.imooc.com", map[string]string{
		"contents": "anoher faked imoock.com",
	})
	return s.Get("http://www.imooc.com")
}

// func inspect(r Retriever) {
// 	fmt.Printf("%T %v\n", r, r)
// 	switch v := r.(type) {
// 	case mock.Retriever:
// 		fmt.Println("Contents", v, v.Contents)
// 	case *real.Retriever:
// 		fmt.Println("UserAgent:", v, v.UserAgent)
// 	}
// }

func main() {
	var r RetrieverPoster
	r = mock.Retriever{"hahah"}
	fmt.Println(session(&r))

	mockRetriever := mock.Retriever{Contents: "this is a fake imooc.com"}
	fmt.Println(session(&mockRetriever))
	// contents := r.(mock.Retriever)
	// fmt.Println(contents)

	// inspect(r)
	// r = &real.Retriever{
	// 	UserAgent: "Mozilla/5.0",
	// 	TimeOut:   time.Minute,
	// }

	// inspect(r)

	// fmt.Println(download(r))

	// if mockRetriever, ok := r.(mock.Retriever); ok {
	// 	fmt.Println(mockRetriever, mockRetriever.Contents)
	// } else {
	// 	fmt.Println("not a mock retrieve")
	// }

}
