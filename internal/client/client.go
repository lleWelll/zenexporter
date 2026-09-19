package client

import (
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
	// без Referer на страницу приложения экспорт отвечает 302 на /a/
	appUrl = baseUrl + "/a/"
)

type HttpClient struct {
	http *http.Client
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
	if err := c.login(cfg.Login, cfg.Password); err != nil {
		return nil, err
	}
	if err := c.verifyAuth(); err != nil {
		return nil, err
	}

	data, err := c.fetchData()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *HttpClient) login(login, password string) error {
	form := url.Values{}
	form.Set("noredirect", "true")
	form.Set("login", login)
	form.Set("password", password)

	req, err := http.NewRequest(http.MethodPost, enterUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("login: build request for %s: %w", enterUrl, err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

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

type profileResp struct {
	User json.RawMessage `json:"user"`
}

func (c *HttpClient) verifyAuth() error {
	resp, err := c.http.Get(profileUrl) // jar уже несёт авторизованный PHPSESSID
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var p profileResp
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return fmt.Errorf("verify auth: not authenticated (invalid response)")
	}
	if len(p.User) == 0 || string(p.User) == "null" {
		return fmt.Errorf("verify auth: not authenticated")
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
	req.Header.Set("Referer", appUrl)

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
