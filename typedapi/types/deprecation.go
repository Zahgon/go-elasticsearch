package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/deprecationlevel"
)

type Deprecation struct {
	Details *string `json:"details,omitempty"`

	Level deprecationlevel.DeprecationLevel `json:"level"`

	Message                     string                     `json:"message"`
	Meta_                       map[string]json.RawMessage `json:"_meta,omitempty"`
	ResolveDuringRollingUpgrade bool                       `json:"resolve_during_rolling_upgrade"`

	Url string `json:"url"`
}

func (s *Deprecation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDeprecation() *Deprecation { _ = "STUB: not implemented"; return nil }
