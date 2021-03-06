package main

import "net/http"

// type HttpClient interface {
// 	Do(req *http.Request) (*http.Response, error)
// }

type Client struct {
	Token  string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

func (c *Client) GetUserBookList(login string) *ResponseBookSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/users/%s/repos", login)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", c.Token)

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}