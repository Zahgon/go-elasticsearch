package types

type RetentionPolicy struct {
	Field string `json:"field"`

	MaxAge Duration `json:"max_age"`
}

func (s *RetentionPolicy) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRetentionPolicy() *RetentionPolicy { _ = "STUB: not implemented"; return nil }

type RetentionPolicyVariant interface {
	RetentionPolicyCaster() *RetentionPolicy
}

func (s *RetentionPolicy) RetentionPolicyCaster() *RetentionPolicy {
	_ = "STUB: not implemented"
	return nil
}
