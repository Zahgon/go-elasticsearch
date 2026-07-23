package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationalternate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationcasefirst"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationdecomposition"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationstrength"
)

type IcuCollationTokenFilter struct {
	Alternate              *icucollationalternate.IcuCollationAlternate         `json:"alternate,omitempty"`
	CaseFirst              *icucollationcasefirst.IcuCollationCaseFirst         `json:"caseFirst,omitempty"`
	CaseLevel              *bool                                                `json:"caseLevel,omitempty"`
	Country                *string                                              `json:"country,omitempty"`
	Decomposition          *icucollationdecomposition.IcuCollationDecomposition `json:"decomposition,omitempty"`
	HiraganaQuaternaryMode *bool                                                `json:"hiraganaQuaternaryMode,omitempty"`
	Language               *string                                              `json:"language,omitempty"`
	Numeric                *bool                                                `json:"numeric,omitempty"`
	Rules                  *string                                              `json:"rules,omitempty"`
	Strength               *icucollationstrength.IcuCollationStrength           `json:"strength,omitempty"`
	Type                   string                                               `json:"type,omitempty"`
	VariableTop            *string                                              `json:"variableTop,omitempty"`
	Variant                *string                                              `json:"variant,omitempty"`
	Version                *string                                              `json:"version,omitempty"`
}

func (s *IcuCollationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s IcuCollationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIcuCollationTokenFilter() *IcuCollationTokenFilter { _ = "STUB: not implemented"; return nil }

type IcuCollationTokenFilterVariant interface {
	IcuCollationTokenFilterCaster() *IcuCollationTokenFilter
}

func (s *IcuCollationTokenFilter) IcuCollationTokenFilterCaster() *IcuCollationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *IcuCollationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
