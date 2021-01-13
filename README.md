# go-yuque-api
a SDK of yuque API written in go

## examples
```
func main() {
	token := "[token]"
	user := api.GetUserInfo(token)
	fmt.Printf("Name: %s, user_id: %d, login: %s\n", user.Data.Name, user.Data.Id, user.Data.Login)
}
```
See [examples](https://github.com/my-Sakura/go-yuque-api/tree/main/examples) folder for more detailed examples.