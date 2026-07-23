package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectionscheme"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/httpinputmethod"
)

type _webhookAction struct {
	v *types.WebhookAction
}

func NewWebhookAction() *_webhookAction { _ = "STUB: not implemented"; return nil }

func (s *_webhookAction) Auth(auth types.HttpInputAuthenticationVariant) *_webhookAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_webhookAction) Body(body string) *_webhookAction { _ = "STUB: not implemented"; return nil }

func (s *_webhookAction) ConnectionTimeout(duration types.DurationVariant) *_webhookAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_webhookAction) Headers(headers map[string]string) *_webhookAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_webhookAction) AddHeader(key string, value string) *_webhookAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_webhookAction) Host(host string) *_webhookAction { _ = "STUB: not implemented"; return nil }

func (s *_webhookAction) Method(method httpinputmethod.HttpInputMethod) *_webhookAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_webhookAction) Params(params map[string]string) *_webhookAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_webhookAction) AddParam(key string, value string) *_webhookAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_webhookAction) Path(path string) *_webhookAction { _ = "STUB: not implemented"; return nil }

func (s *_webhookAction) Port(port uint) *_webhookAction { _ = "STUB: not implemented"; return nil }

func (s *_webhookAction) Proxy(proxy types.HttpInputProxyVariant) *_webhookAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_webhookAction) ReadTimeout(duration types.DurationVariant) *_webhookAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_webhookAction) Scheme(scheme connectionscheme.ConnectionScheme) *_webhookAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_webhookAction) Url(url string) *_webhookAction { _ = "STUB: not implemented"; return nil }

func (s *_webhookAction) WebhookActionCaster() *types.WebhookAction {
	_ = "STUB: not implemented"
	return nil
}
