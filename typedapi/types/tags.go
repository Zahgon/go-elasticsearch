package types

type Tags struct {
	Alias_        string            `json:"_alias"`
	Id_           string            `json:"_id"`
	Organisation_ string            `json:"_organisation"`
	Tags          map[string]string `json:"-"`
	Type_         string            `json:"_type"`
}

func (s *Tags) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s Tags) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewTags() *Tags { _ = "STUB: not implemented"; return nil }
