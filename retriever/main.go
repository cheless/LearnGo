package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)
// Only methods are specified in the interface
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}
// combination of 2 interfaces
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
			"name": "ccmouse",
			"course": "golang",
		})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url) // Implement Get() & Post() at the same time
}
// switch assertion
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r) // T is Type, v is value
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
/*
The specific file interface implemented is only
specified in the main()
 */

func main() {
	var r Retriever // r is a Retriever interface

	// The mock.Retriever object is copied into r
	r = &mock.Retriever{"This is a fake imooc.com"} // "This..." is the Contents
	inspect(r)

	// r points to the object of real.Retriever
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	}else {
		fmt.Println("Not a mock type")
	}

	var r1 = &mock.Retriever{}
	fmt.Println("Try a session:")
	fmt.Println(session(r1))

}
