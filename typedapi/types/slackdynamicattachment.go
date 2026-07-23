package types

type SlackDynamicAttachment struct {
	AttachmentTemplate SlackAttachment `json:"attachment_template"`
	ListPath           string          `json:"list_path"`
}

func (s *SlackDynamicAttachment) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSlackDynamicAttachment() *SlackDynamicAttachment { _ = "STUB: not implemented"; return nil }

type SlackDynamicAttachmentVariant interface {
	SlackDynamicAttachmentCaster() *SlackDynamicAttachment
}

func (s *SlackDynamicAttachment) SlackDynamicAttachmentCaster() *SlackDynamicAttachment {
	_ = "STUB: not implemented"
	return nil
}
