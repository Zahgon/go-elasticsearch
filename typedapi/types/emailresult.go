package types

type EmailResult struct {
	Account *string `json:"account,omitempty"`
	Message Email   `json:"message"`
	Reason  *string `json:"reason,omitempty"`
}

func (s *EmailResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEmailResult() *EmailResult { _ = "STUB: not implemented"; return nil }
