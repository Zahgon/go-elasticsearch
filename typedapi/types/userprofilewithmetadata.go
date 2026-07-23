package types

import (
	"encoding/json"
)

type UserProfileWithMetadata struct {
	Data             map[string]json.RawMessage `json:"data"`
	Doc_             UserProfileHitMetadata     `json:"_doc"`
	Enabled          *bool                      `json:"enabled,omitempty"`
	Labels           map[string]json.RawMessage `json:"labels"`
	LastSynchronized int64                      `json:"last_synchronized"`
	Uid              string                     `json:"uid"`
	User             UserProfileUser            `json:"user"`
}

func (s *UserProfileWithMetadata) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUserProfileWithMetadata() *UserProfileWithMetadata { _ = "STUB: not implemented"; return nil }
