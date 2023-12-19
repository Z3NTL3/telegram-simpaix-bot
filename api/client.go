/*
Work for: Simpaix.net
Author: Z3NTL3
License: GNU
*/

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"z3ntl3/telegram-simpaix-bot/api/models"
	"z3ntl3/telegram-simpaix-bot/api/payloads"
)

const (
	BaseURL = "https://api.telegram.org"
)

type Client struct {
	token      string
	httpclient http.Client
	headers    http.Header
}

func NewClient(token string, timeout time.Duration) *Client {
	return &Client{
		token,
		http.Client{Timeout: timeout},
		http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
}

func Transform[T any](model interface{}) (*T, error) {
	m, ok := model.(T)

	if !ok {
		return nil, errors.New("err")
	}

	return &m, nil
}

// Only HTTP/HTTPS proxies are supported
func (c *Client) SetProxy(proxy string) error {
	uri, err := url.Parse(proxy)
	if err != nil {
		return err
	}

	c.httpclient.Transport = &http.Transport{
		Proxy: http.ProxyURL(uri),
	}

	return nil
}

func (c *Client) DeleteProxy(proxy string) {
	c.httpclient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment, // default
	}
}

func (c *Client) build_request(endpoint, http_method string, body_payload any) (*http.Request, error) {
	var body io.Reader

	if http_method == "post" {
		jsonRaw, err := json.Marshal(&body_payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonRaw)
	}

	req, err := http.NewRequest(
		http_method,
		fmt.Sprintf("%s/bot%s/%s", BaseURL, c.token, endpoint),
		body,
	)
	fmt.Println(
		req.URL)
	if err != nil {
		return nil, err
	}

	req.Header = c.headers
	return req, nil
}

func (c *Client) Authorize() (interface{}, error) {
	req, err := c.build_request("getMe", http.MethodGet, nil)
	if err != nil {
		return nil, nil
	}

	res, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf(
			"[Telegram API]: %s", res.Status,
		))
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	model := models.API_Resp[models.User]{}

	if err := json.Unmarshal(b, &model); err != nil {
		return nil, err
	}

	return model, nil
}

func (c *Client) GetUpdates(opts payloads.MUpdate) (interface{}, error) {
	req, err := c.build_request("getUpdates", http.MethodPost, opts)
	if err != nil {
		return nil, err
	}

	res, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("[Telegram API] %s - %s", res.Status, string(b)))
	}

	m := models.API_Resp[[]models.Update]{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return m, nil
}
