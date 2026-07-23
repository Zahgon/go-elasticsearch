package types

type EmailAttachmentContainer struct {
	Data      *DataEmailAttachment      `json:"data,omitempty"`
	Http      *HttpEmailAttachment      `json:"http,omitempty"`
	Reporting *ReportingEmailAttachment `json:"reporting,omitempty"`
}

func NewEmailAttachmentContainer() *EmailAttachmentContainer { _ = "STUB: not implemented"; return nil }

type EmailAttachmentContainerVariant interface {
	EmailAttachmentContainerCaster() *EmailAttachmentContainer
}

func (s *EmailAttachmentContainer) EmailAttachmentContainerCaster() *EmailAttachmentContainer {
	_ = "STUB: not implemented"
	return nil
}
