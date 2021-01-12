package main

import (
	"fmt"

	"github.com/my-Sakura/vercel-test/api"
)

func main() {
	token := "YLN7hYz4iKmWSs1MfyLDrNY2IqZaM2ZabOOmpIAX"
	doc := api.UpdateDocument(token, "my-sakura/miniproject", "30043054", "asdasdasdfadfasf")
	//api.CreateDocument(token, "hello", "world")
	//doc := api.GetDocumentInfo(token, "my-sakura/miniproject", "hello")
	fmt.Println(doc.Data.Id, doc.Data.Title)
}
