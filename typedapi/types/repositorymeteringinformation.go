package types

type RepositoryMeteringInformation struct {
	Archived bool `json:"archived"`

	ClusterVersion *int64 `json:"cluster_version,omitempty"`

	RepositoryEphemeralId string `json:"repository_ephemeral_id"`

	RepositoryLocation RepositoryLocation `json:"repository_location"`

	RepositoryName string `json:"repository_name"`

	RepositoryStartedAt int64 `json:"repository_started_at"`

	RepositoryStoppedAt *int64 `json:"repository_stopped_at,omitempty"`

	RepositoryType string `json:"repository_type"`

	RequestCounts RequestCounts `json:"request_counts"`
}

func (s *RepositoryMeteringInformation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRepositoryMeteringInformation() *RepositoryMeteringInformation {
	_ = "STUB: not implemented"
	return nil
}
