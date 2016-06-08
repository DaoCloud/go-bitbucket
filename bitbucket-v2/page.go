package bitbucket_v2

type Page struct {
	PageLen  int    `json:"pagelen"`
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	Next     string `json:"next"`
	previous string `json:"previous"`
}
