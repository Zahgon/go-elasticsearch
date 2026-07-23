package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _httpInputProxy struct {
	v *types.HttpInputProxy
}

func NewHttpInputProxy(port uint) *_httpInputProxy { _ = "STUB: not implemented"; return nil }

func (s *_httpInputProxy) Host(host string) *_httpInputProxy { _ = "STUB: not implemented"; return nil }

func (s *_httpInputProxy) Port(port uint) *_httpInputProxy { _ = "STUB: not implemented"; return nil }

func (s *_httpInputProxy) HttpInputProxyCaster() *types.HttpInputProxy {
	_ = "STUB: not implemented"
	return nil
}
