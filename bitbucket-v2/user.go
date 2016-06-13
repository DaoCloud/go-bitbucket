package bitbucket_v2

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"dispaly_name"`
	UUID        string `json:"uuid"`
	Type        string `json:"type"`
}

type AutherType struct {
	User
	Raw string `json:"raw"`
}
