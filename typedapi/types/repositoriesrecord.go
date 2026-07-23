package types

type RepositoriesRecord struct {
	Id *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (s *RepositoriesRecord) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRepositoriesRecord() *RepositoriesRecord { _ = "STUB: not implemented"; return nil }
