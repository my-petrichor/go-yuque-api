package yuque

import (
	"encoding/json"
	"fmt"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type Searcher struct {
	*Client
}

type SearcherOption struct {
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

func newSearch(c *Client) *Searcher {
	return &Searcher{
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
func (s *Searcher) Work(kind, keyword string, options ...SearcherOption) (*internal.SearchSerializer, error) {
	var opt SearcherOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	if opt.Offset == 0 {
		opt.Offset = 1
	}
	var (
		url    = fmt.Sprintf(s.BaseURL + internal.SearchPath)
		search = internal.SearchSerializer{}
		body   = struct {
			Type    string `json:"type"`
			Q       string `json:"q"`
			Scope   string `json:"scope"`
			Offset  int    `json:"offset"`
			Related bool   `json:"related"`
		}{
			Type:    kind,
			Q:       keyword,
			Scope:   opt.Scope,
			Offset:  opt.Offset,
			Related: opt.Related,
		}
	)

	respBody, err := s.do(url, option{Body: body})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &search)
	if err != nil {
		return nil, err
	}

	return &search, nil
}
