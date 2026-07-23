package types

type EncryptedReasoningDetail struct {
	Data string `json:"data"`

	Format *string `json:"format,omitempty"`

	Id *string `json:"id,omitempty"`

	Index *int   `json:"index,omitempty"`
	Type  string `json:"type,omitempty"`
}

func (s *EncryptedReasoningDetail) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s EncryptedReasoningDetail) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewEncryptedReasoningDetail() *EncryptedReasoningDetail { _ = "STUB: not implemented"; return nil }

type EncryptedReasoningDetailVariant interface {
	EncryptedReasoningDetailCaster() *EncryptedReasoningDetail
}

func (s *EncryptedReasoningDetail) EncryptedReasoningDetailCaster() *EncryptedReasoningDetail {
	_ = "STUB: not implemented"
	return nil
}

func (s *EncryptedReasoningDetail) ReasoningDetailCaster() *ReasoningDetail {
	_ = "STUB: not implemented"
	return nil
}
