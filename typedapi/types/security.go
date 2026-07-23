package types

type Security struct {
	Anonymous          FeatureToggle               `json:"anonymous"`
	ApiKeyService      FeatureToggle               `json:"api_key_service"`
	Audit              Audit                       `json:"audit"`
	Available          bool                        `json:"available"`
	Enabled            bool                        `json:"enabled"`
	Fips140            FeatureToggle               `json:"fips_140"`
	Ipfilter           IpFilter                    `json:"ipfilter"`
	OperatorPrivileges Base                        `json:"operator_privileges"`
	Realms             map[string]XpackRealm       `json:"realms"`
	RoleMapping        map[string]XpackRoleMapping `json:"role_mapping"`
	Roles              SecurityRoles               `json:"roles"`
	Ssl                Ssl                         `json:"ssl"`
	SystemKey          *FeatureToggle              `json:"system_key,omitempty"`
	TokenService       FeatureToggle               `json:"token_service"`
}

func (s *Security) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSecurity() *Security { _ = "STUB: not implemented"; return nil }
