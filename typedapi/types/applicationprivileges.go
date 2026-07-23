package types

type ApplicationPrivileges struct {
	Application string `json:"application"`

	Privileges []string `json:"privileges"`

	Resources []string `json:"resources"`
}

func (s *ApplicationPrivileges) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewApplicationPrivileges() *ApplicationPrivileges { _ = "STUB: not implemented"; return nil }

type ApplicationPrivilegesVariant interface {
	ApplicationPrivilegesCaster() *ApplicationPrivileges
}

func (s *ApplicationPrivileges) ApplicationPrivilegesCaster() *ApplicationPrivileges {
	_ = "STUB: not implemented"
	return nil
}
