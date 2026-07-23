package types

type TopClassEntry struct {
	ClassName        string  `json:"class_name"`
	ClassProbability Float64 `json:"class_probability"`
	ClassScore       Float64 `json:"class_score"`
}

func (s *TopClassEntry) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTopClassEntry() *TopClassEntry { _ = "STUB: not implemented"; return nil }
