package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this is a fake imooc.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
