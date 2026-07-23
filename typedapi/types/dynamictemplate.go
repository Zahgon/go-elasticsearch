package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/matchtype"
)

type DynamicTemplate struct {
	Mapping            Property             `json:"mapping,omitempty"`
	Match              []string             `json:"match,omitempty"`
	MatchMappingType   []string             `json:"match_mapping_type,omitempty"`
	MatchPattern       *matchtype.MatchType `json:"match_pattern,omitempty"`
	PathMatch          []string             `json:"path_match,omitempty"`
	PathUnmatch        []string             `json:"path_unmatch,omitempty"`
	Runtime            *RuntimeField        `json:"runtime,omitempty"`
	Unmatch            []string             `json:"unmatch,omitempty"`
	UnmatchMappingType []string             `json:"unmatch_mapping_type,omitempty"`
}

func (s *DynamicTemplate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDynamicTemplate() *DynamicTemplate { _ = "STUB: not implemented"; return nil }

type DynamicTemplateVariant interface {
	DynamicTemplateCaster() *DynamicTemplate
}

func (s *DynamicTemplate) DynamicTemplateCaster() *DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}
