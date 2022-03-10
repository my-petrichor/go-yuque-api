package yuque

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type Client struct {
	Token string
	// default is "https://www.yuque.com/api/v2"
	BaseURL string

	// http request User-Agent
	UserAgent string

	User     *User
	Document *Document
	Repo     *Repo
	Group    *Group
	Search   *Searcher
}

type clientOption func(*Client)

func NewClient(token string, options ...clientOption) *Client {
	client := &Client{
		Token: token,
	}

	client.User = newUser(client)
	client.Document = newDocument(client)
	client.Group = newGroup(client)
	client.Repo = newRepo(client)
	client.Search = newSearch(client)

	for _, option := range options {
		option(client)
	}

	if client.BaseURL == "" {
		client.BaseURL = internal.BaseURL
	}

	return client
}

func (c *Client) SetOption(opts ...clientOption) {
	for _, option := range opts {
		option(c)
	}
}

func WithBaseURL(url string) clientOption {
	return func(c *Client) {
		c.BaseURL = url
	}
}

func WithUserAgent(userAgent string) clientOption {
	return func(c *Client) {
		c.UserAgent = userAgent
	}
}

type option struct {
	Method  string
	Headers map[string]string
	Body    interface{}
}

func (c *Client) do(url string, options ...option) ([]byte, error) {
	var option option
	if len(options) > 1 {
		return nil, errors.New("options length more than one")
	} else if len(options) == 1 {
		option = options[0]
	}
	var (
		err     error
		req     *http.Request
		timeout = time.Duration(5 * time.Second)
		method  = option.Method
		body    *bytes.Buffer
	)

	if len(method) == 0 {
		method = "GET"
	}

	if option.Body != nil {
		requestBody, err := json.Marshal(option.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(requestBody)
	}

	if body != nil {
		req, err = http.NewRequest(method, url, body)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Auth-Token", c.Token)
	req.Header.Add("Content-type", "application/json")
	for k, v := range option.Headers {
		req.Header.Set(k, v)
	}

	httpClient := &http.Client{
		Timeout: timeout,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(respBody))
	}

	return respBody, nil
}
