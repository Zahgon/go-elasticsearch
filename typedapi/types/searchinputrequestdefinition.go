package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

type SearchInputRequestDefinition struct {
	Body               *SearchInputRequestBody    `json:"body,omitempty"`
	Indices            []string                   `json:"indices,omitempty"`
	IndicesOptions     *IndicesOptions            `json:"indices_options,omitempty"`
	RestTotalHitsAsInt *bool                      `json:"rest_total_hits_as_int,omitempty"`
	SearchType         *searchtype.SearchType     `json:"search_type,omitempty"`
	Template           *SearchTemplateRequestBody `json:"template,omitempty"`
}

func (s *SearchInputRequestDefinition) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSearchInputRequestDefinition() *SearchInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}

type SearchInputRequestDefinitionVariant interface {
	SearchInputRequestDefinitionCaster() *SearchInputRequestDefinition
}

func (s *SearchInputRequestDefinition) SearchInputRequestDefinitionCaster() *SearchInputRequestDefinition {
	_ = "STUB: not implemented"
	return nil
}
