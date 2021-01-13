# go-yuque-api
This library is designed as a simple wrapper around the YuQue API. It's encouraged to read YuQue's docs first to get a simple understanding. 

## examples
```
package main

import (
	"fmt"

	"github.com/my-Sakura/go-yuque-api/api"
)

func main() {
	token := "[token]"
	user := api.GetUserInfo(token)
	fmt.Printf("Name: %s, user_id: %d, login: %s\n", user.Data.Name, user.Data.Id, user.Data.Login)
}
```
See [examples](https://github.com/my-Sakura/go-yuque-api/tree/main/examples) folder for more detailed examples.

## license
[MIT](https://github.com/my-Sakura/go-yuque-api/blob/main/LICENSE) © my-Sakura

