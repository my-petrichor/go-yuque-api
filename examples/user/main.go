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
	user, err := client.User.GetInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
