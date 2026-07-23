package types

type Overlapping struct {
	IndexPatterns []string `json:"index_patterns"`
	Name          string   `json:"name"`
}

func (s *Overlapping) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewOverlapping() *Overlapping { _ = "STUB: not implemented"; return nil }
