package api

import (
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/lostsnow/keqing/internal/lab/api/endpoint"
	"github.com/lostsnow/keqing/internal/lab/api/salt"
	"github.com/lostsnow/keqing/pkg/util"
)

type STokenPayload struct {
	AccountID int64  `json:"account_id"`
	GameToken string `json:"game_token"`
}

type STokenRequest struct {
	Request
	STokenPayload
}

type STokenResponse struct {
	Response
	Data STokenResponseData `json:"data"`
}

type STokenResponseData struct {
	Token    SToken   `json:"token"`
	UserInfo UserInfo `json:"user_info"`
}

func NewSTokenReq(id int64, token string) *STokenRequest {
	payload := STokenPayload{
		AccountID: id,
		GameToken: token,
	}

	return &STokenRequest{
		Request: Request{
			URL:    endpoint.PassportAPI + "/account/ma-cn-session/app/getTokenByGameToken",
			Method: http.MethodPost,
			Headers: map[string]string{
				"x-rpc-app_version":  XRpcVersion,
				"DS":                 GenerateDS(salt.Prod, payload, ""),
				"x-rpc-aigis":        "",
				"Content-Type":       "application/json",
				"Accept":             "application/json",
				"x-rpc-game_biz":     "bbs_cn",
				"x-rpc-sys_version":  "11",
				"x-rpc-device_id":    strings.ReplaceAll(uuid.New().String(), "-", ""),
				"x-rpc-device_fp":    util.RandomString(13),
				"x-rpc-device_name":  "Chrome 108.0.0.0",
				"x-rpc-device_model": "Windows 10 64-bit",
				"x-rpc-app_id":       "bll8iq97cem8",
				"x-rpc-client_type":  "4",
				"User-Agent":         OkHTTPUA,
			},
		},
		STokenPayload: payload,
	}
}

func (r *STokenRequest) Do() (*STokenResponseData, error) {
	var v STokenResponse

	payload := STokenPayload{
		AccountID: r.AccountID,
		GameToken: r.GameToken,
	}

	err := SendRequest(r, payload, &v)
	if err != nil {
		return nil, err
	}

	return &v.Data, nil
}
