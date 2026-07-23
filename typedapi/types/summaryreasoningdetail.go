package types

type SummaryReasoningDetail struct {
	Format *string `json:"format,omitempty"`

	Id *string `json:"id,omitempty"`

	Index *int `json:"index,omitempty"`

	Summary string `json:"summary"`
	Type    string `json:"type,omitempty"`
}

func (s *SummaryReasoningDetail) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SummaryReasoningDetail) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSummaryReasoningDetail() *SummaryReasoningDetail { _ = "STUB: not implemented"; return nil }

type SummaryReasoningDetailVariant interface {
	SummaryReasoningDetailCaster() *SummaryReasoningDetail
}

func (s *SummaryReasoningDetail) SummaryReasoningDetailCaster() *SummaryReasoningDetail {
	_ = "STUB: not implemented"
	return nil
}

func (s *SummaryReasoningDetail) ReasoningDetailCaster() *ReasoningDetail {
	_ = "STUB: not implemented"
	return nil
}
