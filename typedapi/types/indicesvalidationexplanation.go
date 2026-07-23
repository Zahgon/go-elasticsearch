package types

type IndicesValidationExplanation struct {
	Error       *string `json:"error,omitempty"`
	Explanation *string `json:"explanation,omitempty"`
	Index       *string `json:"index,omitempty"`
	Shard       *int    `json:"shard,omitempty"`
	Valid       bool    `json:"valid"`
}

func (s *IndicesValidationExplanation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndicesValidationExplanation() *IndicesValidationExplanation {
	_ = "STUB: not implemented"
	return nil
}
