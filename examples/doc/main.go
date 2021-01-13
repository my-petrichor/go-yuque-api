package main

import (
	"fmt"

	"github.com/my-Sakura/go-yuque-api/api"
)

func main() {
	token := "YLN7hYz4iKmWSs1MfyLDrNY2IqZaM2ZabOOmpIAX"
	namespace := "my-sakura/doc"

	docs := api.GetDocumentList(token, namespace)
	for _, doc := range docs.Data {
		fmt.Printf("slug: %s, title: %s\n", doc.Slug, doc.Title)
	}
}
