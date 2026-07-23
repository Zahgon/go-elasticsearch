package types

type ResolveClusterInfo struct {
	Connected bool `json:"connected"`

	Error *string `json:"error,omitempty"`

	MatchingIndices *bool `json:"matching_indices,omitempty"`

	SkipUnavailable bool `json:"skip_unavailable"`

	Version *ElasticsearchVersionMinInfo `json:"version,omitempty"`
}

func (s *ResolveClusterInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewResolveClusterInfo() *ResolveClusterInfo { _ = "STUB: not implemented"; return nil }
