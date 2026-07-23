package types

type Feature struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

func (s *Feature) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFeature() *Feature { _ = "STUB: not implemented"; return nil }
