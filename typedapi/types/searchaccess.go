package types

type SearchAccess struct {
	AllowRestrictedIndices *bool `json:"allow_restricted_indices,omitempty"`

	FieldSecurity *FieldSecurity `json:"field_security,omitempty"`

	Names []string `json:"names"`

	Query IndicesPrivilegesQuery `json:"query,omitempty"`
}

func (s *SearchAccess) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSearchAccess() *SearchAccess { _ = "STUB: not implemented"; return nil }

type SearchAccessVariant interface {
	SearchAccessCaster() *SearchAccess
}

func (s *SearchAccess) SearchAccessCaster() *SearchAccess { _ = "STUB: not implemented"; return nil }
