package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type LikeDocument struct {
	Doc    json.RawMessage `json:"doc,omitempty"`
	Fields []string        `json:"fields,omitempty"`

	Id_ *string `json:"_id,omitempty"`

	Index_ *string `json:"_index,omitempty"`

	PerFieldAnalyzer map[string]string        `json:"per_field_analyzer,omitempty"`
	Routing          []string                 `json:"routing,omitempty"`
	Version          *int64                   `json:"version,omitempty"`
	VersionType      *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *LikeDocument) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLikeDocument() *LikeDocument { _ = "STUB: not implemented"; return nil }

type LikeDocumentVariant interface {
	LikeDocumentCaster() *LikeDocument
}

func (s *LikeDocument) LikeDocumentCaster() *LikeDocument { _ = "STUB: not implemented"; return nil }

func (s *LikeDocument) LikeCaster() *Like { _ = "STUB: not implemented"; return nil }
