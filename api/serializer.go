package api

type ResponseUserDetailSerializer struct {
	Data UserDetailSerializer `json: "data"`
}

//yuque UserDetailSerializer
type UserDetailSerializer struct {
	Id                 int    `json: "id"`
	Type               string `json: "type"`
	Space_id           int    `json: "space_id"`
	Account_id         int    `json: "account_id"`
	Login              string `json: "login"`
	Name               string `json: "name"`
	Avatar_url         string `json: "avatar_url"`
	Books_count        int    `json: "books_count"`
	Public_books_count int    `json: public_books_count"`
	Followers_count    int    `json: "followers_count"`
	Following_count    int    `json: "following_count"`
	Public             int    `json: "public"`
	Description        string `json: "description"`
	Created_at         string `json: created_at"`
	Updated_at         string `json: "updated_at"`
	Serializer         string `json: "_serializer"`
}

//yuque UserSerializer
type UserSerializer struct {
	Id                 int    `json: "id"`
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

type ResponseBookSerializer struct {
	Data []BookDetailSerializer `json: "data"`
}

type ResponseBookDetailSerializer struct {
	Abilities AbilitiesSerializer  `json: "abilities"`
	Data      BookDetailSerializer `json: "data"`
}

//yuque BookDetailSerializer
type BookDetailSerializer struct {
	Id                 int            `json: "id"`
	Type               string         `json: "type"`
	Slug               string         `json: "slug"`
	Name               string         `json: "name"`
	User_id            int            `json: "user_id"`
	Description        string         `json: "description"`
	Toc                string         `json: "toc"`
	Toc_yml            string         `json: "toc_yml"`
	Creator_id         int            `json: "creator_id"`
	Public             int            `json: "public"`
	Items_count        int            `json: items_count"`
	Likes_count        int            `json: "likes_count"`
	Watches_count      int            `json: "watched_count"`
	Pinned_at          string         `json: "pinned_at"`
	Archived           string         `json: "archived"`
	Content_updated_at string         `json: content_updated_at"`
	Updated_at         string         `json: "updated_at"`
	Created_at         string         `json: "created_at"`
	Namespace          string         `json: "namespace"`
	User               UserSerializer `json: "user"`
	Serializer         string         `json: "_serializer"`
}

type ResponseBookDirectoryStructure struct {
	Data []BookSerializer `json: "data"`
}

type BookSerializer struct {
	Type         string `json: "type"`
	Title        string `json: "title"`
	Uuid         int    `json: "uuid"`
	Url          string `json: "url"`
	Prev_uuid    int    `json: "prev_uuid"`
	Sibling_uuid int    `json: "sibling_uuid"`
	Child_uuid   int    `json: "child_uuid"`
	Parent_uuid  int    `json: "parent_uuid"`
	Doc_id       int    `json: "doc_id"`
	Level        int    `json: "level"`
	Id           int    `json: "id"`
	Open_window  int    `json: "open_window"`
	Visible      int    `json: "visible"`
	Depth        int    `json: "depth"`
	Slug         string `json: "slug"`
}

//get a repo's doc list
type ResponseDocSerializer struct {
	Data []DocSerializer `json: "data"`
}

//yuque DocSerializer
type DocSerializer struct {
	Id                 int              `json: "id"`
	Slug               string           `json: "slug"`
	Title              string           `json: "title"`
	Description        string           `json: "description"`
	User_id            int              `json: "user_id"`
	Book_id            int              `json: "book_id"`
	Format             string           `json: "format"`
	Public             int              `json: "public"`
	Status             int              `json: "status"`
	View_status        int              `json: view_status"`
	Read_status        int              `json: read_status"`
	Likes_count        int              `json: "likes_count"`
	Comments_count     int              `json: comments_count"`
	Content_updated_at string           `json: content_updated_at"`
	Created_at         string           `json: "created_at"`
	Updated_at         string           `json: "updated_at"`
	Published_at       string           `json: "published_at"`
	First_published_at string           `json: "first_published_at"`
	Draft_version      int              `json: "draft_version"`
	Last_editor_id     int              `json: "last_editor_id"`
	Word_count         int              `json: "word_count"`
	Cover              string           `json: "cover"`
	Custom_description string           `json: "custom_description"`
	Last_editor        EditorSerializer `json: "last_editor"`
	Book               string           `json: "book"`
	Serializer         string           `json: "_serializer"`
}

type EditorSerializer struct {
	Id              int    `json: "id"`
	Type            string `json: "type"`
	Login           string `json: "login"`
	Name            string `json: "name"`
	Description     string `json: "description"`
	Avatar_url      string `json: "avatar_url"`
	Followers_count int    `json: "followers_count"`
	Following_count int    `json: "following_count"`
	Created_at      string `json: created_at"`
	Updated_at      string `json: "updated_at"`
	Serializer      string `json: "_serializer"`
}

type ResponseDocDetailSerializer struct {
	Abilities AbilitiesSerializer `json: "abilities"`
	Data      DocDetailSerializer `json: "data"`
}

type AbilitiesSerializer struct {
	Update  bool           `json: "update"`
	Destroy bool           `json: "destroy"`
	Doc     DocumentCreate `json: "doc"`
}

type DocumentCreate struct {
	Create bool `json: "create"`
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
