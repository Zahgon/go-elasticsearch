package types

type RepositoryIntegrityIndicatorDetails struct {
	Corrupted             []string `json:"corrupted,omitempty"`
	CorruptedRepositories *int64   `json:"corrupted_repositories,omitempty"`
	TotalRepositories     *int64   `json:"total_repositories,omitempty"`
}

func (s *RepositoryIntegrityIndicatorDetails) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRepositoryIntegrityIndicatorDetails() *RepositoryIntegrityIndicatorDetails {
	_ = "STUB: not implemented"
	return nil
}
