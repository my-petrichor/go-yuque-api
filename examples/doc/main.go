package main

import (
	"fmt"

	"github.com/my-Sakura/go-yuque-api/api"
)

func main() {
	token := "[token]"
	namespace := "[login]/[repo_slug]"
	oldSlug := "[old_document_slug]"

	c := api.NewClient(token)
	id := c.GetDocumentID(namespace, oldSlug)
	newSlug := "[new_document_slug]"
	// change oldSlug to newSlug
	s := c.ChangeDocumentSlug(id, newSlug)
	fmt.Println(id, s)
}
