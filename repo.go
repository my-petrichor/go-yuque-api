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
func (r *Repo) ListAllUnderUser(login string) (*internal.ResponseBookSerializer, error) {
	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoListAllUnderUserPath, login)
		repo = internal.ResponseBookSerializer{}
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
func (r *Repo) ListAllUnderGroup(login string) (*internal.ResponseBookSerializer, error) {
	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoListAllUnderGroupPath, login)
		repo = internal.ResponseBookSerializer{}
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

// Get get repo detail info
func (r *Repo) GetInfo(namespace string) (*internal.ResponseBookDetailSerializer, error) {
	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoGetPath, namespace)
		repo = internal.ResponseBookDetailSerializer{}
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
func (r *Repo) GetDir(namespace string) (*internal.ResponseBookDirectoryStructure, error) {
	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoGetDirPath, namespace)
		repo = internal.ResponseBookDirectoryStructure{}
	)

	respBody, err := r.do(url)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(respBody))
	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

// CreateUnderUser create repo under the user
func (r *Repo) CreateUnderUser(login, newRepoSlug, kind, newRepoName string, options ...RepoOption) (*internal.ResponseBookDetailSerializer, error) {
	var opt RepoOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoCreateUnderUserPath, login)
		repo = internal.ResponseBookDetailSerializer{}
		body = struct {
			Name        string `json:"name"`
			Slug        string `json:"slug"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Public      int    `json:"public"`
		}{
			Name:        newRepoName,
			Slug:        newRepoSlug,
			Description: opt.Description,
			Type:        kind,
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
func (r *Repo) CreateUnderGroup(login, newRepoSlug, kind, newRepoName string, options ...RepoOption) (*internal.ResponseBookDetailSerializer, error) {
	var opt RepoOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoCreateUnderGroupPath, login)
		repo = internal.ResponseBookDetailSerializer{}
		body = struct {
			Name        string `json:"name"`
			Slug        string `json:"slug"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Public      int    `json:"public"`
		}{
			Name:        newRepoName,
			Slug:        newRepoSlug,
			Description: opt.Description,
			Type:        kind,
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
func (r *Repo) Update(namespace string, options ...RepoOption) (*internal.ResponseBookDetailSerializer, error) {
	var opt RepoOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoUpdatePath, namespace)
		repo = internal.ResponseBookDetailSerializer{}
		body = struct {
			Name        string `json:"name"`
			Slug        string `json:"slug"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Public      int    `json:"public"`
			Toc         string `json:"toc"`
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
func (r *Repo) Delete(namespace string) (*internal.ResponseBookDetailSerializer, error) {
	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoDeletePath, namespace)
		repo = internal.ResponseBookDetailSerializer{}
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
