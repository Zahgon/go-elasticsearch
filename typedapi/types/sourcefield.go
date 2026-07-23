package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sourcefieldmode"
)

type SourceField struct {
	Compress          *bool                            `json:"compress,omitempty"`
	CompressThreshold *string                          `json:"compress_threshold,omitempty"`
	Enabled           *bool                            `json:"enabled,omitempty"`
	Excludes          []string                         `json:"excludes,omitempty"`
	Includes          []string                         `json:"includes,omitempty"`
	Mode              *sourcefieldmode.SourceFieldMode `json:"mode,omitempty"`
}

func (s *SourceField) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSourceField() *SourceField { _ = "STUB: not implemented"; return nil }

type SourceFieldVariant interface {
	SourceFieldCaster() *SourceField
}

func (s *SourceField) SourceFieldCaster() *SourceField { _ = "STUB: not implemented"; return nil }
