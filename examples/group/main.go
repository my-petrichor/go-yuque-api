package main

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
)

func main() {
	token := "[token]"
	client := yuque.NewClient(token)
	group, err := client.Group.ListAll("my-sakura")
	if err != nil {
		panic(err)
	}

	fmt.Println(group)
}
