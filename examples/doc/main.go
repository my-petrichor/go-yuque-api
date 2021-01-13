package main

import (
	"fmt"

	"github.com/my-Sakura/go-yuque-api/api"
)

func main() {
	token := "[token]"
	namespace := "[login]/[repo_slug]"

	docs := api.GetDocumentList(token, namespace)
	for _, doc := range docs.Data {
		fmt.Printf("slug: %s, title: %s\n", doc.Slug, doc.Title)
	}
}
