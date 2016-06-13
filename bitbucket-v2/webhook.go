package bitbucket_v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type WebHook struct {
	Uuid         string    `json:"uuid"`
	Description  string    `json:"description,omitempty"`
	Url          string    `json:"url,omitempty"`
	Subject_type string    `json:"-"`
	Subject      Repo      `json:"-"`
	Active       bool      `json:"active"`
	Events       []string  `json:"events,omitempty"`
	Created_at   time.Time `json:"-"`
}

type WebhookInfo struct {
	Page
	Values []WebHook
}

//type WebHookReq struct {
//Description string   `json:"description"`
//Url         string   `json:"url"`
//Active      string   `json:"active"`
//Events      []string `json:"events"`
/*}*/

var MISS_ARGS = errors.New("one or more requird param is missing")

func (this *Client) ListWebhooks(owner, slug string, index int) (*WebhookInfo, error) {
	webhook := &WebhookInfo{}

	if owner == "" || slug == "" {
		return nil, MISS_ARGS
	}

	path := fmt.Sprintf("/repositories/%v/%v/hooks", owner, slug)
	params := url.Values{}
	params.Add("page", strconv.Itoa(index))
	if err := this.do("GET", path, params, nil, "", &webhook); err != nil {
		return nil, err
	}

	return webhook, nil
}

func NewWebhook(url, desc string, active bool, events []string) (*WebHook, error) {
	req := WebHook{
		Url:         url,
		Description: desc,
		Active:      active,
		Events:      events,
	}

	if events == nil {
		req.Events = []string{"repo:push", "issue:created", "issue:updated"}
	}
	return &req, nil
}

//Post json request to create a new webhook
func (this *Client) CreateUpdateWebHook(method, owner, slug string, webHook *WebHook) error {

	reqjson, err := json.Marshal(*webHook)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/repositories/%v/%v/hooks/%v", owner, slug, webHook.Uuid)
	return this.do(method, path, nil, nil, string(reqjson), webHook)
}

func (this *Client) GetDeleteWebHook(method, owner, slug string, webHook *WebHook) error {
	if owner == "" || slug == "" {
		return MISS_ARGS
	}

	path := fmt.Sprintf("/repositories/%v/%v/hooks/%v", owner, slug, webHook.Uuid)
	this.do(method, path, nil, nil, "", webHook)
	return nil
}
