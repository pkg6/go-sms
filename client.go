package gosms

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	log2 "log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	HeaderUserAgentKey   = "User-Agent"
	HeaderContentTypeKey = "Content-Type"
	FormContentType      = "application/x-www-form-urlencoded;charset=utf-8"
	FormASCIIContentType = "application/x-www-form-urlencoded"
	JsonContentType      = "application/json; charset=utf-8"
	JsonpContentType     = "application/javascript; charset=utf-8"
	JsonASCIIContentType = "application/json"
)

type Client struct {
	debug bool
	log   ILogger
	url   *url.URL

	BaseURL    string
	QueryParam MapStrings
	Header     MapStrings
	Cookie     MapStrings
	TimeOut    int
	httpClient *http.Client
	//只有调用do方法的时候才能调用
	Request  *http.Request
	Response *http.Response
}

func NewClient(baseURL string, fns ...func(client *Client)) *Client {
	client := Client{}.Clone()
	client.SetBaseURL(baseURL)
	if client.httpClient == nil {
		client.httpClient = http.DefaultClient
	}
	if client.log == nil {
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
	c.log = nil
	c.url, _ = url.Parse("")
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
	c.log = log
	return c
}

// WithUserAgent 携带User-Agent
func (c *Client) WithUserAgent(userAgent string) *Client {
	c.SetHeader(HeaderUserAgentKey, userAgent)
	return c
}

// WithContentType 如果设置Content-Type就不需要进行覆盖
func (c *Client) WithContentType(contentType string) *Client {
	c.SetHeader(HeaderContentTypeKey, contentType)
	return c
}

// WithBasicAuth 携带Authorization
func (c *Client) WithBasicAuth(username, password string) *Client {
	if c.debug {
		c.log.Debug("with BasicAuth username: %s ,password:", username, password)
	}
	c.WithToken(base64.StdEncoding.EncodeToString([]byte(username+":"+password)), "", "")
	return c
}

// WithToken 携带token
func (c *Client) WithToken(token, tokenKey, tokenType string) *Client {
	if tokenType == "" {
		token = "Basic " + token
	} else {
		token = tokenType + token
	}
	if tokenKey == "" {
		tokenKey = "Authorization"
	}
	if c.debug {
		c.log.Debug("set Token Key=%s value", tokenType, token)
	}
	c.SetHeader(tokenKey, token)
	return c
}

// SetHeaders 批量设置header
func (c *Client) SetHeaders(params MapStrings) *Client {
	for p, v := range params {
		c.SetHeader(p, v)
	}
	return c
}

// SetHeader 单独设置header
func (c *Client) SetHeader(key, value string) *Client {
	c.Header.Set(key, value)
	return c
}

// 如果设置就不需要进行覆盖
func (c *Client) header(key, value string) *Client {
	c.Header.SetForce(key, value, false)
	return c
}

// SetQueryParams 设置url请求参数
func (c *Client) SetQueryParams(params MapStrings) *Client {
	for p, v := range params {
		c.SetQueryParam(p, v)
	}
	return c
}

// SetQueryParam 设置url请求参数
func (c *Client) SetQueryParam(key, value string) *Client {
	c.QueryParam.Set(key, value)
	return c
}

// BuildUrl 生成完成的url参数复制给URL
func (c *Client) BuildUrl(maps ...MapStrings) {
	Url, _ := url.Parse(c.BaseURL)
	q := Url.Query()
	for _, m := range maps {
		for k, v := range m {
			q.Set(k, v)
		}
	}
	Url.RawQuery = q.Encode()
	c.url = Url
}

// Get get请求
func (c *Client) Get(query MapStrings) ([]byte, error) {
	c.BuildUrl(c.QueryParam, query)
	return c.Do(http.MethodGet, c.url.String(), nil, nil, nil)
}

// FileInfo 上传文件基本信息
type FileInfo struct {
	//提交的时候，服务端需要对应一个字段名称
	Name string
	// 读取文件 os.Open("text.go")
	File *os.File
}

// PostFiles 多文件上传
func (c *Client) PostFiles(files []FileInfo, params MapStrings) ([]byte, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	for _, file := range files {
		stat, err := file.File.Stat()
		if err != nil {
			return nil, err
		}
		if c.debug {
			c.log.Debug("Client.Upload fileName=%s fileSize=%v fileMode=%v fileModTime=%s", stat.Name(), stat.Size(), stat.Mode(), stat.ModTime())
		}
		fileWriter, err := bodyWriter.CreateFormFile(file.Name, stat.Name())
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(fileWriter, file.File)
		if err != nil {
			return nil, err
		}
	}
	c.SetHeader(HeaderContentTypeKey, bodyWriter.FormDataContentType())
	_ = bodyWriter.Close()
	for key, val := range params {
		_ = bodyWriter.WriteField(key, val)
	}
	c.BuildUrl(c.QueryParam)
	return c.Do(http.MethodPost, c.url.String(), bodyBuf, nil, nil)
}

// PostFile 单文件上传
// file, _ := os.OpenFile("test.md", os.O_RDONLY, os.ModePerm)
// defer file.Close()
// UploadOne("file",file,nil)
func (c *Client) PostFile(name string, file *os.File, params MapStrings) ([]byte, error) {
	return c.PostFiles([]FileInfo{{Name: name, File: file}}, params)
}

// PostForm 表单提交
func (c *Client) PostForm(params url.Values) ([]byte, error) {
	c.header(HeaderContentTypeKey, FormContentType)
	c.BuildUrl(c.QueryParam)
	return c.Do(http.MethodPost, c.url.String(), strings.NewReader(params.Encode()), nil, nil)
}

// PostJson json提交
func (c *Client) PostJson(body any) ([]byte, error) {
	c.header(HeaderContentTypeKey, JsonContentType)
	var jsonStr string
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
		c.BuildUrl(c.QueryParam)
		return c.Do(http.MethodPost, c.url.String(), strings.NewReader(jsonStr), nil, nil)
	}
	return nil, nil
}

// Do 所有的请求都可以走这个方法
func (c *Client) Do(method, url string, body io.Reader, header MapStrings, cookie MapStrings) ([]byte, error) {
	var err error
	c.Request, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
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
		c.log.Debug("Client.Do.Request %s %s", method, url)
		c.log.Debug("Client.Do.Request Header  %s", headers)
		c.log.Debug("Client.Do.Request Cookie  %s", cookies)
	}
	c.Response, err = c.httpClient.Do(c.Request)
	if err != nil {
		c.log.Error("client.Do.httpClient.Do err=%v", err)
		return nil, err
	}
	defer c.Response.Body.Close()
	bodyByte, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		c.log.Error("Client.Request.ioutil.ReadAll err=%v", err)
		return nil, err
	}
	if c.debug {
		c.log.Debug("Client.Do.Response %s %s", c.Response.Status, method)
		c.log.Debug("Client.Do.Response Header  %s", c.Response.Header)
		c.log.Debug("Client.Do.Response Cookie  %s", c.Response.Cookies())
		c.log.Debug("Client.Do.Response body  %s", string(bodyByte))
	}
	return bodyByte, err
}
