package types

type EmailBody struct {
	Html *string `json:"html,omitempty"`
	Text *string `json:"text,omitempty"`
}

func (s *EmailBody) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEmailBody() *EmailBody { _ = "STUB: not implemented"; return nil }

type EmailBodyVariant interface {
	EmailBodyCaster() *EmailBody
}

func (s *EmailBody) EmailBodyCaster() *EmailBody { _ = "STUB: not implemented"; return nil }
