package main

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
)

func main() {
	token := "YLN7hYz4iKmWSs1MfyLDrNY2IqZaM2ZabOOmpIAX"
	client := yuque.NewClient(token)
	search, err := client.Search.Start("doc", "history", yuque.SearchOption{Offset: 10})
	if err != nil {
		panic(err)
	}

	fmt.Println(search)
}
