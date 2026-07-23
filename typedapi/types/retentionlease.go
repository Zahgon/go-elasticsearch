package types

type RetentionLease struct {
	Period Duration `json:"period"`
}

func (s *RetentionLease) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRetentionLease() *RetentionLease { _ = "STUB: not implemented"; return nil }

type RetentionLeaseVariant interface {
	RetentionLeaseCaster() *RetentionLease
}

func (s *RetentionLease) RetentionLeaseCaster() *RetentionLease {
	_ = "STUB: not implemented"
	return nil
}
