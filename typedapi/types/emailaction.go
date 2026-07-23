package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/emailpriority"
)

type EmailAction struct {
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

func (s *EmailAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEmailAction() *EmailAction { _ = "STUB: not implemented"; return nil }

type EmailActionVariant interface {
	EmailActionCaster() *EmailAction
}

func (s *EmailAction) EmailActionCaster() *EmailAction { _ = "STUB: not implemented"; return nil }
