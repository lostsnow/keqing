package api

type GameToken struct {
	UID   string `json:"uid"`
	Token string `json:"token"`
}

type CookieToken struct {
	UID         string `json:"uid"`
	CookieToken string `json:"cookie_token"`
}

type SToken struct {
	TokenType int    `json:"token_type"`
	Token     string `json:"token"`
}

type UserInfo struct {
	AccountName string `json:"account_name"`
	MID         string `json:"mid"`
	AID         string `json:"aid"`
}
