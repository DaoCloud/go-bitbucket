package bitbucket_v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Webhook struct {
	UUID        string
	Description string
	Url         string
	SubjectType string `json:"subject_type"`
	Subject     Repo
	Active      bool
	Created     time.Time `json:"created_at"`
}

type WebhookInfo struct {
	Page
	Values Webhook
}

type WebHookReq struct {
	Description string            `json:"description"`
	Url         string            `json:"url"`
	Active      string            `json:"active"`
	Events      map[string]string `json:"events"`
}

var MISS_ARGS = errors.New("one or more requird param is missing")

func (this *Client) ListWebhooks(owner, slug string) (*WebhookInfo, error) {
	webhook := &WebhookInfo{}

	if owner == "" || slug == "" {
		return nil, MISS_ARGS
	}

	path := fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%v/%v/hooks/", owner, slug)
	if err := this.do("GET", path, nil, nil, &webhook); err != nil {
		return nil, err
	}

	return webhook, nil
}

func NewWebhookJson(url, desc, active string, events map[string]string) ([]byte, error) {
	req := WebHookReq{
		Url:         url,
		Description: desc,
		Active:      active,
		Events:      events,
	}

	if reqjson, err := json.Marshal(req); err != nil {
		return nil, err
	} else {
		return reqjson, nil
	}
}

//Post json request to create a new webhook
func (this *Client) Webhook(method, owner, slug, url, desc, active string,
	events map[string]string) error {

	if owner == "" || slug == "" || url == "" {
		return MISS_ARGS
	}

	reqjson, err := NewWebhookJson(url, desc, active, events)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%v/%v/hooks", owner, slug)
	this.do(method, path, nil, reqjson, nil)
	return nil
}

func (this *Client) DeleteWebhoook(owner, slug, uuid string) error {
	if owner == "" || slug == "" || uuid == "" {
		return MISS_ARGS
	}

	path := fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%v/%v/hooks/%v", owner, slug, uuid)
	this.do("DELETE", path, nil, nil, nil)
	return nil
}
