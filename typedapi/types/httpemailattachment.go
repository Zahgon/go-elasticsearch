package types

type HttpEmailAttachment struct {
	ContentType *string                     `json:"content_type,omitempty"`
	Inline      *bool                       `json:"inline,omitempty"`
	Request     *HttpInputRequestDefinition `json:"request,omitempty"`
}

func (s *HttpEmailAttachment) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHttpEmailAttachment() *HttpEmailAttachment { _ = "STUB: not implemented"; return nil }

type HttpEmailAttachmentVariant interface {
	HttpEmailAttachmentCaster() *HttpEmailAttachment
}

func (s *HttpEmailAttachment) HttpEmailAttachmentCaster() *HttpEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}
