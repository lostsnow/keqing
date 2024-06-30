package api

import (
	"bytes"
	"crypto/md5" //nolint:gosec
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/spf13/viper"

	"github.com/lostsnow/keqing/internal/lab/api/salt"
	"github.com/lostsnow/keqing/pkg/util"
)

const (
	XRpcVersion = "2.44.1"
)

const (
	LabUA       = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) miHoYoBBS/" + XRpcVersion
	LabMobileUA = "Mozilla/5.0 (Linux; Android 12) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/106.0.5249.126 Mobile Safari/537.36 miHoYoBBS/" + XRpcVersion //nolint:lll
	OkHTTPUA    = "okhttp/4.8.0"
)

var httpClient *http.Client

var (
	ErrInvalidPayload     = errors.New("invalid payload")
	ErrRequestBodyInvalid = errors.New("request body invalid")
)

type RequestInterface interface {
	GetURL() string
	GetMethod() string
	GetHeaders() map[string]string
}

type ResponseInterface interface {
	GetRetCode() int
}

type Request struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

func (r *Request) GetURL() string {
	return r.URL
}

func (r *Request) GetMethod() string {
	return r.Method
}

func (r *Request) GetHeaders() map[string]string {
	return r.Headers
}

type Response struct {
	RetCode int    `json:"retcode"`
	Message string `json:"message"`
}

func (r *Response) GetRetCode() int {
	return r.RetCode
}

func SendRequest(r RequestInterface, payload any, v ResponseInterface) error {
	pl, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("%w for %s", ErrInvalidPayload, r.GetURL())
	}

	client, err := initHTTPClient()
	if err != nil {
		return err
	}

	var req *http.Request
	if http.MethodPost == r.GetMethod() {
		req, err = http.NewRequest(r.GetMethod(), r.GetURL(), bytes.NewBuffer(pl))
	} else {
		u, _ := url.Parse(r.GetURL())

		if payload != nil {
			params, ok := payload.(map[string]string)
			if ok {
				q := u.Query()
				for k, v := range params {
					q.Set(k, v)
					u.RawQuery = q.Encode()
				}
			}
		}

		req, err = http.NewRequest(r.GetMethod(), u.String(), nil)
	}

	if err != nil {
		return fmt.Errorf("new request %s failed: %w", r.GetURL(), err)
	}

	headers := r.GetHeaders()
	if len(headers) == 0 {
		req.Header.Add("User-Agent", "okhttp/4.8.0")

		if http.MethodPost == r.GetMethod() {
			req.Header.Add("Content-Type", "application/json")
		}
	} else {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do request %s failed: %w", r.GetURL(), err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("request %s parse body failed: %w", r.GetURL(), err)
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return fmt.Errorf("request %s unmarshal body %s failed: %w", r.GetURL(), body, err)
	}

	if v.GetRetCode() != 0 {
		return fmt.Errorf("%w, URL: %s, BODY: %s", ErrRequestBodyInvalid, r.GetURL(), body)
	}

	return nil
}

func initHTTPClient() (*http.Client, error) {
	if httpClient != nil {
		return httpClient, nil
	}

	client := &http.Client{}

	proxy := viper.GetString("proxy.mihoyo")

	if proxy != "" {
		p, err := url.Parse(proxy)
		if err != nil {
			return nil, fmt.Errorf("invalid mihoyo proxy url %s: %w", proxy, err)
		}

		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(p),
		}
	}

	httpClient = client

	return httpClient, nil
}

func GenerateDS(s string, payload any, query string) string {
	t := time.Now().Unix()

	var r string
	if s == salt.Prod {
		r = util.RandomString(6)
	} else {
		ri, _ := util.RandInt64(100000, 999999)
		r = strconv.FormatInt(ri, 10)
	}

	var b []byte
	if payload != nil {
		b, _ = json.Marshal(payload) //nolint:errchkjson
	}

	h := md5.New() //nolint:gosec
	h.Write([]byte(fmt.Sprintf("salt=%s&t=%d&r=%s&b=%s&q=%s", s, t, r, b, query)))

	return fmt.Sprintf("%d,%s,%s", t, r, hex.EncodeToString(h.Sum(nil)))
}

func GenerateLabHeaders(cookie string, query string) map[string]string {
	return map[string]string{
		"DS":                GenerateDS(salt.X4, nil, query),
		"Origin":            "https://webstatic.mihoyo.com",
		"Cookie":            cookie,
		"x-rpc-app_version": XRpcVersion,
		"User-Agent":        LabMobileUA,
		"x-rpc-client_type": "5",
		"Referer":           "https://webstatic.mihoyo.com/",
	}
}
