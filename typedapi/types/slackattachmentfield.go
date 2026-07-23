package types

type SlackAttachmentField struct {
	Int   bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

func (s *SlackAttachmentField) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSlackAttachmentField() *SlackAttachmentField { _ = "STUB: not implemented"; return nil }

type SlackAttachmentFieldVariant interface {
	SlackAttachmentFieldCaster() *SlackAttachmentField
}

func (s *SlackAttachmentField) SlackAttachmentFieldCaster() *SlackAttachmentField {
	_ = "STUB: not implemented"
	return nil
}
