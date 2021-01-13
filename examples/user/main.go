package main

import (
	"fmt"

	"github.com/my-Sakura/go-yuque-api/api"
)

func main() {
	token := "[token]"
	user := api.GetUserInfo(token)
	fmt.Printf("Name: %s, user_id: %d, login: %s\n", user.Data.Name, user.Data.Id, user.Data.Login)
}
