package main

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
)

func main() {
	token := "[token]"
	client := yuque.NewClient(token)
	repo, err := client.Repo.ListAllUnderUser("my-sakura")
	if err != nil {
		panic(err)
	}

	fmt.Println(repo)
}
