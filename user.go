package yuque

import (
	"encoding/json"
	"fmt"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type User struct {
	*client
}

func newUser(c *client) *User {
	return &User{
		c,
	}
}

// Get get user info
func (u *User) Get() (*internal.ResponseUserDetailSerializer, error) {
	var (
		url  = fmt.Sprintf(u.BaseURL + internal.UserGetPath)
		user = internal.ResponseUserDetailSerializer{}
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
