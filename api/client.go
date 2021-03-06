package api

import "net/http"

type Client struct {
	HTTPClient *http.Client
	Token      string
	Namespace  string
	Slug       string
	Login      string
}

func NewClient(token string) *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		Token:      token,
	}
}
