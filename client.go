package gosms

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	log2 "log"
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
	debug      bool
	BaseURL    string
	QueryParam MapStrings
	Header     MapStrings
	Cookie     MapStrings
	TimeOut    int
	httpClient *http.Client
	Log        ILogger
	//只有调用do方法的时候才能调用
	Request  *http.Request
	Response *http.Response
}

func NewClient(baseURL string, fns ...func(client *Client)) *Client {
	client := Client{}.Clone()
	client.BaseURL = baseURL
	if client.httpClient == nil {
		client.httpClient = http.DefaultClient
	}
	if client.Log == nil {
		client.SetLog(Logger{Log: log2.Default()}.I())
	}
	if client.TimeOut == 0 {
		client.SetTimeOut(10)
	}
	for _, fn := range fns {
		fn(client)
	}
	return client
}

func (c Client) Clone() *Client {
	c.debug = false
	c.BaseURL = ""
	c.QueryParam = MapStrings{}
	c.Header = MapStrings{}
	c.Cookie = MapStrings{}
	c.TimeOut = 0
	c.httpClient = nil
	c.Log = nil
	c.Request = nil
	c.Response = nil
	return &c
}
func (c *Client) Debug() *Client {
	c.debug = true
	return c
}
func (c *Client) SetBaseURL(url string) *Client {
	c.BaseURL = strings.TrimRight(url, "/")
	return c
}
func (c *Client) SetTimeOut(timeOut int) *Client {
	c.TimeOut = timeOut
	return c
}
func (c *Client) SetLog(log ILogger) *Client {
	c.Log = log
	return c
}
func (c *Client) SetQueryParams(params MapStrings) *Client {
	c.QueryParam = MergeMapsString(c.QueryParam, params)
	return c
}
func (c *Client) SetQueryParam(key, value string) *Client {
	c.QueryParam.Set(key, value)
	return c
}
func (c *Client) SetTransport(transport *http.Transport) *Client {
	c.httpClient.Transport = transport
	return c
}
func (c *Client) SetBasicAuth(username, password string) *Client {
	if c.debug {
		c.Log.Debug("set BasicAuth username: %s ,password:", username, password)
	}
	c.SetHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
	return c
}

// GetUrl 参数进行拼接
func (c *Client) GetUrl(maps ...MapStrings) *url.URL {
	parse, _ := url.Parse(c.BaseURL)
	q := parse.Query()
	for _, m := range maps {
		for k, v := range m {
			q.Set(k, v)
		}
	}
	parse.RawQuery = q.Encode()
	return parse
}
func (c *Client) SetHeaders(headers MapStrings) *Client {
	c.Header = MergeMapsString(c.Header, headers)
	return c
}
func (c *Client) SetHeader(key, value string) *Client {
	c.Header.Set(key, value)
	return c
}
func (c *Client) WithUserAgent(userAgent string) *Client {
	c.Header.Set("User-Agent", userAgent)
	return c
}
func (c *Client) SetContentType(contentType string) *Client {
	c.Header.SetForce("Content-Type", contentType, false)
	return c
}

// Get get请求
func (c *Client) Get(query MapStrings) ([]byte, error) {
	return c.Do(http.MethodGet, c.GetUrl(c.QueryParam, query).String(), nil, nil, nil)
}

// PostForm 表单提交
func (c *Client) PostForm(params url.Values) ([]byte, error) {
	c.SetContentType(FormContentType)
	return c.Do(http.MethodPost, c.GetUrl(c.QueryParam).String(), strings.NewReader(params.Encode()), nil, nil)
}

// PostJson json提交
func (c *Client) PostJson(body any) ([]byte, error) {
	c.SetContentType(JsonContentType)
	var jsonStr string
	var err error
	var verify bool
	if str, ok := body.(string); ok {
		jsonStr = str
		verify = true
	} else {
		if jsonByte, err := json.Marshal(body); err == nil {
			jsonStr = string(jsonByte)
			verify = true
		}
	}
	if verify {
		return c.Do(http.MethodPost, c.GetUrl(c.QueryParam).String(), strings.NewReader(jsonStr), nil, nil)
	}
	return []byte(""), err
}

func (c *Client) Do(method, url string, body io.Reader, header MapStrings, cookie MapStrings) ([]byte, error) {
	var err error
	c.Request, err = http.NewRequest(method, url, body)
	if err != nil {
		c.Log.Error("Client.Request.NewRequest err=%v", err)
		return []byte(""), err
	}
	c.httpClient.Timeout = time.Duration(c.TimeOut) * time.Second
	headers := MergeMapsString(c.Header, header)
	for hk, hv := range headers {
		c.Request.Header.Set(hk, hv)
	}
	cookies := MergeMapsString(c.Cookie, cookie)
	for ck, cv := range cookies {
		c.Request.AddCookie(&http.Cookie{
			Name:  ck,
			Value: cv,
		})
	}
	if c.debug {
		c.Log.Debug("Client.Do.Request %s %s", method, url)
		c.Log.Debug("Client.Do.Request Header  %s", headers)
		c.Log.Debug("Client.Do.Request Cookie  %s", cookies)
	}
	c.Response, err = c.httpClient.Do(c.Request)
	if err != nil {
		c.Log.Error("client.Do.httpClient.Do err=%v", err)
		return []byte(""), err
	}
	defer c.Response.Body.Close()
	bodyByte, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		c.Log.Error("Client.Request.ioutil.ReadAll err=%v", err)
		return []byte(""), err
	}
	if c.debug {
		c.Log.Debug("Client.Do.Response %s %s", c.Response.Status, method)
		c.Log.Debug("Client.Do.Response Header  %s", c.Response.Header)
		c.Log.Debug("Client.Do.Response Cookie  %s", c.Response.Cookies())
		c.Log.Debug("Client.Do.Response body  %s", string(bodyByte))
	}
	return bodyByte, err
}
