package types

type SharedFileSystemRepository struct {
	Settings SharedFileSystemRepositorySettings `json:"settings"`

	Type string  `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

func (s *SharedFileSystemRepository) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SharedFileSystemRepository) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSharedFileSystemRepository() *SharedFileSystemRepository {
	_ = "STUB: not implemented"
	return nil
}

type SharedFileSystemRepositoryVariant interface {
	SharedFileSystemRepositoryCaster() *SharedFileSystemRepository
}

func (s *SharedFileSystemRepository) SharedFileSystemRepositoryCaster() *SharedFileSystemRepository {
	_ = "STUB: not implemented"
	return nil
}

func (s *SharedFileSystemRepository) RepositoryCaster() *Repository {
	_ = "STUB: not implemented"
	return nil
}
