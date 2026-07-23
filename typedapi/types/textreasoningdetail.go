package types

type TextReasoningDetail struct {
	Format *string `json:"format,omitempty"`

	Id *string `json:"id,omitempty"`

	Index *int `json:"index,omitempty"`

	Signature *string `json:"signature,omitempty"`

	Text *string `json:"text,omitempty"`
	Type string  `json:"type,omitempty"`
}

func (s *TextReasoningDetail) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s TextReasoningDetail) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTextReasoningDetail() *TextReasoningDetail { _ = "STUB: not implemented"; return nil }

type TextReasoningDetailVariant interface {
	TextReasoningDetailCaster() *TextReasoningDetail
}

func (s *TextReasoningDetail) TextReasoningDetailCaster() *TextReasoningDetail {
	_ = "STUB: not implemented"
	return nil
}

func (s *TextReasoningDetail) ReasoningDetailCaster() *ReasoningDetail {
	_ = "STUB: not implemented"
	return nil
}
