package yuque

import (
	"encoding/json"
	"fmt"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type User struct {
	*Client
}

func newUser(c *Client) *User {
	return &User{
		c,
	}
}

// Get get user info
func (u *User) GetInfo() (*ResponseUserDetailSerializer, error) {
	var (
		url  = fmt.Sprintf(u.baseURL + internal.UserGetPath)
		user = ResponseUserDetailSerializer{}
	)

	respBody, err := u.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
