package types

type SourceOnlyRepository struct {
	Settings SourceOnlyRepositorySettings `json:"settings"`

	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

func (s *SourceOnlyRepository) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SourceOnlyRepository) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSourceOnlyRepository() *SourceOnlyRepository { _ = "STUB: not implemented"; return nil }

type SourceOnlyRepositoryVariant interface {
	SourceOnlyRepositoryCaster() *SourceOnlyRepository
}

func (s *SourceOnlyRepository) SourceOnlyRepositoryCaster() *SourceOnlyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *SourceOnlyRepository) RepositoryCaster() *Repository {
	_ = "STUB: not implemented"
	return nil
}
