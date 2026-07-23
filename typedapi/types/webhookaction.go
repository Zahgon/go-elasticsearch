package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectionscheme"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/httpinputmethod"
)

type WebhookAction struct {
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

func (s *WebhookAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWebhookAction() *WebhookAction { _ = "STUB: not implemented"; return nil }

type WebhookActionVariant interface {
	WebhookActionCaster() *WebhookAction
}

func (s *WebhookAction) WebhookActionCaster() *WebhookAction { _ = "STUB: not implemented"; return nil }
