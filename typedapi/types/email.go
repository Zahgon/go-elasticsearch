package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/emailpriority"
)

type Email struct {
	Attachments map[string]EmailAttachmentContainer `json:"attachments,omitempty"`
	Bcc         []string                            `json:"bcc,omitempty"`
	Body        *EmailBody                          `json:"body,omitempty"`
	Cc          []string                            `json:"cc,omitempty"`
	From        *string                             `json:"from,omitempty"`
	Id          *string                             `json:"id,omitempty"`
	Priority    *emailpriority.EmailPriority        `json:"priority,omitempty"`
	ReplyTo     []string                            `json:"reply_to,omitempty"`
	SentDate    DateTime                            `json:"sent_date,omitempty"`
	Subject     string                              `json:"subject"`
	To          []string                            `json:"to"`
}

func (s *Email) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEmail() *Email { _ = "STUB: not implemented"; return nil }
