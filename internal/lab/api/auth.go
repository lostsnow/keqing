package api

type GameToken struct {
	UID   string `json:"uid"`
	Token string `json:"token"`
}

//nolint:tagliatelle
type CookieToken struct {
	UID         string `json:"uid"`
	CookieToken string `json:"cookie_token"`
}

//nolint:tagliatelle
type SToken struct {
	TokenType int    `json:"token_type"`
	Token     string `json:"token"`
}

//nolint:tagliatelle
type UserInfo struct {
	AccountName string `json:"account_name"`
	MID         string `json:"mid"`
	AID         string `json:"aid"`
}
