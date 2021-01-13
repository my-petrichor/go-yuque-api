package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetDocumentList(token, namespace string) *ResponseDocSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs", namespace)

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

	doc := ResponseDocSerializer{}
	err = json.Unmarshal(body, &doc)
	// for i, docSerializer := range doc.Data {
	// 	fmt.Println(i, docSerializer.Id, docSerializer.Slug, docSerializer.Title)
	// }

	return &doc
}

func GetDocumentInfo(token, namespace, slug string) *ResponseDocDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs/%s", namespace, slug)

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

	doc := ResponseDocDetailSerializer{}
	json.Unmarshal(body, &doc)

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
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	doc := ResponseDocDetailSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}

func UpdateDocument(token, namespace, id, changeContent string) *ResponseDocDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs/%s?body=%s", namespace, id, changeContent)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	doc := ResponseDocDetailSerializer{}
	err = json.Unmarshal(body, &doc)
	if err != nil {
		fmt.Println(err)
	}

	return &doc
}
