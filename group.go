package yuque

import (
	"encoding/json"
	"fmt"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type Group struct {
	*client
}

func newGroup(c *client) *Group {
	return &Group{
		c,
	}
}

// ListAll list all groups that user join in
func (g *Group) ListAll(login string) (*internal.ResponseUserSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupListPath, login)
		group = internal.ResponseUserSerializer{}
	)

	respBody, err := g.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// ListPublic list all public group that user join in
func (g *Group) ListPublic() (*internal.ResponseUserSerializer, error) {
	var (
		url   = g.BaseURL + internal.GroupListPublicPath
		group = internal.ResponseUserSerializer{}
	)

	respBody, err := g.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// Get get single group detail info
func (g *Group) Get(groupLogin string) (*internal.ResponseUserDetailSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupGetPath, groupLogin)
		group = internal.ResponseUserDetailSerializer{}
	)

	respBody, err := g.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// GetMembers get group member info
func (g *Group) GetMembers(groupLogin string) (*internal.ResponseGroupUserSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupGetMemberPath, groupLogin)
		group = internal.ResponseGroupUserSerializer{}
	)

	respBody, err := g.do(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// Create create a group
func (g *Group) Create(repoName, newLogin, description string) (*internal.ResponseUserDetailSerializer, error) {
	var (
		url   = g.BaseURL + internal.GroupCreatePath
		group = internal.ResponseUserDetailSerializer{}
	)

	respBody, err := g.do(url, Option{Method: "POST", Body: map[string]string{
		"name":        repoName,
		"login":       newLogin,
		"description": description,
	}})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// Update update group info
func (g *Group) Update(groupLogin, newGroupLogin, newGroupName, newDescription string) (*internal.ResponseUserDetailSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupUpdatePath, groupLogin)
		group = internal.ResponseUserDetailSerializer{}
	)

	respBody, err := g.do(url, Option{Method: "PUT", Body: map[string]string{
		"name":        newGroupName,
		"login":       newGroupLogin,
		"description": newDescription,
	}})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// UpdateMember update group member authority
// role: 0 - manager  1 - ordinary
func (g *Group) UpdateMember(groupLogin, login string, role int) (*internal.ResponseGroupUserDetailSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupUpdateMemberPath, groupLogin, login)
		group = internal.ResponseGroupUserDetailSerializer{}
	)

	respBody, err := g.do(url, Option{Method: "PUT", Body: map[string]int{
		"role": role,
	}})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// Delete delete group
func (g *Group) Delete(groupLogin string) (*internal.ResponseUserDetailSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupDeletePath, groupLogin)
		group = internal.ResponseUserDetailSerializer{}
	)

	respBody, err := g.do(url, Option{Method: "DELETE"})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// DeleteMember delete group member
func (g *Group) DeleteMember(groupLogin, login string) (*internal.ResponseGroupUserDetailSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupDeleteMemberPath, groupLogin, login)
		group = internal.ResponseGroupUserDetailSerializer{}
	)

	respBody, err := g.do(url, Option{Method: "DELETE"})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}
