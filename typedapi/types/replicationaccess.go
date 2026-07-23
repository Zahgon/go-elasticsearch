package types

type ReplicationAccess struct {
	AllowRestrictedIndices *bool `json:"allow_restricted_indices,omitempty"`

	Names []string `json:"names"`
}

func (s *ReplicationAccess) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReplicationAccess() *ReplicationAccess { _ = "STUB: not implemented"; return nil }

type ReplicationAccessVariant interface {
	ReplicationAccessCaster() *ReplicationAccess
}

func (s *ReplicationAccess) ReplicationAccessCaster() *ReplicationAccess {
	_ = "STUB: not implemented"
	return nil
}
