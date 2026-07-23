package types

type PublishedClusterStates struct {
	CompatibleDiffs *int64 `json:"compatible_diffs,omitempty"`

	FullStates *int64 `json:"full_states,omitempty"`

	IncompatibleDiffs *int64 `json:"incompatible_diffs,omitempty"`
}

func (s *PublishedClusterStates) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPublishedClusterStates() *PublishedClusterStates { _ = "STUB: not implemented"; return nil }
