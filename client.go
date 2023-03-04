package gosms

import (
	"encoding/json"
	"errors"
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
	Url      string         `json:"url"`
	Query    MapStrings     `json:"query"`
	Headers  MapStrings     `json:"headers"`
	Debug    bool           `json:"debug"`
	Timeout  int            `json:"timeout"`
	Response *http.Response `json:"response"`
	BodyByte []byte         `json:"body_byte"`
}

func (c Client) Clone() *Client {
	if c.Timeout == 0 {
		c.Timeout = 10
	}
	return &c
}

func (c *Client) Get(assign any, query MapStrings, headers MapStrings) error {
	host := c.uriQueryMaps(c.Url, c.Query, query).String()
	allHeaders := MergeMapsString(c.Headers, headers)
	_, err := c.Request(http.MethodGet, host, nil, allHeaders, assign)
	return err
}
func (c *Client) PostForm(assign any, params url.Values, headers MapStrings) error {
	host := c.uriQueryMaps(c.Url, c.Query).String()
	allHeaders := MergeMapsString(c.Headers, headers)
	allHeaders.SetForce("Content-Type", FormContentType, false)
	_, err := c.Request(http.MethodPost, host, strings.NewReader(params.Encode()), allHeaders, assign)
	return err
}

// PostJson params MapStrings｜MapStringAny｜jsonString
func (c *Client) PostJson(assign any, body any, headers MapStrings) error {
	host := c.uriQueryMaps(c.Url, c.Query).String()
	allHeaders := MergeMapsString(c.Headers, headers)
	allHeaders.SetForce("Content-Type", JsonContentType, false)
	var jsonStr string
	var err error
	var verify bool
	//json字符串
	if str, ok := body.(string); ok {
		jsonStr = str
		verify = true
	}
	// 转json字符串
	if maps, ok := body.(MapStrings); ok {
		jsonStr, err = maps.ToJson()
		verify = true
	}
	// 转json字符串
	if maps, ok := body.(MapStringAny); ok {
		jsonStr, err = maps.ToJson()
		verify = true
	}
	if verify {
		_, err = c.Request(http.MethodPost, host, strings.NewReader(jsonStr), allHeaders, assign)
		return err
	}
	return errors.New("body must be string or MapStrings or MapStringAny")
}

// Request 任意发送请求
func (c *Client) Request(method, url string, body io.Reader, headers MapStrings, assign any) (*http.Response, error) {
	client := &http.Client{Timeout: time.Duration(c.Timeout) * time.Second}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Printf("%s: %s http.NewRequest error=%v", method, url, err)
		return c.Response, err
	}
	for headerKey, headerVal := range headers {
		req.Header.Set(headerKey, headerVal)
	}
	c.Response, err = client.Do(req)
	if err != nil {
		log.Printf("%s: %s client.Do error=%v", method, url, err)
		return c.Response, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Response.Body)
	c.BodyByte, _ = ioutil.ReadAll(c.Response.Body)
	if assign != nil {
		if err = json.Unmarshal(c.BodyByte, assign); err != nil {
			if c.Debug {
				log.Printf("Response json.Unmarshal error=%v", err)
			}
		}
	}
	if c.Debug {
		// GET 200 http://github.com/pkg6/go-sms headers body
		log.Printf("Response %s %s %s headers: %s body: %s", c.Response.Status, method, url, c.Response.Header, string(c.BodyByte))
	}
	return c.Response, err
}

// UriQueryMaps 参数进行拼接
func (c *Client) uriQueryMaps(uri string, maps ...MapStrings) *url.URL {
	parse, _ := url.Parse(uri)
	q := parse.Query()
	for _, m := range maps {
		for k, v := range m {
			q.Set(k, v)
		}
	}
	parse.RawQuery = q.Encode()
	return parse
}
