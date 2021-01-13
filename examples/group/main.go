package main

import (
	"fmt"

	"github.com/my-Sakura/go-yuque-api/api"
)

func main() {
	token := "[token]"
	login := "[login]"

	groups := api.GetGroupList(token, login)
	for _, group := range groups.Data {
		fmt.Printf("group_login: %s, group_name: %s\n", group.Login, group.Name)
	}
}
