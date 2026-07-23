package types

type RoleDescriptorWrapper struct {
	RoleDescriptor RoleDescriptorRead `json:"role_descriptor"`
}

func NewRoleDescriptorWrapper() *RoleDescriptorWrapper { _ = "STUB: not implemented"; return nil }
