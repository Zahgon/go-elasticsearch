package types

type IlmPolicy struct {
	Meta_  Metadata `json:"_meta,omitempty"`
	Phases Phases   `json:"phases"`
}

func (s *IlmPolicy) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIlmPolicy() *IlmPolicy { _ = "STUB: not implemented"; return nil }

type IlmPolicyVariant interface {
	IlmPolicyCaster() *IlmPolicy
}

func (s *IlmPolicy) IlmPolicyCaster() *IlmPolicy { _ = "STUB: not implemented"; return nil }
