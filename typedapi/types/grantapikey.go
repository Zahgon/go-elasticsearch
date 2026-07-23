package types

type GrantApiKey struct {
	Expiration *string `json:"expiration,omitempty"`

	Metadata Metadata `json:"metadata,omitempty"`
	Name     string   `json:"name"`

	RoleDescriptors []map[string]RoleDescriptor `json:"role_descriptors,omitempty"`
}

func (s *GrantApiKey) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGrantApiKey() *GrantApiKey { _ = "STUB: not implemented"; return nil }

type GrantApiKeyVariant interface {
	GrantApiKeyCaster() *GrantApiKey
}

func (s *GrantApiKey) GrantApiKeyCaster() *GrantApiKey { _ = "STUB: not implemented"; return nil }
