package main

import (
	"fmt"

	"github.com/my-Sakura/go-yuque-api/api"
)

func main() {
	token := "[token]"
	login := "[login]"
	books := api.GetUserBookList(token, login)
	for _, book := range books.Data {
		fmt.Printf("bookSlug: %s, bookName: %s\n", book.Slug, book.Name)
	}
}
