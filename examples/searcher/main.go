package main

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
)

const (
	token = "[token]"
)

func main() {
	client := yuque.NewClient(token)
	search, err := client.Search.Work("keyword", yuque.SearcherOption{Kind: "doc", Offset: 10})
	if err != nil {
		panic(err)
	}

	fmt.Println(search)
}
