package types

type ReportingEmailAttachment struct {
	Inline   *bool                       `json:"inline,omitempty"`
	Interval Duration                    `json:"interval,omitempty"`
	Request  *HttpInputRequestDefinition `json:"request,omitempty"`
	Retries  *int                        `json:"retries,omitempty"`
	Url      string                      `json:"url"`
}

func (s *ReportingEmailAttachment) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewReportingEmailAttachment() *ReportingEmailAttachment { _ = "STUB: not implemented"; return nil }

type ReportingEmailAttachmentVariant interface {
	ReportingEmailAttachmentCaster() *ReportingEmailAttachment
}

func (s *ReportingEmailAttachment) ReportingEmailAttachmentCaster() *ReportingEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}
