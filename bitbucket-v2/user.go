package bitbucket_v2

import "time"

type User struct {
	Username    string    `json:"username"`
	Website     string    `json:"website"`
	DisplayName string    `json:"dispaly_name"`
	UUID        string    `json:"uuid"`
	Type        string    `json:"type"`
	Created     time.Time `json:"created_on"`
	Links       struct {
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
	} `json:"links"`
}

type AutherType struct {
	User
	Raw string `json:"raw"`
}

func (c *Client) Userinfo() (*User, error) {
	user := User{}

	if err := c.do("GET", "/user", nil, nil, "", &user); err != nil {
		return nil, err
	}

	return &user, nil
}
