package api

import (
	"net/http"

	"github.com/lostsnow/keqing/internal/lab/api/endpoint"
)

//nolint:tagliatelle
type CookieTokenRequest struct {
	Request
	AccountID string `json:"account_id"`
	GameToken string `json:"game_token"`
}

type CookieTokenResponse struct {
	Response
	Data CookieToken `json:"data"`
}

func NewCookieTokenReq(id, token string) *CookieTokenRequest {
	return &CookieTokenRequest{
		Request: Request{
			URL:    endpoint.APITakumi + "/auth/api/getCookieAccountInfoByGameToken",
			Method: http.MethodGet,
		},
		AccountID: id,
		GameToken: token,
	}
}

func (r *CookieTokenRequest) Do() (*CookieToken, error) {
	var v CookieTokenResponse

	payload := map[string]string{
		"account_id": r.AccountID,
		"game_token": r.GameToken,
	}

	err := SendRequest(r, payload, &v)
	if err != nil {
		return nil, err
	}

	return &v.Data, nil
}
