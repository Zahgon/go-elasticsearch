package types

type SoftDeletes struct {
	Enabled *bool `json:"enabled,omitempty"`

	RetentionLease *RetentionLease `json:"retention_lease,omitempty"`
}

func (s *SoftDeletes) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSoftDeletes() *SoftDeletes { _ = "STUB: not implemented"; return nil }

type SoftDeletesVariant interface {
	SoftDeletesCaster() *SoftDeletes
}

func (s *SoftDeletes) SoftDeletesCaster() *SoftDeletes { _ = "STUB: not implemented"; return nil }
