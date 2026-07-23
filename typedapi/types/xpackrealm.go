package types

type XpackRealm struct {
	Available                 bool         `json:"available"`
	Cache                     []RealmCache `json:"cache,omitempty"`
	Enabled                   bool         `json:"enabled"`
	HasAuthorizationRealms    []bool       `json:"has_authorization_realms,omitempty"`
	HasDefaultUsernamePattern []bool       `json:"has_default_username_pattern,omitempty"`
	HasTruststore             []bool       `json:"has_truststore,omitempty"`
	IsAuthenticationDelegated []bool       `json:"is_authentication_delegated,omitempty"`
	Name                      []string     `json:"name,omitempty"`
	Order                     []int64      `json:"order,omitempty"`
	Size                      []int64      `json:"size,omitempty"`
}

func (s *XpackRealm) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewXpackRealm() *XpackRealm { _ = "STUB: not implemented"; return nil }
