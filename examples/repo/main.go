package main

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-api/internal"
)

const (
	token      = "[token]"
	login      = "[login]"
	groupLogin = "[groupLogin]"
	namespace  = "[namespace]"
)

func main() {
	client := yuque.NewClient(token)
	repo, err := ListAllUnderUser(client)
	if err != nil {
		panic(err)
	}

	fmt.Println(repo)
}

func ListAllUnderUser(client *yuque.Client) (*yuque.ResponseBookSerializer, error) {
	return client.Repo.ListAllUnderUser(login)
}

func ListAllUnderGroup(client *yuque.Client) (*yuque.ResponseBookSerializer, error) {
	return client.Repo.ListAllUnderGroup(groupLogin)
}

func GetInfo(client *yuque.Client) (*yuque.ResponseBookDetailSerializer, error) {
	return client.Repo.GetInfo(namespace)
}

func GetDir(client *yuque.Client) (*yuque.ResponseBookDirectoryStructure, error) {
	return client.Repo.GetDir(namespace)
}

func CreateUnderUser(client *yuque.Client) (*yuque.ResponseBookDetailSerializer, error) {
	return client.Repo.CreateUnderUser(login, "newslugq", "Book", "newRepoName", yuque.RepoOption{
		Description: "this is a test",
	})
}

func CreateUnderGroup(client *yuque.Client) (*internal.ResponseBookDetailSerializer, error) {
	return client.Repo.CreateUnderGroup(groupLogin, "newslug", "Book", "newRepoName", yuque.RepoOption{
		Description: "this is a test",
	})
}

func Update(client *yuque.Client) (*internal.ResponseBookDetailSerializer, error) {
	return client.Repo.Update(namespace, yuque.RepoOption{
		Slug:        "aaa",
		Name:        "newName",
		Description: "newDescription",
		Kind:        "Book",
		Public:      4,
	})
}

func Delete(client *yuque.Client) (*internal.ResponseBookDetailSerializer, error) {
	return client.Repo.Delete(namespace)
}
