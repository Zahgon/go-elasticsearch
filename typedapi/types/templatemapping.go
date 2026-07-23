package types

import (
	"encoding/json"
)

type TemplateMapping struct {
	Aliases       map[string]Alias           `json:"aliases"`
	IndexPatterns []string                   `json:"index_patterns"`
	Mappings      TypeMapping                `json:"mappings"`
	Order         int                        `json:"order"`
	Settings      map[string]json.RawMessage `json:"settings"`
	Version       *int64                     `json:"version,omitempty"`
}

func (s *TemplateMapping) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTemplateMapping() *TemplateMapping { _ = "STUB: not implemented"; return nil }
