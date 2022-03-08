package main

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
)

func main() {
	token := "[token]"
	client := yuque.NewClient(token)
	user, err := client.User.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
