package yuque

import (
	"encoding/json"
	"fmt"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type Repo struct {
	*client
}

func newRepo(c *client) *Repo {
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
func (r *Repo) Get(namespace string) (*internal.ResponseBookDetailSerializer, error) {
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

	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

// CreateUnderUser create repo under the user
/*
 public
 0 - private
 1 - all user
 2 - space member
 3 - all user under space (include external contact)
 4 - only repo
*/
/*
 kind
 "Book" - library
 "Design" - board
*/
func (r *Repo) CreateUnderUser(login, newSlug, newName, description, kind string, public int) (*internal.ResponseBookDetailSerializer, error) {
	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoCreateUnderUserPath, login, public)
		repo = internal.ResponseBookDetailSerializer{}
	)

	respBody, err := r.do(url, Option{Method: "POST", Body: map[string]string{
		"name":        newName,
		"slug":        newSlug,
		"description": description,
		"type":        kind,
	}})
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
/*
 public
 0 - private
 1 - all user
 2 - space member
 3 - all user under space (include external contact)
 4 - only repo
*/
/*
 kind
 "Book" - library
 "Design" - board
*/
func (r *Repo) CreateUnderGroup(login, newSlug, newName, description, kind string, public int) (*internal.ResponseBookDetailSerializer, error) {
	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoCreateUnderGroupPath, login, public)
		repo = internal.ResponseBookDetailSerializer{}
	)

	respBody, err := r.do(url, Option{Method: "POST", Body: map[string]string{
		"name":        newName,
		"slug":        newSlug,
		"description": description,
		"type":        kind,
	}})
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
/*
 public
 0 - private
 1 - all user
 2 - space member
 3 - all user under space (include external contact)
 4 - only repo
*/
func (r *Repo) Update(namespace, newName, newSlug, newToc, description string, public int) (*internal.ResponseBookDetailSerializer, error) {
	var (
		url  = fmt.Sprintf(r.BaseURL+internal.RepoUpdatePath, namespace, public)
		repo = internal.ResponseBookDetailSerializer{}
	)

	respBody, err := r.do(url, Option{Method: "PUT", Body: map[string]string{
		"name":        newName,
		"slug":        newSlug,
		"toc":         newToc,
		"description": description,
	}})
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

	respBody, err := r.do(url, Option{Method: "DELETE"})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &repo)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}
