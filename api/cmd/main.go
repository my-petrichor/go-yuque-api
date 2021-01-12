package main

import (
	"fmt"

	"github.com/my-Sakura/go-yuque-api/api"
)

func main() {
	token := "YLN7hYz4iKmWSs1MfyLDrNY2IqZaM2ZabOOmpIAX"
	namespace := "my-sakura/doc"
	// login := "my-sakura"
	//doc := api.UpdateDocument(token, "my-sakura/miniproject", "30043054", "asdasdasdfadfasf")
	//api.GetDocumentList(token, namespace)
	//api.CreateDocument(token, "hello", "world")
	//doc := api.GetDocumentInfo(token, "my-sakura/miniproject", "hello")
	//fmt.Println(doc.Data.Id, doc.Data.Title)
	//doc := api.GetDocumentList(token, namespace)
	// for i, docSerializer := range doc.Data {
	// 	fmt.Println(i, docSerializer.Id, docSerializer.Slug, docSerializer.Title)
	// }
	//doc := api.DeleteDocument(token, namespace, 30032649)
	//fmt.Println(doc.Data.Id, doc.Data.Slug, doc.Data.Body, doc.Data.Title, doc.Data.Format)
	// user := api.GetUserInfo(token)
	// fmt.Println(user.Data.Id, user.Data.Login, user.Data.Books_count)
	// book := api.GetGroupBookList(token, login)
	// for _, v := range book.Data {
	// 	fmt.Println(v.Name)
	// }
	// book := api.CreateGrouprBook(token, login, "te", "dads", 1)
	// fmt.Println(book.Data.Name, book.Data.Public)

	// book := api.GetBookInfo(token, namespace)
	// fmt.Println(book.Data.Name, book.Abilities.Doc.Create)
	// book := api.UpdateBook(token, namespace, "newname", "newslug")
	// fmt.Println(book.Abilities.Update, book.Data.Name)
	// book := api.DeleteBook(token, namespace)
	// fmt.Println(book.Data.Name)
	book := api.GetBookDirectoryStructure(token, namespace)
	for i, v := range book.Data {
		fmt.Println(i, v.Id)
	}
}
