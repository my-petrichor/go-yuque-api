package yuque

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type client struct {
	Token   string
	BaseURL string

	User     *User
	Document *Document
	Repo     *Repo
	Group    *Group
	Search   *Search
}

type clientOption func(*client)

func NewClient(token string, options ...clientOption) *client {
	client := &client{
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

func (c *client) SetOption(opts ...clientOption) {
	for _, option := range opts {
		option(c)
	}
}

func (c *client) WithBaseURL(url string) clientOption {
	return func(c *client) {
		c.BaseURL = url
	}
}

// Option set http requset params
// need to set User-Agent Header
type Option struct {
	Method  string
	Headers map[string]string
	Body    interface{}
}

func (c *client) do(url string, options ...Option) ([]byte, error) {
	var option Option
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
		fmt.Println(string(requestBody), method, "requestBody---dasdasdas")
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
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-type", "application/json")
	}
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
