package types

type ReadOnlyUrlRepository struct {
	Settings ReadOnlyUrlRepositorySettings `json:"settings"`

	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

func (s *ReadOnlyUrlRepository) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ReadOnlyUrlRepository) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewReadOnlyUrlRepository() *ReadOnlyUrlRepository { _ = "STUB: not implemented"; return nil }

type ReadOnlyUrlRepositoryVariant interface {
	ReadOnlyUrlRepositoryCaster() *ReadOnlyUrlRepository
}

func (s *ReadOnlyUrlRepository) ReadOnlyUrlRepositoryCaster() *ReadOnlyUrlRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReadOnlyUrlRepository) RepositoryCaster() *Repository {
	_ = "STUB: not implemented"
	return nil
}
