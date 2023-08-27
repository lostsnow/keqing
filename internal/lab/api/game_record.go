package api

import (
	"fmt"
	"net/http"

	"github.com/lostsnow/keqing/internal/lab/api/endpoint"
)

type GameRecordRequest struct {
	Request
	AccountID   string `json:"account_id"`
	CookieToken string `json:"cookie_token"`
}

type GameRecordResponse struct {
	Response
	Data GameRecordResponseData `json:"data"`
}

type GameRecordResponseData struct {
	GameRecords []GameRecord `json:"list"`
}

func NewGameRecordReq(id, cookieToken string) *GameRecordRequest {
	return &GameRecordRequest{
		Request: Request{
			URL:    endpoint.APITakumiRecord + "/game_record/card/wapi/getGameRecordCard",
			Method: http.MethodGet,
			Headers: GenerateLabHeaders(
				fmt.Sprintf("account_id=%s;cookie_token=%s", id, cookieToken), "uid="+id),
		},
		AccountID:   id,
		CookieToken: cookieToken,
	}
}

func (r *GameRecordRequest) Do() (*GameRecordResponseData, error) {
	var v GameRecordResponse

	payload := map[string]string{
		"uid": r.AccountID,
	}

	err := SendRequest(r, payload, &v)
	if err != nil {
		return nil, err
	}

	return &v.Data, nil
}
