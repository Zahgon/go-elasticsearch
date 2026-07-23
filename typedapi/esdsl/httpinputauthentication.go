package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _httpInputAuthentication struct {
	v *types.HttpInputAuthentication
}

func NewHttpInputAuthentication(basic types.HttpInputBasicAuthenticationVariant) *_httpInputAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputAuthentication) Basic(basic types.HttpInputBasicAuthenticationVariant) *_httpInputAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputAuthentication) HttpInputAuthenticationCaster() *types.HttpInputAuthentication {
	_ = "STUB: not implemented"
	return nil
}
