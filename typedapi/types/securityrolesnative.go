package types

type SecurityRolesNative struct {
	Dls  bool  `json:"dls"`
	Fls  bool  `json:"fls"`
	Size int64 `json:"size"`
}

func (s *SecurityRolesNative) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSecurityRolesNative() *SecurityRolesNative { _ = "STUB: not implemented"; return nil }
