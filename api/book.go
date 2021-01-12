package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUserBookList(token, login string) *ResponseBookSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/users/%s/repos", login)

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

	book := ResponseBookSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

//todo:
//Currently, it is not possible to obtain the team knowledge base list,
//and the obtained result is the same as the user knowledge base list
func GetGroupBookList(token, login string) *ResponseBookSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s/repos", login)

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

	book := ResponseBookSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}
	for i, bookDetailSerializer := range book.Data {
		fmt.Println(i, bookDetailSerializer.Type, bookDetailSerializer.Slug, bookDetailSerializer.Name)
	}

	return &book
}

func GetBookInfo(token, namespace string) *ResponseBookDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s", namespace)

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

	book := ResponseBookDetailSerializer{}
	json.Unmarshal(body, &book)

	return &book
}

func CreateUserBook(token, login, slug, name string) *ResponseBookDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/users/%s/repos?slug=%s&name=%s", login, slug, name)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

/*public:
0  private
1  public
2  visible for space members
3  visible for everyone int space members
4  visible for book members
*/
func CreateGrouprBook(token, login, slug, name string, public int) *ResponseBookDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s/repos?slug=%s&name=%s&public=%d", login, slug, name, public)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func UpdateBook(token, namespace, name, slug string) *ResponseBookDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s?name=%s&slug=%s", namespace, name, slug)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func DeleteBook(token, namespace string) *ResponseBookDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s", namespace)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

//need repo abilities.read authority
func GetBookDirectoryStructure(token, namespace string) *ResponseBookDirectoryStructure {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/toc", namespace)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	book := ResponseBookDirectoryStructure{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}
