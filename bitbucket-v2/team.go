package bitbucket_v2

import (
	"errors"
	"net/url"
	"strconv"
)

type Teams struct {
	Page
	Values []*User `json:"values"`
}

func (c *Client) ListTeams(role string, index int) (*Teams, error) {
	teams := Teams{}
	if role == "" {
		return nil, errors.New("Missing role")
	}

	path := "/teams"
	params := url.Values{}
	params.Add("page", strconv.Itoa(index))
	params.Add("role", role)
	if err := c.do("GET", path, params, nil, "", &teams); err != nil {
		return nil, err
	}

	return &teams, nil
}
