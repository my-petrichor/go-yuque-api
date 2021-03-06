package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetUserInfo() *ResponseUserDetailSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/user")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	user := ResponseUserDetailSerializer{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
	}

	return &user
}

func GetUserInfo(token string) *ResponseUserDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/user")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	user := ResponseUserDetailSerializer{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
	}

	return &user
}
