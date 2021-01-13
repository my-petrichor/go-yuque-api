package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Get the list of organizations that members have joined
func GetGroupList(token, login string) *ResponseUserSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/users/%s/groups", login)

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

	group := ResponseUserSerializer{}
	err = json.Unmarshal(body, &group)
	if err != nil {
		fmt.Println(err)
	}

	return &group
}

//Get the list of public organizations
func GetPublicGroupList(token string) *ResponseUserSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups")

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

	group := ResponseUserSerializer{}
	err = json.Unmarshal(body, &group)
	if err != nil {
		fmt.Println(err)
	}

	return &group
}

//Create group
func CreateGroup(token, changeName, changeLogin string) *ResponseUserDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups?name=%s&login=%s", changeName, changeLogin)

	req, err := http.NewRequest("POST", url, nil)
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

//Get single organization details
func GetGroupDetailInfo(token, login string) *ResponseUserDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s", login)

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

func GetGroupMemberInfo(token, login string) *ResponseGroupUserSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s/users", login)

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

	group := ResponseGroupUserSerializer{}
	err = json.Unmarshal(body, &group)
	if err != nil {
		fmt.Println(err)
	}

	return &group
}

func UpdateGroupDetailInfo(token, login, changeName, changeLogin string) *ResponseUserDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s?name=%s&login=%s&description", login, changeName, changeLogin)

	req, err := http.NewRequest("PUT", url, nil)
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

func UpdateGroupMember(token, groupLogin, login string) *ResponseGroupUserDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s/users/%s", groupLogin, login)

	req, err := http.NewRequest("PUT", url, nil)
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
	group := &ResponseGroupUserDetailSerializer{}
	err = json.Unmarshal(body, &group)
	if err != nil {
		fmt.Println(err)
	}

	return group
}

func DeleteGroup(token, login string) *ResponseUserDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s", login)

	req, err := http.NewRequest("DELETE", url, nil)
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

func DeleteGroupMember(token, groupLogin, login string) *ResponseGroupUserDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s/users/%s", groupLogin, login)

	req, err := http.NewRequest("DELETE", url, nil)
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

	user := ResponseGroupUserDetailSerializer{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
	}

	return &user
}
