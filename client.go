package gosms

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	FormContentType      = "application/x-www-form-urlencoded;charset=utf-8"
	FormASCIIContentType = "application/x-www-form-urlencoded"
	JsonContentType      = "application/json; charset=utf-8"
	JsonpContentType     = "application/javascript; charset=utf-8"
	JsonASCIIContentType = "application/json"
)

type Client struct {
	Debug bool `json:"debug"`
	//url
	Url string `json:"url"`
	//url query
	Query MapStrings `json:"query"`
	// request herader
	Header MapStrings `json:"headers"`
	// request timeout
	Timeout int `json:"timeout"`
	//response
	Response *http.Response `json:"response"`
	//response body
	BodyByte []byte `json:"body_byte"`
}

func (c Client) Clone() *Client {
	if c.Timeout == 0 {
		c.Timeout = 10
	}
	return &c
}

func (c *Client) SetUrl(url string) {
	c.Url = url
}

func (c *Client) SetQuery(query MapStrings) {
	c.Query = query
}

// GetFullUrl 参数进行拼接
func (c *Client) GetFullUrl(maps ...MapStrings) *url.URL {
	parse, _ := url.Parse(c.Url)
	q := parse.Query()
	for _, m := range maps {
		for k, v := range m {
			q.Set(k, v)
		}
	}
	parse.RawQuery = q.Encode()
	return parse
}

func (c *Client) SetBasicAuth(username, password string) {
	auth := username + ":" + password
	c.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))
}

func (c *Client) SetContentType(contentType string) {
	c.Header.SetForce("Content-Type", contentType, false)
}

// Get get请求
func (c *Client) Get(assign any, query MapStrings, header MapStrings) error {
	host := c.GetFullUrl(c.Query, query).String()
	_, err := c.Request(http.MethodGet, host, nil, header, assign)
	return err
}

// PostForm 表单提交
func (c *Client) PostForm(assign any, params url.Values, header MapStrings) error {
	host := c.GetFullUrl(c.Query).String()
	c.SetContentType(FormContentType)
	_, err := c.Request(http.MethodPost, host, strings.NewReader(params.Encode()), header, assign)
	return err
}

// PostJson json提交
func (c *Client) PostJson(assign any, jsonBody any, header MapStrings) error {
	host := c.GetFullUrl(c.Query).String()
	c.SetContentType(JsonContentType)
	var jsonStr string
	var err error
	var verify bool
	//json字符串
	if str, ok := jsonBody.(string); ok {
		jsonStr = str
		verify = true
	} else {
		jsonByte, err := json.Marshal(jsonBody)
		if err == nil {
			jsonStr = string(jsonByte)
			verify = true
		}
	}
	if verify {
		_, err = c.Request(http.MethodPost, host, strings.NewReader(jsonStr), header, assign)
		return err
	}
	return err
}

// Request 任意发送请求
func (c *Client) Request(method, fullUrl string, body io.Reader, header MapStrings, assign any) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Duration(c.Timeout) * time.Second,
	}
	if c.Debug {
		log.Printf("Request %s %s", method, fullUrl)
		log.Printf("Response headers: %s", header)
		log.Printf("Response body: %s", body)
	}
	req, err := http.NewRequest(method, fullUrl, body)
	if err != nil {
		log.Printf("%s: %s http.NewRequest error=%v", method, fullUrl, err)
		return c.Response, err
	}
	headers := MergeMapsString(c.Header, header)
	for headerKey, headerVal := range headers {
		req.Header.Set(headerKey, headerVal)
	}
	if userAgent := req.Header.Get("User-Agent"); userAgent == "" {
		req.Header.Set("User-Agent", "github.com/pkg6/go-sms")
	}
	c.Response, err = client.Do(req)
	if err != nil {
		log.Printf("%s: %s client.Do error=%v", method, fullUrl, err)
		return c.Response, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Response.Body)
	c.BodyByte, _ = ioutil.ReadAll(c.Response.Body)
	if assign != nil {
		_ = json.Unmarshal(c.BodyByte, assign)
	}
	if c.Debug {
		log.Printf("Response %s %s %s", c.Response.Status, method, fullUrl)
		log.Printf("Response headers: %s", c.Response.Header)
		log.Printf("Response body: %s", string(c.BodyByte))
	}
	return c.Response, err
}
