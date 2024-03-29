package yuque

import (
	"encoding/json"
	"fmt"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type Repo struct {
	*Client
}

type RepoOption struct {
	Slug        string
	Name        string
	Description string
	//"Book", "Design", "Sheet", "Column", "Resource", "Thread"
	Kind string
	/*
		0 - private
		1 - all user
		2 - space member
		3 - all user under space (include external contact)
		4 - only repo
	*/
	Public int
	// TODO Directory information for Book Repo
	// NewToc string
}

func newRepo(c *Client) *Repo {
	return &Repo{
		c,
	}
}

// ListAllUnderUser list all repo under user
func (r *Repo) ListAllUnderUser() (*ResponseBookSerializer, error) {
	var (
		url  = fmt.Sprintf(r.baseURL+internal.RepoListAllUnderUserPath, r.login)
		repo = ResponseBookSerializer{}
	)

	respBody, err := r.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

// ListAllUnderGroup list all repo under group
func (r *Repo) ListAllUnderGroup(groupLogin string) (*ResponseBookSerializer, error) {
	var (
		url  = fmt.Sprintf(r.baseURL+internal.RepoListAllUnderGroupPath, groupLogin)
		repo = ResponseBookSerializer{}
	)

	respBody, err := r.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

// Get get repo info
func (r *Repo) GetInfo(namespace string) (*ResponseBookDetailSerializer, error) {
	var (
		url  = fmt.Sprintf(r.baseURL+internal.RepoGetPath, namespace)
		repo = ResponseBookDetailSerializer{}
	)

	respBody, err := r.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

// GetDir get repo's directory
func (r *Repo) GetDir(namespace string) (*ResponseBookDirectoryStructure, error) {
	var (
		url  = fmt.Sprintf(r.baseURL+internal.RepoGetDirPath, namespace)
		repo = ResponseBookDirectoryStructure{}
	)

	respBody, err := r.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

// CreateUnderUser create repo under the user
func (r *Repo) CreateUnderUser(repoSlug, repoKind, repoName string, options ...RepoOption) (*ResponseBookDetailSerializer, error) {
	var opt RepoOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	var (
		url  = fmt.Sprintf(r.baseURL+internal.RepoCreateUnderUserPath, r.login)
		repo = ResponseBookDetailSerializer{}
		body = struct {
			Name        string `json:"name"`
			Slug        string `json:"slug"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Public      int    `json:"public"`
		}{
			Name:        repoName,
			Slug:        repoSlug,
			Description: opt.Description,
			Type:        repoKind,
			Public:      opt.Public,
		}
	)

	respBody, err := r.do(url, option{Method: "POST", Body: body})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

// CreateUnderGroup create repo under the group
func (r *Repo) CreateUnderGroup(repoSlug, repoKind, repoName string, options ...RepoOption) (*ResponseBookDetailSerializer, error) {
	var opt RepoOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	var (
		url  = fmt.Sprintf(r.baseURL+internal.RepoCreateUnderGroupPath, r.login)
		repo = ResponseBookDetailSerializer{}
		body = struct {
			Name        string `json:"name"`
			Slug        string `json:"slug"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Public      int    `json:"public"`
		}{
			Name:        repoName,
			Slug:        repoSlug,
			Description: opt.Description,
			Type:        repoKind,
			Public:      opt.Public,
		}
	)

	respBody, err := r.do(url, option{Method: "POST", Body: body})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

// Update update repo info
func (r *Repo) Update(namespace string, options ...RepoOption) (*ResponseBookDetailSerializer, error) {
	var opt RepoOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	var (
		url  = fmt.Sprintf(r.baseURL+internal.RepoUpdatePath, namespace)
		repo = ResponseBookDetailSerializer{}
		body = struct {
			Name        string `json:"name"`
			Slug        string `json:"slug"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Public      int    `json:"public"`
		}{
			Name:        opt.Name,
			Slug:        opt.Slug,
			Description: opt.Description,
			Type:        opt.Kind,
			Public:      opt.Public,
		}
	)

	respBody, err := r.do(url, option{Method: "PUT", Body: body})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

// Delete delete repo
func (r *Repo) Delete(namespace string) (*ResponseBookDetailSerializer, error) {
	var (
		url  = fmt.Sprintf(r.baseURL+internal.RepoDeletePath, namespace)
		repo = ResponseBookDetailSerializer{}
	)

	respBody, err := r.do(url, option{Method: "DELETE"})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}
