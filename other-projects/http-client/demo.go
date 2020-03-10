package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type DemoRequest struct {
	Limit int
}

// 实现继承, 从而在NewRequest时添加相应属性
func (d DemoRequest) EmptyBody()  {}
func (d DemoRequest) Authorized() {}

func (d *DemoRequest) GetParams() url.Values {
	params := url.Values{}
	params.Set("limit", fmt.Sprint(d.Limit))
	return params
}

type DemoResponse struct {
	Response
	Words string `json:"words"`
}

func (c *Client) GetDemo(ctx context.Context, request *DemoRequest) (*DemoResponse, *http.Response, error) {
	params := request.GetParams()
	req, err := c.NewRequest(http.MethodGet, "/", params, request)
	if err != nil {
		return &DemoResponse{}, new(http.Response), err
	}

	result := new(DemoResponse)
	httpResp, err := c.Do(ctx, req, result)
	return result, httpResp, nil
}
