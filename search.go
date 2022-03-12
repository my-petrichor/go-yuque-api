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
		"topic"
		"repo"
		"doc"
		"artboard"
		"group"
	*/
	Kind string

	// TODO Search path, The path on the team or knowledge base URL
	// Scope string
}

func newSearcher(c *Client) *Searcher {
	return &Searcher{
		c,
	}
}

// Get get user info
func (s *Searcher) Work(keyword string, options ...SearcherOption) (*SearchSerializer, error) {
	var opt SearcherOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	if opt.Offset == 0 {
		opt.Offset = 1
	}
	if opt.Kind == "" {
		opt.Kind = "doc"
	}

	var (
		url    = fmt.Sprintf(s.baseURL + internal.SearchPath)
		search = SearchSerializer{}
		body   = struct {
			Type   string `json:"type"`
			Q      string `json:"q"`
			Offset int    `json:"offset"`
		}{
			Type:   opt.Kind,
			Q:      keyword,
			Offset: opt.Offset,
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
