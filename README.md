# go-yuque-api

This library is designed as a simple wrapper around the YuQue API. It's encouraged to read YuQue's docs first to get a simple understanding.

## introduction

| Name      | Description                                                     | Example                                              |
| --------- | --------------------------------------------------------------- | ---------------------------------------------------- |
| id        | 数据的唯一编号/主键                                             | 1984                                                 |
| login     | 用户／团队的唯一名称用户／团队编号                              | 用户：用户个人路径团队：如语雀团队，login 值为 yuque |
| book_slug | 仓库唯一名称                                                    | 如语雀开发者文档这个仓库，book_slug 值为 developer   |
| namespace | 仓库的唯一名称需要组合 :login/:book_slug 或可以直接使用仓库编号 | yuque/developer                                      |
| slug      | 文档唯一名称                                                    | 如语雀官方开发者文档的 slug 值为 api                 |

## notice

- 匿名请求，IP 限制, 200 次/小时
- 传递 Token 的情况下，每个用户（基于 Token 关联到的账户），5000 次/小时

## examples

```
package main

import (
	"fmt"

	"github.com/my-Sakura/go-yuque-api/api"
)

func main() {
	token := "[token]"
	client := yuque.NewClient(token)
	user, err := client.User.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
```

See [examples](https://github.com/my-Sakura/go-yuque-api/tree/main/examples) folder for more detailed examples.

## license

[MIT](https://github.com/my-Sakura/go-yuque-api/blob/main/LICENSE) © my-Sakura
