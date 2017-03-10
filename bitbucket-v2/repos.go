package bitbucket_v2

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/DaoCloud/go-bitbucket/bitbucket"
)

type Repos struct {
	Page
	Values []*Repo
}

type Repo struct {
	Scm         string    `json:"scm"`
	Website     string    `json:"website"`
	HasWike     bool      `json:"has_wiki"`
	Name        string    `json:"name"`
	ForkPolicy  string    `json:"fork_policy"`
	UUID        string    `json:"uuid"`
	Language    string    `json:"language"`
	Created     time.Time `json:"created_on"`
	FullName    string    `json:"full_name"`
	HasIssues   bool      `json:"has_issues"`
	Owner       User      `json:"owner"`
	Updated     time.Time `json:"updated_on"`
	Size        int64     `json:"size"`
	Type        string    `json:"type"`
	Private     bool      `json:"is_private"`
	Description string    `json:"description"`
	Links       struct {
		Clone []struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"clone"`
	} `json:"links"`
}

type TagsInfo struct {
	Page
	Values []*TargetValue `json:"values"`
}

type BranchInfo struct {
	Page
	Values []*TargetValue `json:"values"`
}

type ForkInfo struct {
	Page
	Values []*Repo `json:"values"`
}

type TargetValue struct {
	Type       string `json:"type"` //"tag" or "commit"
	Name       string `json:"name"`
	Repository Repo   `json:"repository"`
	Target     Refs   `json:"target"`
}

type Refs struct {
	Hash    string     `json:"hash"`
	Author  AutherType `json:"author"`
	Date    time.Time  `json:"date"`
	Message string     `json:"message"`
	Type    string     `json:"type"`
}

// Gets the repositories owned by the individual or team account.
func (c *Client) ListRepos(owner string, index int) (*Repos, error) {
	repos := Repos{}
	if owner == "" {
		return nil, nil
	}

	if owner == "self" {
		client_ng := bitbucket.New(c.ConsumerKey, c.ConsumerSecret, c.AccessToken, c.TokenSecret)
		if user, err := client_ng.Users.Current(); err != nil {
			return nil, err
		} else {
			owner = user.User.Username
		}
	}

	path := fmt.Sprintf("/repositories/%v", owner)
	params := url.Values{}
	params.Add("page", strconv.Itoa(index))
	if err := c.do("GET", path, params, nil, "", &repos); err != nil {
		return nil, err
	}

	return &repos, nil
}

func (c *Client) RepoInfo(owner, slug string) (*Repo, error) {
	repo := Repo{}
	if owner == "" || slug == "" {
		return nil, nil
	}

	path := fmt.Sprintf("/repositories/%v/%v", owner, slug)
	if err := c.do("GET", path, nil, nil, "", &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

func (c *Client) Tags(owner, slug string, index int) ([]*TargetValue, error) {
	ab := []*TargetValue{}
	if owner == "" || slug == "" {
		return nil, nil
	}
	path := fmt.Sprintf("/repositories/%v/%v/refs/tags", owner, slug)
	params := url.Values{}

	if index > 0 {
		params.Set("page", strconv.Itoa(index))
	}

	for {
		tags := TagsInfo{}
		if err := c.do("GET", path, params, nil, "", &tags); err != nil {
			return nil, err
		}
		ab = append(ab, tags.Values...)

		if tags.Next == "" || index > 0 {
			return ab, nil
		}

		u, err := url.Parse(tags.Next)
		if err != nil {
			return nil, err
		}

		params.Set("page", u.Query().Get("page"))
		path = strings.Replace(u.Path, "/2.0", "", 1)
	}
}

func (c *Client) Branches(owner, slug string, index int) (*BranchInfo, error) {
	branches := BranchInfo{}
	if owner == "" || slug == "" {
		return nil, nil
	}

	path := fmt.Sprintf("/repositories/%v/%v/refs/branches", owner, slug)
	params := url.Values{}
	params.Add("page", strconv.Itoa(index))
	if err := c.do("GET", path, params, nil, "", &branches); err != nil {
		return nil, err
	}

	return &branches, nil
}

func (c *Client) Forks(owner, slug string, index int) (*ForkInfo, error) {
	forks := ForkInfo{}
	if owner == "" || slug == "" {
		return nil, nil
	}

	path := fmt.Sprintf("/repositories/%v/%v/forks", owner, slug)
	params := url.Values{}
	params.Add("page", strconv.Itoa(index))
	if err := c.do("GET", path, params, nil, "", &forks); err != nil {
		return nil, err
	}

	return &forks, nil
}
