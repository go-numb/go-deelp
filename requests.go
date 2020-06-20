package deepl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	API = "https://api.deepl.com"
)

type Client struct {
	key        string
	HTTPClient *http.Client
}

func New(key string) *Client {
	return &Client{
		key: key,
		HTTPClient: &http.Client{
			// 環境次第では30秒設定に
			Timeout: 10 * time.Second,
		},
	}
}

func (p *Client) Do(req Requester, results interface{}) error {
	u, err := url.ParseRequestURI(API)
	if err != nil {
		return err
	}
	u.Path = req.Path()

	// GETの場合はQueryに
	// POSTの場合はBodyParamにAuthを含める
	switch req.Method() {
	case http.MethodGet:
		q := &url.Values{}
		q.Set("auth_key", p.key)
		// Auth and requests queries
		u.RawQuery = q.Encode() + "&" + req.Query()

	case http.MethodPost:
		req.SetAuth(p.key)
	}

	buf := req.Param()
	r, err := http.NewRequest(req.Method(), u.String(), buf)
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Content-Length", fmt.Sprintf("%d", buf.Size()))
	r.Header.Set("User-Agent", "go client for DeepL")
	r.Header.Set("Accept", "*/*")
	res, err := p.HTTPClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(results); err != nil {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(body, results); err != nil {
			return fmt.Errorf("can not unmrashal data: %v", string(body))
		}

		return err
	}

	return nil
}
