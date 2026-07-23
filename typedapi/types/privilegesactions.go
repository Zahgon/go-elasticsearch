package types

type PrivilegesActions struct {
	Actions     []string `json:"actions"`
	Application *string  `json:"application,omitempty"`
	Metadata    Metadata `json:"metadata,omitempty"`
	Name        *string  `json:"name,omitempty"`
}

func (s *PrivilegesActions) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPrivilegesActions() *PrivilegesActions { _ = "STUB: not implemented"; return nil }

type PrivilegesActionsVariant interface {
	PrivilegesActionsCaster() *PrivilegesActions
}

func (s *PrivilegesActions) PrivilegesActionsCaster() *PrivilegesActions {
	_ = "STUB: not implemented"
	return nil
}
