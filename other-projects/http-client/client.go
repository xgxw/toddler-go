package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-cleanhttp"
)

const (
	BaseURL = "https://localhost.com"

	// ContentTypeJSON HTTP Content-Type header for JSON data
	ContentTypeJSON = "application/json; charset=utf-8"
)

type Options struct {
	Token      string `json:"token"`
	HTTPClient *http.Client
}

type Client struct {
	BaseURL *url.URL `json:"base_url"`
	Token   string   `json:"token"`

	client *http.Client
}

func NewClient(opts Options) *Client {
	baseURL, _ := url.Parse(BaseURL)

	httpc := opts.HTTPClient
	if httpc == nil {
		httpc = cleanhttp.DefaultPooledClient()
	}

	return &Client{
		BaseURL: baseURL,
		Token:   opts.Token,
		client:  httpc,
	}
}

// NewRequest 创建一个 Request. 通过 body 实现的方法判断对Request进行哪些定制
func (c *Client) NewRequest(method, uri string, params url.Values, body interface{}) (*http.Request, error) {
	var err error

	// 如果body继承了AuthorizedRequester, 那么params带上token
	if _, ok := body.(AuthorizedRequester); ok {
		params, err = c.paramsWithToken(params)
		if err != nil {
			return nil, err
		}
	}

	u, err := c.resolveURL(uri, params)
	if err != nil {
		return nil, err
	}

	var isJSON bool
	var buf io.ReadWriter
	// 如果body继承了EmptyBodyRequester, 则调整encoder, 并且判断 ContentType
	if _, ok := body.(EmptyBodyRequester); body != nil && !ok {
		buf = new(bytes.Buffer)
		encoder := json.NewEncoder(buf)

		if b, ok := body.(CustomizedEncoderRequester); ok {
			b.ConfigureEncoder(encoder)
		}

		err = encoder.Encode(body)
		if err != nil {
			return nil, err
		}

		isJSON = true
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if isJSON {
		req.Header.Set("Content-Type", ContentTypeJSON)
	}
	return req, nil
}

// Do 发起HTTP调用
func (c *Client) Do(ctx context.Context, req *http.Request, v ResponseCarrier) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	if err == nil {
		// 接口响应的错误一并返回 由调用方处理
		if v != nil && v.GetErrorCode() != 0 {
			err = &ResponseError{
				Code:     v.GetErrorCode(),
				Message:  v.GetErrorMessage(),
				Response: resp,
			}
		}
	}

	return resp, err
}

// resolveURL 处理URL请求路径和query参数
func (c *Client) resolveURL(uri string, params url.Values) (*url.URL, error) {
	rel, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	if params != nil {
		u.RawQuery = params.Encode()
	}

	return u, nil
}

// paramsWithToken 在query参数中增加access_token
func (c *Client) paramsWithToken(params url.Values) (url.Values, error) {
	if params == nil {
		params = url.Values{}
	}

	params.Set("token", c.Token)

	return params, nil
}
