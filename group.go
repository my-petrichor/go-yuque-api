package yuque

import (
	"encoding/json"
	"fmt"

	"github.com/my-Sakura/go-yuque-api/internal"
)

type Group struct {
	*Client
}

type GroupOption struct {
	Login       string
	Name        string
	Description string
}

func newGroup(c *Client) *Group {
	return &Group{
		c,
	}
}

// ListAll list all groups that user join in
func (g *Group) ListAll(login string) (*ResponseUserSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupListPath, login)
		group = ResponseUserSerializer{}
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

// ListPublic list all public group
func (g *Group) ListPublic() (*ResponseUserSerializer, error) {
	var (
		url   = g.BaseURL + internal.GroupListPublicPath
		group = ResponseUserSerializer{}
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

// Get get single group info
func (g *Group) GetInfo(groupLogin string) (*ResponseUserDetailSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupGetPath, groupLogin)
		group = ResponseUserDetailSerializer{}
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
func (g *Group) GetMember(groupLogin string) (*ResponseGroupUserSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupGetMemberPath, groupLogin)
		group = ResponseGroupUserSerializer{}
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
func (g *Group) Create(newGroupName, newGroupLogin string, options ...GroupOption) (*ResponseUserDetailSerializer, error) {
	var opt GroupOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	var (
		url   = g.BaseURL + internal.GroupCreatePath
		group = ResponseUserDetailSerializer{}
		body  = struct {
			Name        string `json:"name"`
			Login       string `json:"login"`
			Description string `json:"description"`
		}{
			Name:        newGroupName,
			Login:       newGroupLogin,
			Description: opt.Description,
		}
	)

	respBody, err := g.do(url, option{Method: "POST", Body: body})
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
func (g *Group) Update(groupLogin string, options ...GroupOption) (*ResponseUserDetailSerializer, error) {
	var opt GroupOption
	if len(options) > 1 {
		return nil, ErrTooManyOptions
	} else if len(options) == 1 {
		opt = options[0]
	}

	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupUpdatePath, groupLogin)
		group = ResponseUserDetailSerializer{}
		body  = struct {
			Name        string `json:"name"`
			Login       string `json:"login"`
			Description string `json:"description"`
		}{
			Name:        opt.Name,
			Login:       opt.Login,
			Description: opt.Description,
		}
	)

	respBody, err := g.do(url, option{Method: "PUT", Body: body})
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
func (g *Group) UpdateMember(groupLogin, login string, role int) (*ResponseGroupUserDetailSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupUpdateMemberPath, groupLogin, login)
		group = ResponseGroupUserDetailSerializer{}
	)

	respBody, err := g.do(url, option{Method: "PUT", Body: map[string]int{
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
func (g *Group) Delete(groupLogin string) (*ResponseUserDetailSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupDeletePath, groupLogin)
		group = ResponseUserDetailSerializer{}
	)

	respBody, err := g.do(url, option{Method: "DELETE"})
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
func (g *Group) DeleteMember(groupLogin, login string) (*ResponseGroupUserDetailSerializer, error) {
	var (
		url   = fmt.Sprintf(g.BaseURL+internal.GroupDeleteMemberPath, groupLogin, login)
		group = ResponseGroupUserDetailSerializer{}
	)

	respBody, err := g.do(url, option{Method: "DELETE"})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}
