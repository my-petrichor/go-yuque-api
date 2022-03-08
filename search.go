package yuque

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type Search struct {
	*client
}

type SearchOption struct {
	// split page example: 1, 2, 3...
	Offset int
	/*
		Search path, The path on the team or knowledge base URL
		For example, "scope = lark" is search for the content of the lark group,
		"scope = lark/help" is search for the content under the "lark/help" knowledge base
	*/
	Scope string
	// true is search with yourself
	Related bool
}

func newSearch(c *client) *Search {
	return &Search{
		c,
	}
}

// Get get user info
/*
 kind
 "topic"
 "repo"
 "doc"
 "artboard"
 "group"
 "user"
 "attachment"
*/
func (s *Search) Start(kind, keyword string, options ...SearchOption) (*internal.ResponseUserDetailSerializer, error) {
	var option SearchOption
	if len(options) > 1 {
		return nil, errors.New("options length more than one")
	} else if len(options) == 1 {
		option = options[0]
	}

	var (
		url  = fmt.Sprintf(s.BaseURL + internal.SearchPath)
		user = internal.ResponseUserDetailSerializer{}
		body = struct {
			Type    string `json:"type"`
			Q       string `json:"q"`
			Scope   string `json:"scope"`
			Offset  int    `json:"offset"`
			Related bool   `json:"related"`
		}{
			Type:    kind,
			Q:       keyword,
			Scope:   option.Scope,
			Offset:  option.Offset,
			Related: option.Related,
		}
	)

	respBody, err := s.do(url, Option{Body: body})
	if err != nil {
		return nil, err
	}

	fmt.Println(string(respBody), "----------")
	err = json.Unmarshal(respBody, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
