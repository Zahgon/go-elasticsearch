package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectionscheme"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/httpinputmethod"
)

type _httpInputRequestDefinition struct {
	v *types.HttpInputRequestDefinition
}

func NewHttpInputRequestDefinition() *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Auth(auth types.HttpInputAuthenticationVariant) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Body(body string) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) ConnectionTimeout(duration types.DurationVariant) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Headers(headers map[string]string) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) AddHeader(key string, value string) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Host(host string) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Method(method httpinputmethod.HttpInputMethod) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Params(params map[string]string) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) AddParam(key string, value string) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Path(path string) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Port(port uint) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Proxy(proxy types.HttpInputProxyVariant) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) ReadTimeout(duration types.DurationVariant) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Scheme(scheme connectionscheme.ConnectionScheme) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) Url(url string) *_httpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInputRequestDefinition) HttpInputRequestDefinitionCaster() *types.HttpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}
