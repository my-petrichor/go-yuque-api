package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseDocDetailSerializer struct {
	Abilities AbilitiesSerializer `json: "abilities"`
	Data      DocDetailSerializer `json: "data"`
}

type AbilitiesSerializer struct {
	Update  bool `json: "update"`
	Destroy bool `json: "destroy"`
}

//yuque DocDetailSerializer
type DocDetailSerializer struct {
	Id                 int                  `json: "id"`
	Slug               string               `json: "slug"`
	Title              string               `json: "title"`
	Book_id            int                  `json: "book_id"`
	Book               BookDetailSerializer `json: "book"`
	User_id            int                  `json: "user_id"`
	Creator            UserSerializer       `json: "creator"`
	Format             string               `json: "format"`
	Body               string               `json: "body"`
	Body_draft         string               `json: "body_draft"`
	Body_html          string               `json: "body_html"`
	Body_lake          string               `json: body_lake"`
	Body_draft_lake    string               `json: "body_draft_lake"`
	Public             int                  `json: "public"`
	Status             int                  `json: "status"`
	View_status        int                  `json: "view_status"`
	Read_status        int                  `json: "read_status"`
	Like_count         int                  `json: "like_count"`
	Comments_count     int                  `json: comments_count"`
	Content_updated_at string               `json: "content_updated_at"`
	Deleted_at         string               `json: "deleted_at"`
	Created_at         string               `json: "created_at"`
	Updated_at         string               `json: "updated_at"`
	Published_at       string               `json: "published_at"`
	First_published_at string               `json: "first_published_at"`
	Word_count         int                  `json: "word_count"`
	Cover              string               `json: "cover"`
	Description        string               `json: "description"`
	Custon_description string               `json: custon_description"`
	Hits               string               `json: "hits"`
	Serializer         string               `json: "serializer"`
}

//yuque BookSerializer
type BookDetailSerializer struct {
	Id                 int            `json: "id"`
	Type               string         `json: "type"`
	Slug               string         `json: "slug"`
	Name               string         `json: "name"`
	User_id            int            `json: "user_id"`
	Description        string         `json: "description"`
	Creator_id         int            `json: "creator_id"`
	Public             int            `json: "public"`
	Items_count        string         `json: items_count"`
	Likes_count        string         `json: "likes_count"`
	Watches_count      string         `json: "watched_count"`
	Content_updated_at string         `json: content_updated_at"`
	Updated_at         string         `json: "updated_at"`
	Created_at         string         `json: "created_at"`
	Namespace          string         `json: "namespace"`
	User               UserSerializer `json: "user"`
	Serializer         string         `json: "_serializer"`
}

//yuque UserSerializer
type UserSerializer struct {
	Id                 string `json: "id"`
	Type               string `json: "type"`
	Login              string `json: "login"`
	Name               string `json: "name"`
	Description        string `json: "description"`
	Avatar_url         string `json: "avatar_url"`
	Books_count        int    `json: "books_count"`
	Public_books_count int    `json: public_books_count"`
	Followers_count    int    `json: "followers_count"`
	Following_count    int    `json: "following_count"`
	Created_at         string `json: created_at"`
	Updated_at         string `json: "updated_at"`
	Serializer         string `json: "_serializer"`
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

func CreateDocument(token, namespace, slug, title string) *ResponseDocDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs?slug=%s&title=%s", namespace, slug, title)

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

func UpdateDocument(token, namespace, id, content string) *ResponseDocDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/docs/%s?body=%s", namespace, id, content)

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
