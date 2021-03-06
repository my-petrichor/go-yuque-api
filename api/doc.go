package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/FlashFeiFei/yuque/response"
)

func (c *Client) GetDocumentList(namespace string) *ResponseDocSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs", namespace)

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

	doc := ResponseDocSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}

func GetDocumentList(token, namespace string) *ResponseDocSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs", namespace)

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

	doc := ResponseDocSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}

func (c *Client) GetDocumentInfo(namespace, slug string) *ResponseDocDetailSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs/%s", namespace, slug)

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

	doc := ResponseDocDetailSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}

func (c *Client) GetDocumentID(namespace, slug string) int {
	doc := c.GetDocumentInfo(namespace, slug)
	return doc.Data.ID
}

func (c *Client) CreateDocument(namespace, changeSlug, changeTitle string) *ResponseDocDetailSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs?slug=%s&title=%s", namespace, changeSlug, changeTitle)

	req, err := http.NewRequest("POST", url, nil)
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
	doc := ResponseDocDetailSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}

func CreateDocument(token, namespace, changeSlug, changeTitle string) *ResponseDocDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs?slug=%s&title=%s", namespace, changeSlug, changeTitle)

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
	doc := ResponseDocDetailSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}

func (c *Client) ChangeDocumentSlug(id int, slug string) string {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/my-sakura/doc/docs/%d?slug=%s", id, slug)
	req, _ := http.NewRequest("PUT", url, nil)

	req.Header.Add("X-Auth-Token", c.Token)
	resp, _ := c.HTTPClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	doc := response.ResponseDocDetailSerializer{}
	err := json.Unmarshal(body, &doc)
	if err != nil {
		log.Fatal("unmarshal error", err)
	}

	return doc.Data.Slug
}

func (c *Client) UpdateDocument(namespace, changeContent string, id int) *ResponseDocDetailSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs/%d?body=%s", namespace, id, changeContent)

	req, err := http.NewRequest("PUT", url, nil)
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
	doc := ResponseDocDetailSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}

func UpdateDocument(token, namespace, changeContent string, id int) *ResponseDocDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs/%d?body=%s", namespace, id, changeContent)

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
	doc := ResponseDocDetailSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}

func (c *Client) DeleteDocument(token, namespace string, documentId int) *ResponseDocDetailSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs/%d", namespace, documentId)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	doc := ResponseDocDetailSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}

func DeleteDocument(token, namespace string, documentId int) *ResponseDocDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs/%d", namespace, documentId)

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
	doc := ResponseDocDetailSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}
