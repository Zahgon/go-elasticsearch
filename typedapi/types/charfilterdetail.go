package types

type CharFilterDetail struct {
	FilteredText []string `json:"filtered_text"`
	Name         string   `json:"name"`
}

func (s *CharFilterDetail) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCharFilterDetail() *CharFilterDetail { _ = "STUB: not implemented"; return nil }
