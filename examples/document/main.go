package main

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
)

func main() {
	token := "[token]"
	client := yuque.NewClient(token)
	document, err := client.Document.ListAll("my-sakura/blog")
	if err != nil {
		panic(err)
	}

	fmt.Println(document)
}
