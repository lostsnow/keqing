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
	AppID  int    `json:"app_id"`
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
	URL    string `json:"url"`
	AppID  int    `json:"app_id"`
	Device string `json:"device"`
	Ticket string `json:"ticket"`
}

func NewQrcodeFetch() *QrcodeFetchRequest {
	return &QrcodeFetchRequest{
		Request{
			URL:    endpoint.Hk4eSdk + "/hk4e_cn/combo/panda/qrcode/fetch",
			Method: http.MethodPost,
		},
	}
}

func (r *QrcodeFetchRequest) Do() (*QrcodeFetchResponseData, error) {
	payload := QrcodeFetchPayload{
		AppID:  4,
		Device: util.RandomString(64),
	}

	var v QrcodeFetchResponse

	err := SendRequest(r, payload, &v)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(v.Data.URL)
	if err != nil {
		return nil, fmt.Errorf("get QR code login ticket failed: %w", err)
	}

	v.Data.AppID = payload.AppID
	v.Data.Device = payload.Device
	v.Data.Ticket = u.Query().Get("ticket")

	return &v.Data, nil
}

func (r *QrcodeFetchRequest) ToImage(url string) ([]byte, error) {
	var qr []byte

	qr, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return nil, fmt.Errorf("generate login QR code failed: %w", err)
	}

	return qr, nil
}
