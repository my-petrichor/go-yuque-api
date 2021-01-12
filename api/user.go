package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUserInfo(token string) *ResponseUserDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/user")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	user := ResponseUserDetailSerializer{}
	json.Unmarshal(body, &user)

	return &user
}
