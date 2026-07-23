package types

type WebhookResult struct {
	Request  HttpInputRequestResult   `json:"request"`
	Response *HttpInputResponseResult `json:"response,omitempty"`
}

func NewWebhookResult() *WebhookResult { _ = "STUB: not implemented"; return nil }
