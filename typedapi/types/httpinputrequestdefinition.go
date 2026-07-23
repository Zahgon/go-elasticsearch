package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectionscheme"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/httpinputmethod"
)

type HttpInputRequestDefinition struct {
	Auth              *HttpInputAuthentication           `json:"auth,omitempty"`
	Body              *string                            `json:"body,omitempty"`
	ConnectionTimeout Duration                           `json:"connection_timeout,omitempty"`
	Headers           map[string]string                  `json:"headers,omitempty"`
	Host              *string                            `json:"host,omitempty"`
	Method            *httpinputmethod.HttpInputMethod   `json:"method,omitempty"`
	Params            map[string]string                  `json:"params,omitempty"`
	Path              *string                            `json:"path,omitempty"`
	Port              *uint                              `json:"port,omitempty"`
	Proxy             *HttpInputProxy                    `json:"proxy,omitempty"`
	ReadTimeout       Duration                           `json:"read_timeout,omitempty"`
	Scheme            *connectionscheme.ConnectionScheme `json:"scheme,omitempty"`
	Url               *string                            `json:"url,omitempty"`
}

func (s *HttpInputRequestDefinition) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHttpInputRequestDefinition() *HttpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

type HttpInputRequestDefinitionVariant interface {
	HttpInputRequestDefinitionCaster() *HttpInputRequestDefinition
}

func (s *HttpInputRequestDefinition) HttpInputRequestDefinitionCaster() *HttpInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}
