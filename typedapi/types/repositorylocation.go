package types

type RepositoryLocation struct {
	BasePath string `json:"base_path"`

	Bucket *string `json:"bucket,omitempty"`

	Container *string `json:"container,omitempty"`
}

func (s *RepositoryLocation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRepositoryLocation() *RepositoryLocation { _ = "STUB: not implemented"; return nil }
