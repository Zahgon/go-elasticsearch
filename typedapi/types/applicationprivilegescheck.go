package types

type ApplicationPrivilegesCheck struct {
	Application string `json:"application"`

	Privileges []string `json:"privileges"`

	Resources []string `json:"resources"`
}

func (s *ApplicationPrivilegesCheck) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewApplicationPrivilegesCheck() *ApplicationPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}

type ApplicationPrivilegesCheckVariant interface {
	ApplicationPrivilegesCheckCaster() *ApplicationPrivilegesCheck
}

func (s *ApplicationPrivilegesCheck) ApplicationPrivilegesCheckCaster() *ApplicationPrivilegesCheck {
	_ = "STUB: not implemented"
	return nil
}
