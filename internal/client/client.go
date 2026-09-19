package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"zenexport/internal/config"
)

const (
	clientTimeout = 60 * time.Second

	baseUrl     = "https://zenmoney.ru"
	profileUrl  = baseUrl + "/api/s1/profile/"
	enterUrl    = baseUrl + "/login/enter/"
	downloadUrl = baseUrl + "/export/download2011/"
)

type HttpClient struct {
	http *http.Client
}

type credentialsReq struct {
	Noredirect bool   `json:"noredirect"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

func NewHttpClient() *HttpClient {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	return &HttpClient{
		&http.Client{
			Timeout: clientTimeout,
			Jar:     jar,
		},
	}
}

func (c *HttpClient) ExportData(cfg *config.Config) ([]byte, error) {
	if err := c.resolveSessionCookie(); err != nil {
		return nil, err
	}

	if err := c.login(cfg.Login, cfg.Password); err != nil {
		return nil, err
	}

	data, err := c.fetchData()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *HttpClient) resolveSessionCookie() error {
	req, err := http.NewRequest(http.MethodGet, profileUrl, nil)
	if err != nil {
		return fmt.Errorf("resolve session: build request for %s: %w", profileUrl, err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("resolve session: request to %s failed: %w", profileUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("resolve session: unexpected response status from %s: %s", profileUrl, resp.Status)
	}

	return nil
}

func (c *HttpClient) login(login, password string) error {
	creds := credentialsReq{Noredirect: true, Login: login, Password: password}
	body, err := json.Marshal(creds)
	if err != nil {
		return fmt.Errorf("login: marshal credentials: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, enterUrl, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("login: build request for %s: %w", enterUrl, err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("login: request to %s failed: %w", enterUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("login: unexpected response status from %s: %s", enterUrl, resp.Status)
	}

	return nil
}

func (c *HttpClient) fetchData() ([]byte, error) {
	body := url.Values{}
	body.Add("month_filter", "0")
	body.Add("account_filter", "0")

	req, err := http.NewRequest(http.MethodPost, downloadUrl, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, fmt.Errorf("fetch data: build request for %s: %w", downloadUrl, err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch data: request to %s failed: %w", downloadUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("fetch data: unexpected response status from %s: %s", downloadUrl, resp.Status)
	}
	mt, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil || mt != "text/csv" {
		return nil, fmt.Errorf("fetch data: invalid response Content-Type: %s", resp.Header.Get("Content-Type"))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetch data: read response body: %w", err)
	}

	return data, nil
}
