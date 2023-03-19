package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/skip2/go-qrcode"

	"github.com/lostsnow/keqing/internal/lab/api/endpoint"
	"github.com/lostsnow/keqing/pkg/util"
)

type QrcodeFetchPayload struct {
	AppId  int    `json:"app_id"`
	Device string `json:"device"`
}

type QrcodeFetchRequest struct {
	Request
}

type QrcodeFetchResponse struct {
	Response
	Data QrcodeFetchResponseData `json:"data"`
}

type QrcodeFetchResponseData struct {
	Url    string `json:"url"`
	AppId  int    `json:"app_id"`
	Device string `json:"device"`
	Ticket string `json:"ticket"`
}

func NewQrcodeFetch() *QrcodeFetchRequest {
	return &QrcodeFetchRequest{
		Request{
			Url:    endpoint.Hk4eSdk + "/hk4e_cn/combo/panda/qrcode/fetch",
			Method: http.MethodPost,
		},
	}
}

func (r *QrcodeFetchRequest) Do() (*QrcodeFetchResponseData, error) {
	payload := QrcodeFetchPayload{
		AppId:  4,
		Device: util.RandomString(64),
	}
	var v QrcodeFetchResponse
	err := SendRequest(r, payload, &v)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(v.Data.Url)
	if err != nil {
		return nil, fmt.Errorf("get QR code login ticket failed: %s", err)
	}
	v.Data.AppId = payload.AppId
	v.Data.Device = payload.Device
	v.Data.Ticket = u.Query().Get("ticket")

	return &v.Data, nil
}

func (r *QrcodeFetchRequest) ToImage(url string) ([]byte, error) {
	var qr []byte
	qr, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return nil, fmt.Errorf("generate login QR code failed: %s", err)
	}
	return qr, nil
}
