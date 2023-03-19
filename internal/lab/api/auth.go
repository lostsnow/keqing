package api

type GameToken struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
}

type CookieToken struct {
	Uid         string `json:"uid"`
	CookieToken string `json:"cookie_token"`
}

type SToken struct {
	TokenType int    `json:"token_type"`
	Token     string `json:"token"`
}

type UserInfo struct {
	AccountName string `json:"account_name"`
	Mid         string `json:"mid"`
	Aid         string `json:"aid"`
}
