# go-yuque-api

Golang version of yuque api.

## introduction

| Name      | Description          | Example                                                                               |
| --------- | -------------------- | ------------------------------------------------------------------------------------- |
| login     | 用户／团队的唯一标识 | 用户：用户个人路径 团队：如[语雀团队](https://www.yuque.com/yuque/)，login 值为 yuque |
| repo_slug | 仓库唯一标识         | 如[语雀开发者文档](https://www.yuque.com/yuque/developer) repo_slug 值为 developer    |
| namespace | :login/:repo_slug    | yuque/developer                                                                       |
| slug      | 文档唯一名称         | 如[语雀开发者文档](https://www.yuque.com/yuque/developer/api)的 slug 值为 api         |

## notice

- 匿名请求，IP 限制, 200 次/小时
- 传递 Token 的情况下，每个用户（基于 Token 关联到的账户），5000 次/小时
- 对于文档更改接口，如果想要对手动更新过的文档进行更改，需要设置 ForceASL option 为 1 进行强制更改，
  因为手动更改过文档会将文档格式更改为 lake，而语雀接口现在只支持 markdown 类型的文档更改。

## examples

```
package main

import (
	"fmt"

    yuque "github.com/my-Sakura/go-yuque-api"
)

func main() {
	token := "[token]"
	client := yuque.NewClient(token)
	user, err := client.User.GetInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
```

See [examples](https://github.com/my-Sakura/go-yuque-api/tree/main/examples) folder for more detailed examples.

## license

[MIT](https://github.com/my-Sakura/go-yuque-api/blob/main/LICENSE) © my-Sakura
