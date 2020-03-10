package httpclient

import "encoding/json"

// EmptyBodyRequester 接口用于判断请求体是否为空
type EmptyBodyRequester interface {
	EmptyBody()
}

type EmptyBodyRequest struct{}

func (b *EmptyBodyRequest) EmptyBody() {}

// AuthorizedRequester 接口用于判断请求是否需要提供AccessToken
type AuthorizedRequester interface {
	Authorized()
}

type AuthorizedRequest struct{}

func (a *AuthorizedRequest) Authorized() {}

type CustomizedEncoderRequester interface {
	ConfigureEncoder(encoder *json.Encoder)
}
