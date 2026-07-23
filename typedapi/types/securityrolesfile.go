package types

type SecurityRolesFile struct {
	Dls  bool  `json:"dls"`
	Fls  bool  `json:"fls"`
	Size int64 `json:"size"`
}

func (s *SecurityRolesFile) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSecurityRolesFile() *SecurityRolesFile { _ = "STUB: not implemented"; return nil }
