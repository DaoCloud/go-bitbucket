package bitbucket_v2

import (
	"fmt"
	"time"
)

type Repos struct {
	Page
	Values []Repo
}

type Repo struct {
	Scm        string    `json:"scm"`
	Website    string    `json:"website"`
	HasWike    bool      `json:"has_wiki"`
	Name       string    `json:"name"`
	ForkPolicy string    `json:"fork_policy"`
	UUID       string    `json:"uuid"`
	Language   string    `json:"language"`
	Created    time.Time `json:"created_on"`
	FullName   string    `json:"full_name"`
	HasIssues  bool      `json:"has_issues"`
	//Owner       User      `json:"owner"`
	Updated     time.Time `json:"updated_on"`
	Size        int64     `json:"size"`
	Type        string    `json:"type"`
	IsPrivate   bool      `json:"is_private"`
	Description string    `json:"description"`
}

/*type RepoInfo struct {*/
//Page
//Values []TargetValue `json:"values"`
/*}*/

type TagsInfo struct {
	Page
	Values []TargetValue `json:"values"`
}

type BranchInfo struct {
	Page
	Values []TargetValue `json:"values"`
}

type ForkInfo struct {
	Page
	Values []Repo `json:"values"`
}

type TargetValue struct {
	Type       string `json:"type"` //"tag" or "commit"
	Name       string `json:"name"`
	Repository Repo   `json:"repository"`
	Target     Refs   `json:"target"`
}

type Refs struct {
	Hash string `json:"hash"`
	//Author  User      `json:"author"`
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
	Type    string    `json:"type"`
}

// Gets the repositories owned by the individual or team account.
func (this *Client) ListRepos(owner string) (*Repos, error) {
	repos := Repos{}
	if owner == "" {
		return nil, nil
	}

	path := "/repositories/" + owner

	if err := this.do("GET", path, nil, nil, &repos); err != nil {
		return nil, err
	}

	return &repos, nil
}

func (this *Client) RepoInfo(owner, slug string) (*Repo, error) {
	repo := Repo{}
	if owner == "" || slug == "" {
		return nil, nil
	}

	path := fmt.Sprintf("/repositories/%v/%v", owner, slug)
	if err := this.do("GET", path, nil, nil, &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

func (this *Client) Tags(owner, slug string) (*TagsInfo, error) {
	tags := TagsInfo{}
	if owner == "" || slug == "" {
		return nil, nil
	}

	path := fmt.Sprintf("/repositories/%v/%v/refs/tags", owner, slug)
	if err := this.do("GET", path, nil, nil, &tags); err != nil {
		return nil, err
	}

	return &tags, nil
}

func (this *Client) Branches(owner, slug string) (*BranchInfo, error) {
	branches := BranchInfo{}
	if owner == "" || slug == "" {
		return nil, nil
	}

	path := fmt.Sprintf("/repositories/%v/%v/refs/branches", owner, slug)
	if err := this.do("GET", path, nil, nil, &branches); err != nil {
		return nil, err
	}

	return &branches, nil
}

func (this *Client) Forks(owner, slug string) (*ForkInfo, error) {
	forks := ForkInfo{}
	if owner == "" || slug == "" {
		return nil, nil
	}

	path := fmt.Sprintf("/repositories/%v/%v/forks", owner, slug)
	if err := this.do("GET", path, nil, nil, &forks); err != nil {
		return nil, err
	}

	return &forks, nil
}
