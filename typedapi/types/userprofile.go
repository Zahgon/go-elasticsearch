package types

import (
	"encoding/json"
)

type UserProfile struct {
	Data    map[string]json.RawMessage `json:"data"`
	Enabled *bool                      `json:"enabled,omitempty"`
	Labels  map[string]json.RawMessage `json:"labels"`
	Uid     string                     `json:"uid"`
	User    UserProfileUser            `json:"user"`
}

func (s *UserProfile) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUserProfile() *UserProfile { _ = "STUB: not implemented"; return nil }
