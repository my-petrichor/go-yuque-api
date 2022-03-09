package main

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-api/internal"
)

var (
	token      = "[token]"
	groupLogin = "[groupLogin]"
	login      = "[login]"
)

func main() {
	client := yuque.NewClient(token)
	group, err := ListAll(client)
	if err != nil {
		panic(err)
	}

	fmt.Println(group)
}

func ListAll(client *yuque.Client) (*internal.ResponseUserSerializer, error) {
	return client.Group.ListAll(login)
}

func ListPublic(client *yuque.Client) (*internal.ResponseUserSerializer, error) {
	return client.Group.ListPublic()
}

func GetInfo(client *yuque.Client) (*internal.ResponseUserDetailSerializer, error) {
	return client.Group.GetInfo(groupLogin)
}

func GetMembers(client *yuque.Client) (*internal.ResponseGroupUserSerializer, error) {
	return client.Group.GetMembers(groupLogin)
}

func Create(client *yuque.Client) (*internal.ResponseUserDetailSerializer, error) {
	return client.Group.Create("newGroupName", "newGroupLogin", yuque.GroupOption{
		Login:       "newGroupLogin",
		Name:        "newGroupName",
		Description: "newGroupDescription",
	})
}

func Update(client *yuque.Client) (*internal.ResponseUserDetailSerializer, error) {
	return client.Group.Update(groupLogin, yuque.GroupOption{
		Login:       "newqwe",
		Name:        "newGroupName",
		Description: "newGroupDescription",
	})
}

func UpdateMember(client *yuque.Client) (*internal.ResponseGroupUserDetailSerializer, error) {
	return client.Group.UpdateMember(groupLogin, login, 0)
}

func Delete(client *yuque.Client) (*internal.ResponseUserDetailSerializer, error) {
	return client.Group.Delete(groupLogin)
}

func DeleteMember(client *yuque.Client) (*internal.ResponseGroupUserDetailSerializer, error) {
	return client.Group.DeleteMember(groupLogin, login)
}
