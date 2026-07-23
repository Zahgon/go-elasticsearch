package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/apikeytype"
)

type ApiKey struct {
	Access *Access `json:"access,omitempty"`

	CertificateIdentity *string `json:"certificate_identity,omitempty"`

	Creation int64 `json:"creation"`

	Expiration *int64 `json:"expiration,omitempty"`

	Id string `json:"id"`

	Invalidated bool `json:"invalidated"`

	Invalidation *int64 `json:"invalidation,omitempty"`

	LimitedBy []map[string]RoleDescriptor `json:"limited_by,omitempty"`

	Metadata Metadata `json:"metadata"`

	Name string `json:"name"`

	ProfileUid *string `json:"profile_uid,omitempty"`

	Realm string `json:"realm"`

	RealmType *string `json:"realm_type,omitempty"`

	RoleDescriptors map[string]RoleDescriptor `json:"role_descriptors,omitempty"`

	Sort_ []FieldValue `json:"_sort,omitempty"`

	Type apikeytype.ApiKeyType `json:"type"`

	Username string `json:"username"`
}

func (s *ApiKey) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewApiKey() *ApiKey { _ = "STUB: not implemented"; return nil }
