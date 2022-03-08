package internal

const (
	BaseURL = "https://www.yuque.com/api/v2"
)

const (
	UserGetPath = "/user"
)

const (
	GroupListPath         = "/users/%s/groups"
	GroupListPublicPath   = "/groups"
	GroupGetPath          = "/groups/%s"
	GroupGetMemberPath    = "/groups/%s/users"
	GroupCreatePath       = "/groups"
	GroupUpdatePath       = "/groups/%s"
	GroupUpdateMemberPath = "/groups/%s/users/%s"
	GroupDeletePath       = "/groups/%s"
	GroupDeleteMemberPath = "/groups/%s/users/%s"
)

const (
	DocumentListPath   = "/repos/%s/docs"
	DocumentGetPath    = "/repos/%s/docs/%s"
	DocumentCreatePath = "/repos/%s/docs?public=%d"
	DocumentUpdatePath = "/repos/%s/docs/%d?public=%d&_force_asi=%d"
	DocumentDeletePath = "/repos/%s/docs/%d"
)

const (
	RepoListAllUnderUserPath  = "/users/%s/repos"
	RepoListAllUnderGroupPath = "/groups/%s/repos"
	RepoGetPath               = "/repos/%s"
	RepoCreateUnderUserPath   = "/users/%s/repos?public=%d"
	RepoCreateUnderGroupPath  = "/groups/%s/repos?public=%d"
	RepoUpdatePath            = "/repos/%s?public=%d"
	RepoDeletePath            = "/repos/%s"
	RepoGetDirPath            = "/repos/%s/toc"
)

const (
	SearchPath = "/search"
)
