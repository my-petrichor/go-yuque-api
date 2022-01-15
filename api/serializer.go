package api

import "time"

type ResponseGroupUserSerializer struct {
	Data []GroupUserSerializer `json:"data"`
}

type ResponseGroupUserDetailSerializer struct {
	Data GroupUserDetailSerializer `json:"data"`
}

type GroupUserDetailSerializer struct {
	Create     bool           `json:"create"`
	Update     bool           `json:"update"`
	Destroy    bool           `json:"destroy"`
	ID         int            `json:"id"`
	GroupID    int            `json:"group_id"`
	UserID     int            `json:"user_id"`
	User       UserSerializer `json:"user"`
	Role       int            `json:"role"`
	Status     int            `json:"status"`
	CreatedAt  string         `json:"created_at"`
	Updatedat  string         `json:"updated_at"`
	Group      UserSerializer `json:"group"`
	Serializer string         `json:"_serializer"`
}

//yuque groupUserSerializer
type GroupUserSerializer struct {
	ID         int            `json:"id"`
	GroupID    int            `json:"group_id"`
	UserID     int            `json:"user_id"`
	User       UserSerializer `json:"user"`
	Role       int            `json:"role"`
	Status     int            `json:"status"`
	CreatedAt  string         `json:"created_at"`
	UpdatedAt  string         `json:"updated_at"`
	Group      UserSerializer `json:"group"`
	Serializer string         `json:"_serializer"`
}

type ResponseUserDetailSerializer struct {
	Abilities AbilitiesSerializer  `json:"abilities"`
	Meta      MetaSerializer       `json:"meta"`
	Data      UserDetailSerializer `json:"data"`
}

//yuque UserDetailSerializer
type UserDetailSerializer struct {
	ID                int    `json:"id"`
	Type              string `json:"type"`
	SpaceID           int    `json:"space_id"`
	AccountID         int    `json:"account_id"`
	OrganizationID    int    `json:"organization_id"`
	Login             string `json:"login"`
	Name              string `json:"name"`
	AvatarURL         string `json:"avatar_url"`
	OwnerID           int    `json:"owner_id"`
	BooksCount        int    `json:"books_count"`
	PublicBooksCount  int    `json:"public_books_count"`
	PublicTopicsCount int    `json:"public_topics_count"`
	MembersCount      int    `json:"members_count"`
	GrainsSum         int    `json:"grains_sum"`
	FollowersCount    int    `json:"followers_count"`
	FollowingCount    int    `json:"following_count"`
	Public            int    `json:"public"`
	Description       string `json:"description"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Serializer        string `json:"_serializer"`
}

type ResponseUserSerializer struct {
	Data []UserSerializer `json:"data"`
}

//yuque UserSerializer
type UserSerializer struct {
	ID                int    `json:"id"`
	Type              string `json:"type"`
	Login             string `json:"login"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	AvatarURL         string `json:"avatar_url"`
	BooksCount        int    `json:"books_count"`
	PublicBooksCount  int    `json:"public_books_count"`
	TopicsCount       int    `json:"topics_count"`
	PublicTopicsCount int    `json:"public_topics_count"`
	MembersCount      int    `json:"members_count"`
	Public            int    `json:"public"`
	FollowersCount    int    `json:"followers_count"`
	FollowingCount    int    `json:"following_count"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Serializer        string `json:"_serializer"`
}

type ResponseBookSerializer struct {
	Data []BookDetailSerializer `json:"data"`
}

type ResponseBookDetailSerializer struct {
	Abilities AbilitiesSerializer  `json:"abilities"`
	Data      BookDetailSerializer `json:"data"`
}

//yuque BookDetailSerializer
type BookDetailSerializer struct {
	ID               int            `json:"id"`
	Type             string         `json:"type"`
	Slug             string         `json:"slug"`
	Name             string         `json:"name"`
	UserID           int            `json:"user_id"`
	Description      string         `json:"description"`
	Toc              string         `json:"toc"`
	TocYml           string         `json:"toc_yml"`
	CreatorID        int            `json:"creator_id"`
	Public           int            `json:"public"`
	ItemsCount       int            `json:"items_count"`
	LikesCount       int            `json:"likes_count"`
	WatchesCount     int            `json:"watched_count"`
	PinnedAt         string         `json:"pinned_at"`
	Archived         string         `json:"archived"`
	ContentUpdatedAt string         `json:"content_updated_at"`
	UpdatedAt        string         `json:"updated_at"`
	CreatedAt        string         `json:"created_at"`
	Namespace        string         `json:"namespace"`
	User             UserSerializer `json:"user"`
	Serializer       string         `json:"_serializer"`
}

type ResponseBookDirectoryStructure struct {
	Data []BookSerializer `json:"data"`
}

type BookSerializer struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	UUID        int    `json:"uuid"`
	URL         string `json:"url"`
	PrevUUID    int    `json:"prev_uuid"`
	SiblingUUID int    `json:"sibling_uuid"`
	ChildUUID   int    `json:"child_uuid"`
	ParentUUID  int    `json:"parent_uuid"`
	DocID       int    `json:"doc_id"`
	Level       int    `json:"level"`
	ID          int    `json:"id"`
	OpenWindow  int    `json:"open_window"`
	Visible     int    `json:"visible"`
	Depth       int    `json:"depth"`
	Slug        string `json:"slug"`
}

//get a repo's doc list
type ResponseDocSerializer struct {
	Data []DocSerializer `json:"data"`
}

//yuque DocSerializer
type DocSerializer struct {
	ID                int              `json:"id"`
	Slug              string           `json:"slug"`
	Title             string           `json:"title"`
	Description       string           `json:"description"`
	UserID            int              `json:"user_id"`
	BookID            int              `json:"book_id"`
	Format            string           `json:"format"`
	Public            int              `json:"public"`
	Status            int              `json:"status"`
	ViewStatus        int              `json:"view_status"`
	ReadStatus        int              `json:"read_status"`
	LikesCount        int              `json:"likes_count"`
	CommentsCount     int              `json:"comments_count"`
	ContentUpdatedAt  string           `json:"content_updated_at"`
	CreatedAt         string           `json:"created_at"`
	UpdatedAt         string           `json:"updated_at"`
	PublishedAt       string           `json:"published_at"`
	FirstPublishedAt  string           `json:"first_published_at"`
	DraftVersion      int              `json:"draft_version"`
	LastEditorID      int              `json:"last_editor_id"`
	WordCount         int              `json:"word_count"`
	Cover             string           `json:"cover"`
	CustomDescription string           `json:"custom_description"`
	LastEditor        EditorSerializer `json:"last_editor"`
	Book              string           `json:"book"`
	Serializer        string           `json:"_serializer"`
}

type EditorSerializer struct {
	ID             int    `json:"id"`
	Type           string `json:"type"`
	Login          string `json:"login"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	AvatarURL      string `json:"avatar_url"`
	FollowersCount int    `json:"followers_count"`
	FollowingCount int    `json:"following_count"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Serializer     string `json:"_serializer"`
}

type ResponseDocDetailSerializer struct {
	Abilities AbilitiesSerializer `json:"abilities"`
	Data      DocDetailSerializer `json:"data"`
}

type AbilitiesSerializer struct {
	Update    bool                      `json:"update"`
	Destroy   bool                      `json:"destroy"`
	Doc       DocumentCreate            `json:"doc"`
	Read      bool                      `json:"read"`
	GroupUser GroupUserDetailSerializer `json:"group_user"`
	Repo      RepoDetailSerializer      `json:"repo"`
}

type MetaSerializer struct {
	TopicEnable int `json:"topic_enable"`
}

type RepoDetailSerializer struct {
	Create  bool `json:"create"`
	Update  bool `json:"update"`
	Destroy bool `json:"destroy"`
}

type DocumentCreate struct {
	Create bool `json:"create"`
}

//yuque DocDetailSerializer
type DocDetailSerializer struct {
	ID                int                  `json:"id"`
	Slug              string               `json:"slug"`
	Title             string               `json:"title"`
	BookID            int                  `json:"book_id"`
	Book              BookDetailSerializer `json:"book"`
	UserID            int                  `json:"user_id"`
	Creator           UserSerializer       `json:"creator"`
	Format            string               `json:"format"`
	Body              string               `json:"body"`
	BodyDraft         string               `json:"body_draft"`
	BodyHTML          string               `json:"body_html"`
	BodyLake          string               `json:"body_lake"`
	BodyDraftLake     string               `json:"body_draft_lake"`
	Public            int                  `json:"public"`
	Status            int                  `json:"status"`
	ViewStatus        int                  `json:"view_status"`
	ReadStatus        int                  `json:"read_status"`
	LikeCount         int                  `json:"like_count"`
	CommentsCount     int                  `json:"comments_count"`
	ContentUpdatedAt  string               `json:"content_updated_at"`
	DeletedAt         string               `json:"deleted_at"`
	CreatedAt         string               `json:"created_at"`
	UpdatedAt         string               `json:"updated_at"`
	PublishedAt       string               `json:"published_at"`
	FirstPublishedAt  string               `json:"first_published_at"`
	WordCount         int                  `json:"word_count"`
	Cover             string               `json:"cover"`
	Description       string               `json:"description"`
	CustonDescription string               `json:"custon_description"`
	Hits              int                  `json:"hits"`
	Serializer        string               `json:"serializer"`
}

type WebHookPublishDocSerializer struct {
	Data struct {
		ID     int    `json:"id"`
		Slug   string `json:"slug"`
		Title  string `json:"title"`
		BookID int    `json:"book_id"`
		Book   struct {
			ID               int         `json:"id"`
			Type             string      `json:"type"`
			Slug             string      `json:"slug"`
			Name             string      `json:"name"`
			UserID           int         `json:"user_id"`
			Description      string      `json:"description"`
			CreatorID        int         `json:"creator_id"`
			Public           int         `json:"public"`
			ItemsCount       int         `json:"items_count"`
			LikesCount       int         `json:"likes_count"`
			WatchesCount     int         `json:"watches_count"`
			ContentUpdatedAt time.Time   `json:"content_updated_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			CreatedAt        time.Time   `json:"created_at"`
			User             interface{} `json:"user"`
			Serializer       string      `json:"_serializer"`
		} `json:"book"`
		UserID int `json:"user_id"`
		User   struct {
			ID               int         `json:"id"`
			Type             string      `json:"type"`
			Login            string      `json:"login"`
			Name             string      `json:"name"`
			Description      interface{} `json:"description"`
			AvatarURL        string      `json:"avatar_url"`
			BooksCount       int         `json:"books_count"`
			PublicBooksCount int         `json:"public_books_count"`
			FollowersCount   int         `json:"followers_count"`
			FollowingCount   int         `json:"following_count"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			Serializer       string      `json:"_serializer"`
		} `json:"user"`
		Format             string      `json:"format"`
		Body               string      `json:"body"`
		BodyDraft          string      `json:"body_draft"`
		BodyHTML           string      `json:"body_html"`
		Public             int         `json:"public"`
		Status             int         `json:"status"`
		ViewStatus         int         `json:"view_status"`
		ReadStatus         int         `json:"read_status"`
		LikesCount         int         `json:"likes_count"`
		CommentsCount      int         `json:"comments_count"`
		ContentUpdatedAt   time.Time   `json:"content_updated_at"`
		DeletedAt          interface{} `json:"deleted_at"`
		CreatedAt          time.Time   `json:"created_at"`
		UpdatedAt          time.Time   `json:"updated_at"`
		PublishedAt        time.Time   `json:"published_at"`
		FirstPublishedAt   time.Time   `json:"first_published_at"`
		WordCount          int         `json:"word_count"`
		Serializer         string      `json:"_serializer"`
		Path               string      `json:"path"`
		Publish            bool        `json:"publish"`
		ActionType         string      `json:"action_type"`
		WebhookSubjectType string      `json:"webhook_subject_type"`
		ActorID            int         `json:"actor_id"`
	} `json:"data"`
}

type WebHookUpdateDocSerializer struct {
	Data struct {
		ID     int    `json:"id"`
		Slug   string `json:"slug"`
		Title  string `json:"title"`
		BookID int    `json:"book_id"`
		Book   struct {
			ID               int         `json:"id"`
			Type             string      `json:"type"`
			Slug             string      `json:"slug"`
			Name             string      `json:"name"`
			UserID           int         `json:"user_id"`
			Description      string      `json:"description"`
			CreatorID        int         `json:"creator_id"`
			Public           int         `json:"public"`
			ItemsCount       int         `json:"items_count"`
			LikesCount       int         `json:"likes_count"`
			WatchesCount     int         `json:"watches_count"`
			ContentUpdatedAt time.Time   `json:"content_updated_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			CreatedAt        time.Time   `json:"created_at"`
			User             interface{} `json:"user"`
			Serializer       string      `json:"_serializer"`
		} `json:"book"`
		UserID int `json:"user_id"`
		User   struct {
			ID               int         `json:"id"`
			Type             string      `json:"type"`
			Login            string      `json:"login"`
			Name             string      `json:"name"`
			Description      interface{} `json:"description"`
			AvatarURL        string      `json:"avatar_url"`
			BooksCount       int         `json:"books_count"`
			PublicBooksCount int         `json:"public_books_count"`
			FollowersCount   int         `json:"followers_count"`
			FollowingCount   int         `json:"following_count"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			Serializer       string      `json:"_serializer"`
		} `json:"user"`
		Format             string      `json:"format"`
		Body               string      `json:"body"`
		BodyDraft          string      `json:"body_draft"`
		BodyHTML           string      `json:"body_html"`
		Public             int         `json:"public"`
		Status             int         `json:"status"`
		ViewStatus         int         `json:"view_status"`
		ReadStatus         int         `json:"read_status"`
		LikesCount         int         `json:"likes_count"`
		CommentsCount      int         `json:"comments_count"`
		ContentUpdatedAt   time.Time   `json:"content_updated_at"`
		DeletedAt          interface{} `json:"deleted_at"`
		CreatedAt          time.Time   `json:"created_at"`
		UpdatedAt          time.Time   `json:"updated_at"`
		PublishedAt        time.Time   `json:"published_at"`
		FirstPublishedAt   time.Time   `json:"first_published_at"`
		WordCount          int         `json:"word_count"`
		Serializer         string      `json:"_serializer"`
		Path               string      `json:"path"`
		Publish            bool        `json:"publish"`
		ActionType         string      `json:"action_type"`
		WebhookSubjectType string      `json:"webhook_subject_type"`
		ActorID            int         `json:"actor_id"`
	} `json:"data"`
}

type WebHookDeleteDocSerializer struct {
	Data struct {
		ID     int    `json:"id"`
		Slug   string `json:"slug"`
		Title  string `json:"title"`
		BookID int    `json:"book_id"`
		Book   struct {
			ID               int         `json:"id"`
			Type             string      `json:"type"`
			Slug             string      `json:"slug"`
			Name             string      `json:"name"`
			UserID           int         `json:"user_id"`
			Description      string      `json:"description"`
			CreatorID        int         `json:"creator_id"`
			Public           int         `json:"public"`
			ItemsCount       int         `json:"items_count"`
			LikesCount       int         `json:"likes_count"`
			WatchesCount     int         `json:"watches_count"`
			ContentUpdatedAt time.Time   `json:"content_updated_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			CreatedAt        time.Time   `json:"created_at"`
			User             interface{} `json:"user"`
			Serializer       string      `json:"_serializer"`
		} `json:"book"`
		UserID int `json:"user_id"`
		User   struct {
			ID               int         `json:"id"`
			Type             string      `json:"type"`
			Login            string      `json:"login"`
			Name             string      `json:"name"`
			Description      interface{} `json:"description"`
			AvatarURL        string      `json:"avatar_url"`
			BooksCount       int         `json:"books_count"`
			PublicBooksCount int         `json:"public_books_count"`
			FollowersCount   int         `json:"followers_count"`
			FollowingCount   int         `json:"following_count"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			Serializer       string      `json:"_serializer"`
		} `json:"user"`
		Format             string    `json:"format"`
		Body               string    `json:"body"`
		BodyDraft          string    `json:"body_draft"`
		BodyHTML           string    `json:"body_html"`
		Public             int       `json:"public"`
		Status             int       `json:"status"`
		ViewStatus         int       `json:"view_status"`
		ReadStatus         int       `json:"read_status"`
		LikesCount         int       `json:"likes_count"`
		CommentsCount      int       `json:"comments_count"`
		ContentUpdatedAt   time.Time `json:"content_updated_at"`
		DeletedAt          time.Time `json:"deleted_at"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		PublishedAt        time.Time `json:"published_at"`
		FirstPublishedAt   time.Time `json:"first_published_at"`
		WordCount          int       `json:"word_count"`
		Serializer         string    `json:"_serializer"`
		Path               string    `json:"path"`
		Publish            bool      `json:"publish"`
		ActionType         string    `json:"action_type"`
		WebhookSubjectType string    `json:"webhook_subject_type"`
		ActorID            int       `json:"actor_id"`
	} `json:"data"`
}

type WebHookAddCommentSerializer struct {
	Data struct {
		ID         int         `json:"id"`
		UserID     int         `json:"user_id"`
		ParentID   interface{} `json:"parent_id"`
		BodyHTML   string      `json:"body_html"`
		LikesCount int         `json:"likes_count"`
		Mood       int         `json:"mood"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		Status     int         `json:"status"`
		ToUserID   interface{} `json:"to_user_id"`
		Type       interface{} `json:"type"`
		User       struct {
			ID             int         `json:"id"`
			Type           string      `json:"type"`
			Login          string      `json:"login"`
			Name           string      `json:"name"`
			Description    interface{} `json:"description"`
			AvatarURL      string      `json:"avatar_url"`
			FollowersCount int         `json:"followers_count"`
			FollowingCount int         `json:"following_count"`
			CreatedAt      time.Time   `json:"created_at"`
			UpdatedAt      time.Time   `json:"updated_at"`
			Serializer     string      `json:"_serializer"`
		} `json:"user"`
		Serializer  string `json:"_serializer"`
		Commentable struct {
			ID     int    `json:"id"`
			Slug   string `json:"slug"`
			Title  string `json:"title"`
			BookID int    `json:"book_id"`
			Book   struct {
				ID               int         `json:"id"`
				Type             string      `json:"type"`
				Slug             string      `json:"slug"`
				Name             string      `json:"name"`
				UserID           int         `json:"user_id"`
				Description      string      `json:"description"`
				CreatorID        int         `json:"creator_id"`
				Public           int         `json:"public"`
				ItemsCount       int         `json:"items_count"`
				LikesCount       int         `json:"likes_count"`
				WatchesCount     int         `json:"watches_count"`
				ContentUpdatedAt time.Time   `json:"content_updated_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				CreatedAt        time.Time   `json:"created_at"`
				User             interface{} `json:"user"`
				Serializer       string      `json:"_serializer"`
			} `json:"book"`
			UserID           int         `json:"user_id"`
			Format           string      `json:"format"`
			Body             string      `json:"body"`
			BodyDraft        string      `json:"body_draft"`
			BodyHTML         string      `json:"body_html"`
			Public           int         `json:"public"`
			Status           int         `json:"status"`
			ViewStatus       int         `json:"view_status"`
			ReadStatus       int         `json:"read_status"`
			LikesCount       int         `json:"likes_count"`
			CommentsCount    int         `json:"comments_count"`
			ContentUpdatedAt time.Time   `json:"content_updated_at"`
			DeletedAt        interface{} `json:"deleted_at"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			PublishedAt      time.Time   `json:"published_at"`
			FirstPublishedAt time.Time   `json:"first_published_at"`
			WordCount        int         `json:"word_count"`
			User             interface{} `json:"user"`
			Serializer       string      `json:"_serializer"`
			Path             string      `json:"path"`
		} `json:"commentable"`
		Path               string `json:"path"`
		ActionType         string `json:"action_type"`
		WebhookSubjectType string `json:"webhook_subject_type"`
		ActorID            int    `json:"actor_id"`
	} `json:"data"`
}

type WebHookUpdateCommentSerializer struct {
	Data struct {
		ID         int         `json:"id"`
		UserID     int         `json:"user_id"`
		ParentID   interface{} `json:"parent_id"`
		BodyHTML   string      `json:"body_html"`
		LikesCount int         `json:"likes_count"`
		Mood       int         `json:"mood"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		Status     int         `json:"status"`
		ToUserID   interface{} `json:"to_user_id"`
		Type       interface{} `json:"type"`
		User       struct {
			ID             int         `json:"id"`
			Type           string      `json:"type"`
			Login          string      `json:"login"`
			Name           string      `json:"name"`
			Description    interface{} `json:"description"`
			AvatarURL      string      `json:"avatar_url"`
			FollowersCount int         `json:"followers_count"`
			FollowingCount int         `json:"following_count"`
			CreatedAt      time.Time   `json:"created_at"`
			UpdatedAt      time.Time   `json:"updated_at"`
			Serializer     string      `json:"_serializer"`
		} `json:"user"`
		Serializer  string `json:"_serializer"`
		Commentable struct {
			ID     int    `json:"id"`
			Slug   string `json:"slug"`
			Title  string `json:"title"`
			BookID int    `json:"book_id"`
			Book   struct {
				ID               int         `json:"id"`
				Type             string      `json:"type"`
				Slug             string      `json:"slug"`
				Name             string      `json:"name"`
				UserID           int         `json:"user_id"`
				Description      string      `json:"description"`
				CreatorID        int         `json:"creator_id"`
				Public           int         `json:"public"`
				ItemsCount       int         `json:"items_count"`
				LikesCount       int         `json:"likes_count"`
				WatchesCount     int         `json:"watches_count"`
				ContentUpdatedAt time.Time   `json:"content_updated_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				CreatedAt        time.Time   `json:"created_at"`
				User             interface{} `json:"user"`
				Serializer       string      `json:"_serializer"`
			} `json:"book"`
			UserID           int         `json:"user_id"`
			Format           string      `json:"format"`
			Body             string      `json:"body"`
			BodyDraft        string      `json:"body_draft"`
			BodyHTML         string      `json:"body_html"`
			Public           int         `json:"public"`
			Status           int         `json:"status"`
			ViewStatus       int         `json:"view_status"`
			ReadStatus       int         `json:"read_status"`
			LikesCount       int         `json:"likes_count"`
			CommentsCount    int         `json:"comments_count"`
			ContentUpdatedAt time.Time   `json:"content_updated_at"`
			DeletedAt        interface{} `json:"deleted_at"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			PublishedAt      time.Time   `json:"published_at"`
			FirstPublishedAt time.Time   `json:"first_published_at"`
			WordCount        int         `json:"word_count"`
			User             interface{} `json:"user"`
			Serializer       string      `json:"_serializer"`
			Path             string      `json:"path"`
		} `json:"commentable"`
		Path               string `json:"path"`
		ActionType         string `json:"action_type"`
		WebhookSubjectType string `json:"webhook_subject_type"`
		ActorID            int    `json:"actor_id"`
	} `json:"data"`
}

type WebHookDeleteCommentSerializer struct {
	Data struct {
		ID         int         `json:"id"`
		UserID     int         `json:"user_id"`
		ParentID   interface{} `json:"parent_id"`
		BodyHTML   string      `json:"body_html"`
		LikesCount int         `json:"likes_count"`
		Mood       int         `json:"mood"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		Status     int         `json:"status"`
		ToUserID   interface{} `json:"to_user_id"`
		Type       interface{} `json:"type"`
		User       struct {
			ID             int         `json:"id"`
			Type           string      `json:"type"`
			Login          string      `json:"login"`
			Name           string      `json:"name"`
			Description    interface{} `json:"description"`
			AvatarURL      string      `json:"avatar_url"`
			FollowersCount int         `json:"followers_count"`
			FollowingCount int         `json:"following_count"`
			CreatedAt      time.Time   `json:"created_at"`
			UpdatedAt      time.Time   `json:"updated_at"`
			Serializer     string      `json:"_serializer"`
		} `json:"user"`
		Serializer  string `json:"_serializer"`
		Commentable struct {
			ID     int    `json:"id"`
			Slug   string `json:"slug"`
			Title  string `json:"title"`
			BookID int    `json:"book_id"`
			Book   struct {
				ID               int         `json:"id"`
				Type             string      `json:"type"`
				Slug             string      `json:"slug"`
				Name             string      `json:"name"`
				UserID           int         `json:"user_id"`
				Description      string      `json:"description"`
				CreatorID        int         `json:"creator_id"`
				Public           int         `json:"public"`
				ItemsCount       int         `json:"items_count"`
				LikesCount       int         `json:"likes_count"`
				WatchesCount     int         `json:"watches_count"`
				ContentUpdatedAt time.Time   `json:"content_updated_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				CreatedAt        time.Time   `json:"created_at"`
				User             interface{} `json:"user"`
				Serializer       string      `json:"_serializer"`
			} `json:"book"`
			UserID           int         `json:"user_id"`
			Format           string      `json:"format"`
			Body             string      `json:"body"`
			BodyDraft        string      `json:"body_draft"`
			BodyHTML         string      `json:"body_html"`
			Public           int         `json:"public"`
			Status           int         `json:"status"`
			ViewStatus       int         `json:"view_status"`
			ReadStatus       int         `json:"read_status"`
			LikesCount       int         `json:"likes_count"`
			CommentsCount    int         `json:"comments_count"`
			ContentUpdatedAt time.Time   `json:"content_updated_at"`
			DeletedAt        interface{} `json:"deleted_at"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			PublishedAt      time.Time   `json:"published_at"`
			FirstPublishedAt time.Time   `json:"first_published_at"`
			WordCount        int         `json:"word_count"`
			User             interface{} `json:"user"`
			Serializer       string      `json:"_serializer"`
			Path             string      `json:"path"`
		} `json:"commentable"`
		Path               string `json:"path"`
		ActionType         string `json:"action_type"`
		WebhookSubjectType string `json:"webhook_subject_type"`
		ActorID            int    `json:"actor_id"`
	} `json:"data"`
}

type WebHookAddCommentReplySerializer struct {
	Data struct {
		ID         int         `json:"id"`
		UserID     int         `json:"user_id"`
		ParentID   int         `json:"parent_id"`
		BodyHTML   string      `json:"body_html"`
		LikesCount int         `json:"likes_count"`
		Mood       int         `json:"mood"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		Status     int         `json:"status"`
		ToUserID   interface{} `json:"to_user_id"`
		Type       interface{} `json:"type"`
		User       struct {
			ID             int         `json:"id"`
			Type           string      `json:"type"`
			Login          string      `json:"login"`
			Name           string      `json:"name"`
			Description    interface{} `json:"description"`
			AvatarURL      string      `json:"avatar_url"`
			FollowersCount int         `json:"followers_count"`
			FollowingCount int         `json:"following_count"`
			CreatedAt      time.Time   `json:"created_at"`
			UpdatedAt      time.Time   `json:"updated_at"`
			Serializer     string      `json:"_serializer"`
		} `json:"user"`
		Serializer  string `json:"_serializer"`
		Commentable struct {
			ID     int    `json:"id"`
			Slug   string `json:"slug"`
			Title  string `json:"title"`
			BookID int    `json:"book_id"`
			Book   struct {
				ID               int         `json:"id"`
				Type             string      `json:"type"`
				Slug             string      `json:"slug"`
				Name             string      `json:"name"`
				UserID           int         `json:"user_id"`
				Description      string      `json:"description"`
				CreatorID        int         `json:"creator_id"`
				Public           int         `json:"public"`
				ItemsCount       int         `json:"items_count"`
				LikesCount       int         `json:"likes_count"`
				WatchesCount     int         `json:"watches_count"`
				ContentUpdatedAt time.Time   `json:"content_updated_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				CreatedAt        time.Time   `json:"created_at"`
				User             interface{} `json:"user"`
				Serializer       string      `json:"_serializer"`
			} `json:"book"`
			UserID           int         `json:"user_id"`
			Format           string      `json:"format"`
			Body             string      `json:"body"`
			BodyDraft        string      `json:"body_draft"`
			BodyHTML         string      `json:"body_html"`
			Public           int         `json:"public"`
			Status           int         `json:"status"`
			ViewStatus       int         `json:"view_status"`
			ReadStatus       int         `json:"read_status"`
			LikesCount       int         `json:"likes_count"`
			CommentsCount    int         `json:"comments_count"`
			ContentUpdatedAt time.Time   `json:"content_updated_at"`
			DeletedAt        interface{} `json:"deleted_at"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			PublishedAt      time.Time   `json:"published_at"`
			FirstPublishedAt time.Time   `json:"first_published_at"`
			WordCount        int         `json:"word_count"`
			User             interface{} `json:"user"`
			Serializer       string      `json:"_serializer"`
			Path             string      `json:"path"`
		} `json:"commentable"`
		Path               string `json:"path"`
		ActionType         string `json:"action_type"`
		WebhookSubjectType string `json:"webhook_subject_type"`
		ActorID            int    `json:"actor_id"`
	} `json:"data"`
}

type WebHookUpdateCommentReplySerializer struct {
	Data struct {
		ID         int         `json:"id"`
		UserID     int         `json:"user_id"`
		ParentID   int         `json:"parent_id"`
		BodyHTML   string      `json:"body_html"`
		LikesCount int         `json:"likes_count"`
		Mood       int         `json:"mood"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		Status     int         `json:"status"`
		ToUserID   interface{} `json:"to_user_id"`
		Type       interface{} `json:"type"`
		User       struct {
			ID             int         `json:"id"`
			Type           string      `json:"type"`
			Login          string      `json:"login"`
			Name           string      `json:"name"`
			Description    interface{} `json:"description"`
			AvatarURL      string      `json:"avatar_url"`
			FollowersCount int         `json:"followers_count"`
			FollowingCount int         `json:"following_count"`
			CreatedAt      time.Time   `json:"created_at"`
			UpdatedAt      time.Time   `json:"updated_at"`
			Serializer     string      `json:"_serializer"`
		} `json:"user"`
		Serializer  string `json:"_serializer"`
		Commentable struct {
			ID     int    `json:"id"`
			Slug   string `json:"slug"`
			Title  string `json:"title"`
			BookID int    `json:"book_id"`
			Book   struct {
				ID               int         `json:"id"`
				Type             string      `json:"type"`
				Slug             string      `json:"slug"`
				Name             string      `json:"name"`
				UserID           int         `json:"user_id"`
				Description      string      `json:"description"`
				CreatorID        int         `json:"creator_id"`
				Public           int         `json:"public"`
				ItemsCount       int         `json:"items_count"`
				LikesCount       int         `json:"likes_count"`
				WatchesCount     int         `json:"watches_count"`
				ContentUpdatedAt time.Time   `json:"content_updated_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				CreatedAt        time.Time   `json:"created_at"`
				User             interface{} `json:"user"`
				Serializer       string      `json:"_serializer"`
			} `json:"book"`
			UserID           int         `json:"user_id"`
			Format           string      `json:"format"`
			Body             string      `json:"body"`
			BodyDraft        string      `json:"body_draft"`
			BodyHTML         string      `json:"body_html"`
			Public           int         `json:"public"`
			Status           int         `json:"status"`
			ViewStatus       int         `json:"view_status"`
			ReadStatus       int         `json:"read_status"`
			LikesCount       int         `json:"likes_count"`
			CommentsCount    int         `json:"comments_count"`
			ContentUpdatedAt time.Time   `json:"content_updated_at"`
			DeletedAt        interface{} `json:"deleted_at"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			PublishedAt      time.Time   `json:"published_at"`
			FirstPublishedAt time.Time   `json:"first_published_at"`
			WordCount        int         `json:"word_count"`
			User             interface{} `json:"user"`
			Serializer       string      `json:"_serializer"`
			Path             string      `json:"path"`
		} `json:"commentable"`
		Path               string `json:"path"`
		ActionType         string `json:"action_type"`
		WebhookSubjectType string `json:"webhook_subject_type"`
		ActorID            int    `json:"actor_id"`
	} `json:"data"`
}

type WebHookDeleteCommentReplySerializer struct {
	Data struct {
		ID         int         `json:"id"`
		UserID     int         `json:"user_id"`
		ParentID   int         `json:"parent_id"`
		BodyHTML   string      `json:"body_html"`
		LikesCount int         `json:"likes_count"`
		Mood       int         `json:"mood"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
		Status     int         `json:"status"`
		ToUserID   interface{} `json:"to_user_id"`
		Type       interface{} `json:"type"`
		User       struct {
			ID             int         `json:"id"`
			Type           string      `json:"type"`
			Login          string      `json:"login"`
			Name           string      `json:"name"`
			Description    interface{} `json:"description"`
			AvatarURL      string      `json:"avatar_url"`
			FollowersCount int         `json:"followers_count"`
			FollowingCount int         `json:"following_count"`
			CreatedAt      time.Time   `json:"created_at"`
			UpdatedAt      time.Time   `json:"updated_at"`
			Serializer     string      `json:"_serializer"`
		} `json:"user"`
		Serializer  string `json:"_serializer"`
		Commentable struct {
			ID     int    `json:"id"`
			Slug   string `json:"slug"`
			Title  string `json:"title"`
			BookID int    `json:"book_id"`
			Book   struct {
				ID               int         `json:"id"`
				Type             string      `json:"type"`
				Slug             string      `json:"slug"`
				Name             string      `json:"name"`
				UserID           int         `json:"user_id"`
				Description      string      `json:"description"`
				CreatorID        int         `json:"creator_id"`
				Public           int         `json:"public"`
				ItemsCount       int         `json:"items_count"`
				LikesCount       int         `json:"likes_count"`
				WatchesCount     int         `json:"watches_count"`
				ContentUpdatedAt time.Time   `json:"content_updated_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				CreatedAt        time.Time   `json:"created_at"`
				User             interface{} `json:"user"`
				Serializer       string      `json:"_serializer"`
			} `json:"book"`
			UserID           int         `json:"user_id"`
			Format           string      `json:"format"`
			Body             string      `json:"body"`
			BodyDraft        string      `json:"body_draft"`
			BodyHTML         string      `json:"body_html"`
			Public           int         `json:"public"`
			Status           int         `json:"status"`
			ViewStatus       int         `json:"view_status"`
			ReadStatus       int         `json:"read_status"`
			LikesCount       int         `json:"likes_count"`
			CommentsCount    int         `json:"comments_count"`
			ContentUpdatedAt time.Time   `json:"content_updated_at"`
			DeletedAt        interface{} `json:"deleted_at"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			PublishedAt      time.Time   `json:"published_at"`
			FirstPublishedAt time.Time   `json:"first_published_at"`
			WordCount        int         `json:"word_count"`
			User             interface{} `json:"user"`
			Serializer       string      `json:"_serializer"`
			Path             string      `json:"path"`
		} `json:"commentable"`
		Path               string `json:"path"`
		ActionType         string `json:"action_type"`
		WebhookSubjectType string `json:"webhook_subject_type"`
		ActorID            int    `json:"actor_id"`
	} `json:"data"`
}
