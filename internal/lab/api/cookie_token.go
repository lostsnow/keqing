package api

import (
	"net/http"

	"github.com/lostsnow/keqing/internal/lab/api/endpoint"
)

type CookieTokenRequest struct {
	Request
	AccountId string `json:"account_id"`
	GameToken string `json:"game_token"`
}

type CookieTokenResponse struct {
	Response
	Data CookieToken `json:"data"`
}

func NewCookieTokenReq(id, token string) *CookieTokenRequest {
	return &CookieTokenRequest{
		Request: Request{
			Url:    endpoint.ApiTakumi + "/auth/api/getCookieAccountInfoByGameToken",
			Method: http.MethodGet,
		},
		AccountId: id,
		GameToken: token,
	}
}

func (r *CookieTokenRequest) Do() (*CookieToken, error) {
	var v CookieTokenResponse
	payload := map[string]string{
		"account_id": r.AccountId,
		"game_token": r.GameToken,
	}
	err := SendRequest(r, payload, &v)
	if err != nil {
		return nil, err
	}

	return &v.Data, nil
}
