package main

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
)

const (
	token     = "[token]"
	namespace = "[namespace]"
	slug      = "[slug]"
)

func main() {
	client, err := yuque.NewClient("[token]")
	if err != nil {
		panic(err)
	}
	document, err := Create(client)
	if err != nil {
		panic(err)
	}

	fmt.Println(document)
}

func ListAll(client *yuque.Client) (*yuque.ResponseDocSerializer, error) {
	return client.Document.ListAll(namespace)
}

func GetInfo(client *yuque.Client) (*yuque.ResponseDocDetailSerializer, error) {
	return client.Document.GetInfo(namespace, slug)
}

func Create(client *yuque.Client) (*yuque.ResponseDocDetailSerializer, error) {
	return client.Document.Create(namespace, yuque.DocumentOption{
		Slug:   "slug",
		Title:  "Title",
		Body:   "Body",
		Format: "markdown",
		Public: 0,
	})
}

func Update(client *yuque.Client) (*yuque.ResponseDocDetailSerializer, error) {
	return client.Document.Update(namespace, slug, yuque.DocumentOption{
		Slug:     "newSlug",
		Title:    "newTitle",
		Body:     "newBody",
		Public:   0,
		ForceASL: 1,
	})
}

func Delete(client *yuque.Client) (*yuque.ResponseDocDetailSerializer, error) {
	return client.Document.Delete(namespace, slug)
}
