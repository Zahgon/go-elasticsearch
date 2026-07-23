package types

type RetentionPolicyContainer struct {
	Time *RetentionPolicy `json:"time,omitempty"`
}

func NewRetentionPolicyContainer() *RetentionPolicyContainer { _ = "STUB: not implemented"; return nil }

type RetentionPolicyContainerVariant interface {
	RetentionPolicyContainerCaster() *RetentionPolicyContainer
}

func (s *RetentionPolicyContainer) RetentionPolicyContainerCaster() *RetentionPolicyContainer {
	_ = "STUB: not implemented"
	return nil
}
