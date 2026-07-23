package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _httpHeaders struct {
	v types.HttpHeaders
}

func NewHttpHeaders(httpheaders map[string][]string) *_httpHeaders {
	_ = "STUB: not implemented"
	return nil
}

func (u *_httpHeaders) HttpHeadersCaster() *types.HttpHeaders {
	_ = "STUB: not implemented"
	return nil
}
